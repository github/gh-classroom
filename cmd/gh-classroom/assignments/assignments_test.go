package assignments

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestOutputsHelpTestWithNoArgs(t *testing.T) {
	gock.New("https://api.github.com").
		Get("/classrooms/1234/assignments").
		Reply(200).
		JSON(`[{
    "id": 1,
    "public_repo": false,
    "title": "New assignment here",
    "type": "individual",
    "invite_link": "http://github.localhost/assignment-invitations/594b54b4dcffafea7d9671116e7ae8d4",
    "invitations_enabled": true,
    "slug": "new-assignment-here",
    "students_are_repo_admins": false,
    "feedback_pull_requests_enabled": false,
    "max_teams": null,
    "max_members": null,
    "editor": null,
    "accepted": 0,
    "submissions": 0,
    "passing": 0,
    "language": null,
    "classroom": {
      "id": 1,
      "name": "Classroom over api",
      "archived": false,
      "url": "https://classroom.github.com/classrooms/146-classroom-over-api"
    }
  }]`)
	actual := new(bytes.Buffer)
	command := NewCmdAssignments()
	command.SetOut(actual)
	command.SetErr(actual)
	command.SetArgs([]string{
		fmt.Sprintf("--%s=%s", "classroom-id", "1234"),
	})

	command.Execute()

	expected := "hello there"

	assert.Equal(t, actual.String(), expected, "Actual output should match expected output")
}
