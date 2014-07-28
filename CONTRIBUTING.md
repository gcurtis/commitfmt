Contributing
============

commitfmt was designed to be easily extended. Creating new rules is simple, and there's [plenty of documentation](http://godoc.org/github.com/gcurtis/commitfmt/rules) on how to do it. We're also open to accepting any new features that improve the commitfmt command itself. Just follow these three steps if you want to contribute:

1. Whether you've found a bug or want to add a rule/feature, start by opening an issue on GitHub to get some feedback on what you plan to do.
2. Go ahead and implement your change! Be sure to include tests and documentation for all of your code. You can also find the documentation for commitfmt at <http://godoc.org/github.com/gcurtis/commitfmt>.
3. Send a pull request with your change (preferably as one commit, but it can be multiple if it makes sense). Be sure to reference the issue at the bottom of the commit body with "Closes #123." or "Fixes #123.". Also don't forget to run commitfmt on your message!

Adding New Rules
----------------

* Please add any tests for rules to `main_test.go` as opposed to a new test file. This helps test all of the rules together to ensure that none of them conflict.
* Remember to add your rule and its description to the README.
