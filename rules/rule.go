/*
Package rules contains the various formatting rules used by commitfmt.

Adding Rules

Adding new rules is simple. Create a new type that satisfies rules.Interface,
create a global instance for the rule, and then add that instance to rules.All.
Here is an example "subj-prefix" rule that checks that the commit subject begins
with a certain string.

	// subjprefix.go

	// Exported global variable for the rule. This variable should also be added
	// to the rules.All slice.
	var SubjPrefix = &subjPrefix{}

	// Rule types are unexported to keep the package's API clean.
	type subjPrefix struct{
		prefix string
	}

	// Name returns the name of your rule. The name should be all lowercase and
	// words should be separated by hyphens.
	func (rule *subjPrefix) Name() string {
		return "my-rule"
	}

	// Desc returns a description for your rule. The description should start
	// with a lowercase letter, be one to two sentences and end with a period.
	// Notice how the description changes if the user has configured the rule.
	// This helps the user figure out what's wrong if they violate the rule.
	func (rule *subjPrefix) Desc() string {
		if rule.prefix != "" {
			return fmt.Sprintf(`the commit subject must begin with "%s".`,
				rule.prefix)
		}
		return "the commit subject must begin with a configured prefix."
	}

	// Config configures the rule with a map of settings. This rule allows the
	// user to configure the subject prefix it should check for.
	func (rule *subjPrefix) Config(conf map[string]interface{}) error {
		rule.prefix, ok := conf["prefix"].(string)
		return nil
	}

	// Check should return any violations of your rule (or nil if there aren't
	// any). This example rule checks that the commit subject starts with a
	// configured string. If there is no configured prefix, it is just skipped.
	func (rule *subjPrefix) Check(subject string, body string) []Violation {
		if rule.prefix == "" {
			return nil
		}

		if !strings.HasPrefix(subject, rule.prefix) {
			return []Violation{Violation{rule, 0}}
		}
		return nil
	}

As long as your rule is added to rules.All, it will be automatically be picked
up and checked by commitfmt.

Remember that when calculating the position of a violation, you must take into
account the subject, two newlines, and the body. So if a violation occurs at
index 0 in the body, your rule should return the position len(subject) + 2.

Rule Configuration

Before rules are checked, they will be configured with any user-specified
settings. It completely up to each rule to define and document any settings that
the user can specify. If a rule requires configuration and none is provided,
then the rule should silently skip itself by returning nil when Check is called.

If an error occurs during configuration (for example, if the value of a setting
is the wrong type), then a human-readable error should be returned. Returning an
error will immediately abort any checking and the error will be printed to the
user. Since the error message will be shown to the user, it should follow the
same formatting conventions as a rule's description - start with a lowercase
letter, be one to two sentences and end with a period.

Sometimes it's a good idea to change the rule's description based on its
configuration. For example, a rule's default description might be "the subject
should start with a configured prefix" but after configuration it changes to
"the subject should start with 'Foo'" to be more helpful to the user if the rule
is violated.

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
	// a lowercase letter, be one to two sentences and end with a period. The
	// description will be shown to the user if there's an error, so it should
	// provide helpful information about why rule was violated.
	Desc() string

	// Config configures the rule with a map of settings. If an error is
	// returned, then checking will be aborted and the error will be printed to
	// the user. The map will never be nil.
	Config(conf map[string]interface{}) error

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
