package rules

import (
	"strings"
	"unicode"
)

// SubjSentenceCase checks that the subject adheres to sentence casing, i.e.,
// only the first letter of the first word should be capitalized. This rule does
// its best to detect proper capitalization, but it will need to be ignored for
// pronouns (e.g., "Fix references to Java libraries" will incorrectly trigger
// this rule).
var SubjSentenceCase = &subjSentenceCase{}

type subjSentenceCase struct{}

func (rule *subjSentenceCase) Name() string {
	return "subj-sentence-case"
}

func (rule *subjSentenceCase) Desc() string {
	return "the subject should adhere to sentence casing, i.e., only the " +
		"first letter of the first word should be capitalized."
}

func (rule *subjSentenceCase) Config(conf map[string]interface{}) error {
	return nil
}

func (rule *subjSentenceCase) Check(subject string, body string) []Violation {
	if len(subject) == 0 {
		return nil
	}

	var violations []Violation
	if !unicode.IsUpper(rune(subject[0])) {
		violations = append(violations, Violation{rule, 0})
	}

	words := strings.Split(subject, " ")
	pos := len(words[0]) + 1
	for _, w := range words[1:] {
		if len(w) == 0 {
			continue
		}

		if unicode.IsUpper(rune(w[0])) {
			if !isException(w) {
				violations = append(violations, Violation{rule, pos})
			}
		}

		pos += len(w) + 1
	}

	return violations
}

// isException returns true if a word doesn't violate the rule even though it is
// capitalized in the middle of a sentence.
func isException(word string) bool {
	for _, c := range word[1:] {
		if unicode.IsUpper(c) {
			return true
		}
	}

	return false
}
