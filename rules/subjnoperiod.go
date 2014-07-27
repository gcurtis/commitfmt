package rules

import (
	"strings"
)

// SubjNoPeriod enforces that the subject does not have a period.
var SubjNoPeriod = &subjNoPeriod{}

type subjNoPeriod struct{}

func (rule *subjNoPeriod) Desc() string {
	return "subj-no-period: the subject should not have a period."
}

func (rule *subjNoPeriod) Enforce(subject string, body string) []Violation {
	if strings.HasSuffix(subject, "...") {
		return nil
	}

	if strings.HasSuffix(subject, ".") {
		return []Violation{Violation{rule, len(subject) - 1}}
	}

	return nil
}
