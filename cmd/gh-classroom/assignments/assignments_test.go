package assignments

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestListingAssignments(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/classrooms/1234/assignments").
		MatchParam("page", "1").
		MatchParam("per_page", "30").
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
		"assignments",
		"-c1234",
	})

	err := command.Execute()
	assert.NoError(t, err, "Should not error")
	expected := "1 Assignment for Classroom over api\n\n" + 
	"ID\tTitle\tSubmission Public\tType\tEditor\tInvitation Link\tAccepted\tSubmissions\tPassing\n" + 
	"1\tNew assignment here\tfalse\tindividual\t\thttp://github.localhost/assignment-invitations/594b54b4dcffafea7d9671116e7ae8d4\t0\t0\t0\n"

	assert.Equal(t, expected, actual.String(), "Actual output should match expected output")
}
