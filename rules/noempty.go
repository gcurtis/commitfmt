package rules

// NoEmpty checks that the commit message is not empty.
var NoEmpty = &noEmpty{}

type noEmpty struct{}

func (rule *noEmpty) Desc() string {
	return "no-empty: the commit message cannot be empty."
}

func (rule *noEmpty) Check(subject string, body string) []Violation {
	if subject == "" {
		return []Violation{Violation{rule, 0}}
	}
	return nil
}
