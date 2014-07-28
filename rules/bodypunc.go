package rules

import (
	"strings"
	"unicode"
)

// BodyPunc checks that the body ends with valid punctuation (".", "!", "?")
// unless it ends with a list.
var BodyPunc = &bodyPunc{}

type bodyPunc struct{}

func (rule *bodyPunc) Name() string {
	return "body-punc"
}

func (rule *bodyPunc) Desc() string {
	return `the body should end with valid punctuation (".", "!", "?") unless` +
		` it ends with a list.`
}

func (rule *bodyPunc) Check(subject string, body string) []Violation {
	if len(body) == 0 {
		return nil
	}

	lines := strings.Split(body, "\n")
	lastLine := strings.TrimSpace(lines[len(lines)-1])
	if !inList(lastLine) {
		if !endsWithPunc(lastLine) {
			return []Violation{Violation{rule, len(subject) + len(body) + 2}}
		}
	}

	return nil
}

// inList returns true if a string looks like it belongs to a list.
func inList(s string) bool {
	return s[0] == '-' || s[0] == '+' || s[0] == '*'
}

// endsWithPunc returns true if a string ends with punctuation.
func endsWithPunc(s string) bool {
	last := rune(s[len(s)-1])
	return unicode.IsPunct(last)
}
