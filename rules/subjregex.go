package rules

import (
	"fmt"
	"regexp"
)

// SubjRegex checks that the commit subject matches a configured regex.
var SubjRegex = &subjRegex{DefaultConf: map[string]interface{}{
	"pattern": nil,
}}

type subjRegex struct {
	DefaultConf map[string]interface{}
	pattern     *regexp.Regexp
}

func (rule *subjRegex) Name() string {
	return "subj-regex"
}

func (rule *subjRegex) Desc() string {
	if rule.pattern != nil {
		return fmt.Sprintf(`the subject must match the regex "%s".`,
			rule.pattern.String())
	}
	return "the subject must match a configured regex."
}

func (rule *subjRegex) Config(conf map[string]interface{}) (err error) {
	inter, ok := conf["pattern"]
	if !ok {
		return
	}

	if inter == nil {
		rule.pattern = nil
		return
	}

	regexpStr, ok := inter.(string)
	if !ok {
		err = fmt.Errorf("the pattern must be a string")
		return
	}

	rule.pattern, err = regexp.Compile(regexpStr)
	if err != nil {
		err = fmt.Errorf("the pattern must be a valid regular expression")
		return
	}

	return
}

func (rule *subjRegex) Check(subject string, body string) []Violation {
	if rule.pattern == nil {
		return nil
	}

	if !rule.pattern.MatchString(subject) {
		return []Violation{Violation{rule, 0}}

	}
	return nil
}
