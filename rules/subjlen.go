package rules

// SubjLen checks that the subject doesn't exceed 50 characters.
var SubjLen = &subjLen{}

type subjLen struct{}

func (rule *subjLen) Name() string {
	return "subj-len"
}

func (rule *subjLen) Desc() string {
	return "the subject should not exceed 50 characters."
}

func (rule *subjLen) Config(conf map[string]interface{}) {

}

func (rule *subjLen) Check(subject string, body string) []Violation {
	if len(subject) > 50 {
		return []Violation{Violation{rule, 50}}
	}
	return nil
}
