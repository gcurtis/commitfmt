package main

import (
	"fmt"
	"github.com/gcurtis/commitfmt/rules"
	"testing"
)

func reportHasViolation(rep *report, r rules.Interface) bool {
	for _, v := range rep.violations {
		if v.Rule == r {
			return true
		}
	}

	return false
}

func TestEmptyMessage(t *testing.T) {
	msg := ""
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.NoEmpty) {
		t.Error("Expected violation:", ruleString(rules.NoEmpty))
	}
}

func TestWhitespaceMessage(t *testing.T) {
	msg := " "
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.NoEmpty) {
		t.Error("Expected violation:", ruleString(rules.NoEmpty))
	}
}

func TestValidSubject(t *testing.T) {
	msg := "Subject"
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func TestValidSubjectWithBody(t *testing.T) {
	msg := "Subject\n\nBody."
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func TestMultilineSubject(t *testing.T) {
	msg := "Subject1\nSubject2"
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.SubjOneLine) {
		t.Error("Expected violations:", ruleString(rules.SubjOneLine))
	}
}

func TestSubjectThatIsTooLong(t *testing.T) {
	msg := "This subject line goes over 50 characters=========="
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.SubjLen) {
		t.Error("Expected violations:", ruleString(rules.SubjLen))
	}
}

func TestSubjectWithTitleCase(t *testing.T) {
	msg := "This Subject Is Incorrectly Title Cased"
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.SubjSentenceCase) {
		t.Error("Expected violations:", ruleString(rules.SubjSentenceCase))
	}
}

func TestSubjectWithExtraCapitalizedWords(t *testing.T) {
	msg := "This subject is Incorrectly cased"
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.SubjSentenceCase) {
		t.Error("Expected violations:", ruleString(rules.SubjSentenceCase))
	}
}

func TestSubjectWithAcronym(t *testing.T) {
	msg := "Subject with the acronym ID"
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func TestSubjectWithCamelCase(t *testing.T) {
	msg := "Subject with the class name MyClass in it"
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func TestSubjectWithPeriod(t *testing.T) {
	msg := "This subject ends with a period."
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.SubjNoPeriod) {
		t.Error("Expected violations:", ruleString(rules.SubjNoPeriod))
	}
}

func TestSubjectWithEllipsis(t *testing.T) {
	msg := "This subject ends with ellipsis..."
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func TestBodyWithMultipleParagraphs(t *testing.T) {
	msg := `Subject

Paragraph1.

Paragraph2 with
multiple lines.`
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func TestSubjectWithMultipleSpaces(t *testing.T) {
	msg := "Subject  with multiple spaces"
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.Whitespace) {
		t.Error("Expected violations:", ruleString(rules.Whitespace))
	}
}

func TestBodyWithMultipleSpaces(t *testing.T) {
	msg := `Subject

Body with  multiple spaces.`
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.Whitespace) {
		t.Error("Expected violations:", ruleString(rules.Whitespace))
	}
}

func TestBodyWithMultipleNewlines(t *testing.T) {
	msg := `Subject

Paragraph 1.

Paragraph 2.


Paragraph 3.`
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.Whitespace) {
		t.Error("Expected violations:", ruleString(rules.Whitespace))
	}
}

func TestSubjectWithTrailingSpace(t *testing.T) {
	msg := "Subject with trailing space \n\nBody."
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.Whitespace) {
		t.Error("Expected violations:", ruleString(rules.Whitespace))
	}
}

func TestBodyWithTrailingSpace(t *testing.T) {
	msg := "Subject\n\nParagraph1. \n\nParagraph2."
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.Whitespace) {
		t.Error("Expected violations:", ruleString(rules.Whitespace))
	}
}

func TestBodyWithLineThatIsTooLong(t *testing.T) {
	msg := `Subject

Paragraph that is longer that 72 characters=============================.`
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.BodyLen) {
		t.Error("Expected violations:", ruleString(rules.BodyLen))
	}
}

func TestBodyThatDoesNotEndWithPunctuation(t *testing.T) {
	msg := `Subject

Paragraph that doesn't end with punctuation`
	rep := runRules(msg)

	if !reportHasViolation(rep, rules.BodyPunc) {
		t.Error("Expected violations:", ruleString(rules.BodyPunc))
	}
}

func TestBodyListThatDoesNotEndWithPunctuation(t *testing.T) {
	msg := `Subject

* This is a list item`
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func TestFullMessage(t *testing.T) {
	msg := `Capitalized, short (50 chars or less) summary

More detailed explanatory text, if necessary. Wrap it to about 72
characters or so. In some contexts, the first line is treated as the
subject of an email and the rest of the text as the body. The blank
line separating the summary from the body is critical (unless you omit
the body entirely); tools like rebase can get confused if you run the
two together.

Write your commit message in the imperative: "Fix bug" and not "Fixed
bug" or "Fixes bug." This convention matches up with commit messages
generated by commands like git merge and git revert.

Further paragraphs come after blank lines.

- Bullet points are okay, too

- Typically a hyphen or asterisk is used for the bullet, followed by a
  single space, with blank lines in between, but conventions vary here

- Use a hanging indent`
	rep := runRules(msg)

	if rep.violations != nil {
		t.Error("Unexpected violations:", rep.string())
	}
}

func Example1() {
	msg := "This subject is longer than 50 characters and will trigger an error"
	rep := runRules(msg)
	fmt.Println(rep.string())

	// Output: [1:51] subj-len: the subject should not exceed 50 characters.
	// 	This subject is longer than 50 characters and will trigger an error
	// 	                                                  ^
	// 1 formatting errors were found.
}

func Example2() {
	msg := `This commit message has a Number of different violations that will be caught.

The body is way too long and goes beyond 72 characters per line.  There are unnecessary spaces
in between words and the body doesn't end with punctuation`
	rep := runRules(msg)
	fmt.Println(rep.string())

	// Output: [1:27] subj-sentence-case: the subject should adhere to sentence casing, i.e., only the first letter of the first word should be capitalized.
	// 	This commit message has a Number of different violations that will be caught.
	// 	                          ^
	// [1:51] subj-len: the subject should not exceed 50 characters.
	// 	This commit message has a Number of different violations that will be caught.
	// 	                                                  ^
	// [1:77] subj-no-period: the subject should not end with a period.
	// 	This commit message has a Number of different violations that will be caught.
	// 	                                                                            ^
	// [3:65] whitespace: there should not be any unnecessary spacing, i.e., only one line break between paragraphs, only one space between words, and no trailing whitespace.
	// 	The body is way too long and goes beyond 72 characters per line.  There are unnecessary spaces
	// 	                                                                ^
	// [3:72] body-len: each line of the body should not exceed 72 characters.
	// 	The body is way too long and goes beyond 72 characters per line.  There are unnecessary spaces
	// 	                                                                       ^
	// [4:58] body-punc: the body should end with valid punctuation (".", "!", "?") unless it ends with a list.
	// 	in between words and the body doesn't end with punctuation
	// 	                                                         ^
	// 6 formatting errors were found.
}
