package accepted

import (
	"bytes"
	"testing"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestListingAssignments(t *testing.T) {
	t.Run("with no accepted asignments", func(t *testing.T) {
		defer gock.Off()

		t.Setenv("GITHUB_TOKEN", "999")

		gock.New("https://api.github.com").
			Get("/assignments/1234/accepted_assignments").
			Reply(200).
			JSON(`[]`)

		actual := new(bytes.Buffer)

		f := &cmdutil.Factory{}
		command := NewCmdAcceptedAssignments(f)
		command.SetOut(actual)
		command.SetErr(actual)
		command.SetArgs([]string{
			"-a1234",
		})

		err := command.Execute()
		assert.NoError(t, err, "Should not error")

		expected := "Assignment:  \nID: 0 \n\nID\tSubmitted\tPassing\tCommit Count\tGrade\tFeedback Pull Request URL\tStudent\tRepository\n"
		assert.Equal(t, expected, actual.String(), "Actual output should match expected output")
	})
}
