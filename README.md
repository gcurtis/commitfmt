commitfmt
=========

commitfmt is a git hook that helps validate that a commit message is properly formatted.

Install
-------

The easiest way is to download the [latest release](https://github.com/gcurtis/commitfmt/releases) on GitHub.

If you have Go installed, you can also install it with `go get github.com/gcurtis/commitfmt` and find it in `$GOPATH/bin/commitfmt`.

Usage
-----

commitfmt can be called either from an existing hook with `commitfmt <message-file>` or by linking to it in your repo's git hooks directory. For example: `ln -s commitfmt ~/my-repo/.git/hooks/commit-msg`.

There are times when commitfmt may incorrectly return an error. For example, commitfmt will complain if your message has a long URL that goes past the 72 character limit, even though it may be a properly formatted message. In which case, you can tell git to bypass the commitfmt hook by running `git commit --no-verify`.

Rules
-----

A commit message should have a descriptive subject, an optional body, and be hard-wrapped to the appropriate line length. The message itself should be phrased in the imperative. For example, a message with the subject `Fixed build error` is incorrect. A better subject would be `Fix build error due to misspelled method`. The body should have correct spelling/grammar and consist of full sentences.

Rules around spelling and grammar are difficult to check automatically and would result in too many false-positives. However, the following rules can be automatically checked by commitfmt.

### Subject

* subj-sentence-case - the subject should adhere to sentence casing, i.e., only the first letter of the first word should be capitalized. This rule does its best to detect proper capitalization, but it will need to be ignored for pronouns (e.g., "Fix references to Java libraries" will incorrectly trigger this rule).
* subj-no-period - the subject should not end with a period.
* subj-len - the subject should not exceed 50 characters.
* subj-one-line - the subject should not span multiple lines. Make sure there are two newlines between the subject and body.
* subj-regex - the subject should match a regex configured via the "pattern" setting.

### Body

* body-len - each line of the body should not exceed 72 characters. This rule can be ignored for non-prose (e.g., long URLs, build output, etc.).
* body-punc - the body should end with valid punctuation (".", "!", "?") unless it ends with a list.

### General

* whitespace - there should not be any unnecessary spacing, i.e., only one line break between paragraphs, only one space between words, and no trailing whitespace.
* no-empty - the commit message cannot be empty.

Configuring Rules
-----------------

Rules can be configured by creating a `.commitfmt` JSON file in the root of your repo. To disable a rule, set its value to `false` in the conf file. To customize a rule, set its value to a map of the settings you wish to customize. Refer to a rule's documentation to see what settings it provides. For example:

```json
{
    "subj-sentence-case": false,
    "subj-regex": {
        "pattern": "^Ticket: .+"
    }
}
```
