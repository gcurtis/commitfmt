commitfmt
=========

commitfmt is a git hook that helps validate that a commit message is properly
formatted.

A well formatted commit message looks something like this:

	Capitalized summary that is 50 characters or less

	A optional body that is separated from the subject by a blank line. It
	should be hard-wrapped to 72 characters and consist of full sentences.

	Messsages should be descriptive and written in the imperative. For
	example, a message with the subject "Fixed build error" is incorrect. A
	better subject would be "Fix build error due to misspelled method".

	Paragraphs in the body are separated by blank lines.

	* Lists are also commonly used
	* Bullets are usually astericks or hyphens separated from the text by a
      single space.

If a commit message has formatting errors, commitfmt will let you know where
they occur:

	$ git commit -m "uncapitalized subject that goes beyond the 50 character limit"
	[1:1] subj-sentence-case: the subject should adhere to sentence casing, i.e., only the first letter of the first word should be capitalized.
		uncapitalized subject that goes beyond the 50 character limit
		^
	[1:51] subj-len: the subject should not exceed 50 characters.
		uncapitalized subject that goes beyond the 50 character limit
		                                                  ^
	2 formatting errors were found.

Install
-------

The easiest way is to download the [latest release][1] on GitHub.

If you have Go installed, you can also install it with
`go get github.com/gcurtis/commitfmt`.

[1]: https://github.com/gcurtis/commitfmt/releases

Usage
-----

commitfmt can be called either from an existing hook with
`commitfmt <message-file>` or by linking to it in your repo's git hooks
directory. For example: `ln -s commitfmt ~/my-repo/.git/hooks/commit-msg`.

There are times when commitfmt may incorrectly return an error. For example,
commitfmt will complain if your message has a long URL that goes past the 72
character limit, even though it may be a properly formatted message. In which
case, you can tell git to bypass the commitfmt hook by running
`git commit --no-verify`.

Rules
-----

Descriptions of commitfmt's rules and instructions on how to configure them can
be found in [docs/rules.md](docs/rules.md).

If you're interested in making your own rules, there's documentation on how to
do so in the [godoc](http://godoc.org/github.com/gcurtis/commitfmt/rules) as
well as the [contributing guide](CONTRIBUTING.md).
