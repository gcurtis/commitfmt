package rules

import (
	"strings"
)

// SubjOneLine checks that the subject doesn't span multiple lines.
var SubjOneLine = &subjOneLine{}

type subjOneLine struct{}

func (rule *subjOneLine) Name() string {
	return "subj-one-line"
}

func (rule *subjOneLine) Desc() string {
	return "the subject should not span multiple lines. Make sure there are " +
		"two newlines between the subject and body."
}

func (rule *subjOneLine) Config(conf map[string]interface{}) {

}

func (rule *subjOneLine) Check(subject string, body string) []Violation {
	if index := strings.Index(subject, "\n"); index != -1 {
		return []Violation{Violation{rule, index}}
	}
	return nil
}
