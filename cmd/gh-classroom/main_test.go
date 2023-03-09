package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssignmentsCommandShowsHelp(t *testing.T) {
	actual := bytes.Buffer{}
	expected := "Display a list of assignments for a classroom.\n\nUsage:\n  classroom assignments [flags]\n\nFlags:\n      --classroom-id int   ID of the classroom\n  -h, --help               help for assignments\n      --page int           Page number (default 1)\n      --per-page int       Number of assignments per page (default 30)\n      --web                Open the assignment list in a browser\n"
	cmd := rootCmd()
	cmd.SetOut(&actual)
	cmd.SetArgs([]string{"assignments", "--help"})
	cmd.Execute()
	assert.Equal(t, expected, actual.String())
}
