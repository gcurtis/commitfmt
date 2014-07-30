package rules

import (
	"strings"
)

// BodyLen checks that each line of the body does not exceed 72 characters.
var BodyLen = &bodyLen{}

type bodyLen struct{}

func (rule *bodyLen) Name() string {
	return "body-len"
}

func (rule *bodyLen) Desc() string {
	return "each line of the body should not exceed 72 characters."
}

func (rule *bodyLen) Config(conf map[string]interface{}) error {
	return nil
}

func (rule *bodyLen) Check(subject string, body string) []Violation {
	var violations []Violation
	offset := len(subject) + 2

	lines := strings.Split(body, "\n")
	for _, l := range lines {
		if len(l) > 72 {
			violations = append(violations, Violation{rule, offset + 72})
		}
		offset += len(l) + 1
	}

	return violations
}
