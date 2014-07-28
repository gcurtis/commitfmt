package rules

import (
	"strings"
)

// SubjNoPeriod checks that the subject does not end with a period.
var SubjNoPeriod = &subjNoPeriod{}

type subjNoPeriod struct{}

func (rule *subjNoPeriod) Name() string {
	return "subj-no-period"
}

func (rule *subjNoPeriod) Desc() string {
	return "the subject should not end with a period."
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
