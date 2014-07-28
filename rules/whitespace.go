package rules

// Whitespace checks that there isn't any unnecessary spacing, i.e., only one
// line break between paragraphs, only one space between words, and no trailing
// whitespace.
var Whitespace = &whitespace{}

type whitespace struct{}

func (rule *whitespace) Name() string {
	return "whitespace"
}

func (rule *whitespace) Desc() string {
	return "there should not be any unnecessary spacing, i.e., only one line " +
		"break between paragraphs, only one space between words, and no " +
		"trailing whitespace."
}

func (rule *whitespace) Check(subject string, body string) []Violation {
	msg := subject + "\n\n" + body
	var violations []Violation
	space := 0
	newline := 0
	seenWord := false
	for i, c := range msg {
		if c == ' ' {
			space++
			if seenWord {
				if space > 1 {
					violations = append(violations, Violation{rule, i})
				}
				if msg[i+1] == '\n' || msg[i+1] == '\t' {
					violations = append(violations, Violation{rule, i + 1})
				}
			}
		} else if c == '\n' {
			newline++
			seenWord = false
			if newline > 2 {
				violations = append(violations, Violation{rule, i})
			}
		} else {
			space = 0
			newline = 0
			seenWord = true
		}
	}

	return violations
}
