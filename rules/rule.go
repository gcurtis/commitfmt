/*
Package rules contains the various rules used by commitfmt.

Adding Rules

Adding new rules is simple. Create a new type that satisfies rules.Interface,
create a global instance for the rule, and then add that instance to rules.All.
Here is an example rule named "my-rule" that enforces that the commit subject
begins with the word "Ticket".

	// myrule.go

	// Exported global variable for the rule. This variable should also be added
	// to the rules.All slice.
	var MyRule = &myRule{}

	// Rule types are unexported to keep the package's API clean.
	type myRule struct{}

	// Desc returns a description for your rule. The description should follow
	// the format of: "my-rule: one to two sentence description of the rule
	// ending with a period." Note that when rule names are shown to the user,
	// they use a hyphenated version of their name.
	func (rule *myRule) Desc() string {
		return `my-rule: the commit subject must begin with "Ticket".`
	}

	// Enforce should return any violations of your rule (or nil if there aren't
	// any). This example rule enforces that the commit subject starts with the
	// word "Ticket".
	func (rule *myRule) Enforce(subject string, body string) []Violation {
		if !strings.HasPrefix(subject, "Ticket") {
			return []Violation{Violation{rule, 0}}
		}
		return nil
	}

As long as your rule is added to rules.All, it will be automatically be picked
up and enforced by commitfmt.

*/
package rules

// Violation points to a position in the commit message where a rule was
// violated.
type Violation struct {
	Rule Interface
	Pos  int
}

// Interface defines the methods that all rules must implement.
type Interface interface {
	// Desc returns a description of the rule. The description must follow the
	// format of: "rule-name - one or two sentence description of the rule
	// ending in a period."
	Desc() string

	// Enforce takes a commit subject and body, and returns a list of positions
	// where the rule was violated.
	Enforce(subject string, body string) []Violation
}

// All is a slice of every rule in this package.
var All = []Interface{
	NoEmpty,
	SubjLen,
	SubjOneLine,
	SubjSentenceCase,
	SubjNoPeriod,
	Whitespace,
	BodyLen,
	BodyPunc,
}