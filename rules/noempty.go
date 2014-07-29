package rules

// NoEmpty checks that the commit message is not empty.
var NoEmpty = &noEmpty{}

type noEmpty struct{}

func (rule *noEmpty) Name() string {
	return "no-empty"
}

func (rule *noEmpty) Desc() string {
	return "the commit message cannot be empty."
}

func (rule *noEmpty) Config(conf map[string]interface{}) {

}

func (rule *noEmpty) Check(subject string, body string) []Violation {
	if subject == "" {
		return []Violation{Violation{rule, 0}}
	}
	return nil
}
