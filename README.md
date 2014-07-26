commitfmt
=========

commitfmt is a git hook that helps validate that a commit message is properly formatted.

Usage
-----

commitfmt can be called either from an existing hook with `commitfmt <message-file>` or by renaming it to `commit-msg` and placing it directly in your git hooks directory.

There are times when commitfmt may incorrectly return an error. For example, commitfmt will complain if your message has a long URL that goes past the 72 character limit, even though it may be a properly formatted message. In which case, you can tell git to bypass the commitfmt hook by running `git commit --no-verify`.

Rules
-----

A commit message should have a descriptive subject, an optional body, and be hard-wrapped to the appropriate line length. The message itself should be phrased in the imperative. For example, a message with the subject `Fixed build error` is **incorrect**. A correct subject would be `Fix build error due to misspelled method`. The body should have correct spelling/grammar and consist of full sentences.

Rules around spelling and grammar are difficult to check automatically and would result in too many false-positives. However, the following rules will be automatically enforced.

### Subject

* subj-sentence-case - the subject should adhere to sentence casing, i.e., only the first letter of the first word should be capitalized. This rule does its best to detect proper capitalization, but might need to be ignored for certain messages (e.g., "Fix SomeFunc in MyClassName" will incorrectly trigger this rule).
* subj-no-period - the subject should not have a period.
* subj-len - the subject should not exceed 50 characters.
* subj-one-line - the subject should not span multiple lines.

### Body

* body-len - each line of the body should not exceed 72 characters. This rule can be ignored for non-prose (e.g., long URLs, build output, etc.).
* body-punc - the body should end with valid punctuation (".", "!", "?") unless it ends with a list.

### General

* unnecessary-space - there should not be any unnecessary spacing, i.e., only one line break between paragraphs, only one space between words, and no trailing whitespace.
