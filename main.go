// Package commitfmt provides a git hook that validates the formatting of a
// commit message.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gcurtis/commitfmt/rules"
	"io/ioutil"
	"os"
	"strings"
)

// snipLine is the special line recognized by git that tells it to strip the
// rest of a commit message.
const snipLine = "------------------------ >8 ------------------------"

// commentChar is the character git uses for commenting out lines in commit
// messages.
const commentChar = '#'

// confName is the name of the commitfmt configuration file.
const confName = ".commitfmt"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You must provide a path to a file containing"+
			" the commit message.")
		os.Exit(1)
	}

	path := os.Args[1]
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file \"%s\".\n", path)
		os.Exit(1)
	}
	msg := string(bytes)

	conf := readConf()
	cleaned := cleanMsg(msg)
	report := runRules(cleaned, conf)
	fmt.Println(report.string())
	if len(report.violations) > 0 {
		// Make a best-effort to save the commit message and provide the user
		// with some help before exiting.
		if f, err := ioutil.TempFile("", "commitfmt"); err == nil {
			_, err := f.WriteString(cleaned)
			if err == nil {
				fmt.Fprintf(os.Stderr, "\nYour commit message has been saved. "+
					"You can edit your previous commit message with:\n"+
					"\tgit commit -e -F %[1]s\n"+
					"or you can bypass this check with:\n"+
					"\tgit commit --no-verify -e -F %[1]s\n",
					f.Name())
			}
			f.Close()
		}
		os.Exit(1)
	}
}

// runRules parses a cleaned commit message and then checks every rule found in
// the rules package.
func runRules(cleanMsg string, conf map[string]interface{}) (rep *report) {
	rep = &report{msg: cleanMsg}
	subject, body := parseMsg(cleanMsg)

	for _, rule := range rules.All {
		if conf != nil {
			ruleConf, ok := conf[rule.Name()]
			if ok && ruleConf != nil {
				if ruleConf == false {
					continue
				} else {
					rule.Config(ruleConf.(map[string]interface{}))
				}
			}
		}

		violations := rule.Check(subject, body)
		rep.append(violations...)
	}

	return
}

// cleanMsg removes any commented-out or snipped content from a commit message.
func cleanMsg(msg string) string {
	remComments := bytes.Buffer{}
	split := strings.SplitAfter(msg, "\n")
	for _, line := range split {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, string(commentChar)+" "+snipLine) {
			break
		}
		if strings.HasPrefix(trim, string(commentChar)) {
			continue
		}

		remComments.WriteString(line)
	}
	return strings.TrimSpace(remComments.String())
}

// parseMsg parses a cleaned message by breaking it up into a subject and a
// body.
func parseMsg(cleanMsg string) (subject string, body string) {
	split := strings.SplitN(strings.TrimSpace(cleanMsg), "\n\n", 2)
	subject = split[0]
	if len(split) > 1 {
		body = split[1]
	}
	return
}

func readConf() (conf map[string]interface{}) {
	r, err := os.Open(confName)
	if err != nil {
		return
	}

	err = json.NewDecoder(r).Decode(&conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't parse conf file, proceeding with"+
			" default rules.")
	}
	return
}
