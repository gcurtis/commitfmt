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
	report := runRules(msg, conf)
	fmt.Println(report.string())
	if len(report.violations) > 0 {
		os.Exit(1)
	}
}

// runRules parses a commit message and then checks every rule found in the
// rules package.
func runRules(msg string, conf map[string]interface{}) (rep *report) {
	msg = strings.TrimSpace(msg)
	rep = &report{msg: msg}
	subject, body := parseMsg(msg)

	for _, rule := range rules.All {
		if conf != nil {
			ruleConf, ok := conf[rule.Name()]
			if ok {
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

// parseMsg parses a message by breaking it up into a subject and a body. It
// will also remove any commented-out or snipped content.
func parseMsg(msg string) (subject string, body string) {
	remComments := bytes.Buffer{}
	split := strings.SplitAfter(msg, "\n")
	for _, line := range split {
		trim := strings.TrimSpace(line)
		if strings.Contains(trim, snipLine) {
			break
		}
		if strings.HasPrefix(trim, string(commentChar)) {
			continue
		}

		remComments.WriteString(line)
	}

	split = strings.SplitN(strings.TrimSpace(remComments.String()), "\n\n", 2)
	subject = split[0]
	if len(split) > 1 {
		body = split[1]
	}
	return
}

func readConf() (conf map[string]interface{}) {
	r, err := os.Open(confName)
	if err != nil {
		err = json.NewDecoder(r).Decode(&conf)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Couldn't parse conf file, proceeding with"+
				" default rules.")
		}
	}
	return
}
