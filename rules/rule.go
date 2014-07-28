/*
Package rules contains the various formatting rules used by commitfmt.

Adding Rules

Adding new rules is simple. Create a new type that satisfies rules.Interface,
create a global instance for the rule, and then add that instance to rules.All.
Here is an example rule named "my-rule" that checks that the commit subject
begins with the word "Ticket".

	// myrule.go

	// Exported global variable for the rule. This variable should also be added
	// to the rules.All slice.
	var MyRule = &myRule{}

	// Rule types are unexported to keep the package's API clean.
	type myRule struct{}

	// Name returns the name of your rule. The name should be all lowercase and
	// words should be separated by hyphens.
	func (rule *myRule) Name() string {
		return "my-rule"
	}

	// Desc returns a description for your rule. The description should start
	// with a lowercase letter, be one to two sentences and end with a period.
	func (rule *myRule) Desc() string {
		return `the commit subject must begin with "Ticket".`
	}

	// Check should return any violations of your rule (or nil if there aren't
	// any). This example rule checks that the commit subject starts with the
	// word "Ticket".
	func (rule *myRule) Check(subject string, body string) []Violation {
		if !strings.HasPrefix(subject, "Ticket") {
			return []Violation{Violation{rule, 0}}
		}
		return nil
	}

As long as your rule is added to rules.All, it will be automatically be picked
up and checked by commitfmt.

Remember that when calculating the position of a violation, you must take into
account the subject, two newlines, and the body. So if a violation occurs at
index 0 in the body, your rule should return the position len(subject) + 2.

*/
package rules

// Violation points to a position in the commit message where a rule was
// violated.
type Violation struct {
	Rule Interface // Rule is the rule that was violated.
	Pos  int       // Pos is the string index of where the violation occurred.
}

// Interface defines the methods that all rules must implement.
type Interface interface {
	// Name returns the name of the rule. The name should be all lowercase and
	// words should be separated by hyphens.
	Name() string

	// Desc returns a description of the rule. The description should start with
	// a lowercase letter, be one to two sentences and end with a period.
	Desc() string

	// Check takes a commit subject and body, and returns a list of positions
	// where the rule was violated.
	Check(subject string, body string) []Violation
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
