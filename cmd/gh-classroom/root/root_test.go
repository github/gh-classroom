package root

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignmentsCommandShowsHelp(t *testing.T) {
	actual := bytes.Buffer{}
	expected := "Display a list of assignments for a classroom.\n\nUsage:\n  classroom assignments [flags]\n\nFlags:\n  -c, --classroom-id int   ID of the classroom\n  -h, --help               help for assignments\n      --page int           Page number (default 1)\n      --per-page int       Number of assignments per page (default 30)\n      --web                Open the assignment list in a browser\n"
	cmd := NewRootCmd()
	cmd.SetOut(&actual)
	cmd.SetArgs([]string{"assignments", "--help"})
	cmd.Execute()

	assert.Equal(t, expected, actual.String())
}
