package rules

// SubjLen enforces that the subject doesn't exceed 50 characters.
var SubjLen = &subjLen{}

type subjLen struct{}

func (rule *subjLen) Desc() string {
	return "subj-len: the subject should not exceed 50 characters."
}

func (rule *subjLen) Enforce(subject string, body string) []Violation {
	if len(subject) > 50 {
		return []Violation{Violation{rule, 50}}
	}
	return nil
}
