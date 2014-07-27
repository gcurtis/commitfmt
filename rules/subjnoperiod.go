package rules

import (
	"strings"
)

// SubjNoPeriod checks that the subject does not end with a period.
var SubjNoPeriod = &subjNoPeriod{}

type subjNoPeriod struct{}

func (rule *subjNoPeriod) Desc() string {
	return "subj-no-period: the subject should not end with a period."
}

func (rule *subjNoPeriod) Check(subject string, body string) []Violation {
	if strings.HasSuffix(subject, "...") {
		return nil
	}

	if strings.HasSuffix(subject, ".") {
		return []Violation{Violation{rule, len(subject) - 1}}
	}

	return nil
}
