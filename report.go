package main

import (
	"bytes"
	"fmt"
	"github.com/gcurtis/commitfmt/rules"
	"sort"
	"strings"
)

// report contains a list of rules that were violated in a commit message.
type report struct {
	msg        string            // msg is the commit message.
	violations []rules.Violation // violations is a list of rule violations.
}

// append adds a violation to the report.
func (rep *report) append(v ...rules.Violation) {
	rep.violations = append(rep.violations, v...)
}

// string creates a human-readable string from the report.
func (rep *report) string() string {
	sort.Sort(rep)
	str := ""
	for _, v := range rep.violations {
		lineStart, lineNum, charNum := rep.lineChar(v.Pos)
		desc := v.Rule.Desc()
		context := rep.context(lineStart, charNum, "\t")
		str += fmt.Sprintf("[%d:%d] %s\n%s\n", lineNum, charNum, desc, context)
	}
	str += fmt.Sprintf("%d formatting errors were found.", len(rep.violations))

	return str
}

// lineChar takes a position in the commit message and returns the starting
// point of the line that the position is on, the position's line number and the
// position's character number.
func (rep *report) lineChar(pos int) (lineStart int, lineNum int, charNum int) {
	lineNum = 1
	charNum = 1
	for i := 0; i < pos; i++ {
		if rep.msg[i] == '\n' {
			lineNum++
			lineStart = i + 1
			charNum = 0
		} else {
			charNum++
		}
	}
	return
}

// context creates a "context string" that points to where the error occurred
// within the commit message.
func (rep *report) context(lineStart int, charNum int, prefix string) string {
	line := rep.msg[lineStart:]
	index := strings.Index(line, "\n")
	if index != -1 {
		line = line[:index]
	}

	buf := bytes.Buffer{}
	buf.WriteString(prefix)
	buf.WriteString(line)
	buf.WriteRune('\n')
	buf.WriteString(prefix)
	for i := 0; i < charNum-1; i++ {
		buf.WriteRune(' ')
	}
	buf.WriteRune('^')

	return buf.String()
}

// Len satisfies sort.Interface.
func (rep *report) Len() int {
	return len(rep.violations)
}

// Swap satisfies sort.Interface.
func (rep *report) Swap(i, j int) {
	rep.violations[i], rep.violations[j] = rep.violations[j], rep.violations[i]
}

// Less satisfies sort.Interface.
func (rep *report) Less(i, j int) bool {
	return rep.violations[i].Pos < rep.violations[j].Pos
}
