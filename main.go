// Package contains the entry point for the commitfmt command.
package main

import (
	"fmt"
	"github.com/gcurtis/commitfmt/rules"
	"io/ioutil"
	"os"
	"strings"
)

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

	report := runRules(msg)
	fmt.Println(report.string())
	if len(report.violations) > 0 {
		os.Exit(1)
	}
}

// runRules parses a commit message and then enforces every rule found in the
// rules package.
func runRules(msg string) (rep *report) {
	msg = strings.TrimSpace(msg)
	rep = &report{msg: msg}
	subject, body := parseMsg(msg)

	for _, rule := range rules.All {
		violations := rule.Enforce(subject, body)
		rep.append(violations...)
	}

	return
}

// parseMsg parses a message by breaking it up into a subject and a body.
func parseMsg(msg string) (subject string, body string) {
	split := strings.SplitN(msg, "\n\n", 2)
	subject = split[0]
	if len(split) > 1 {
		body = split[1]
	}
	return
}
