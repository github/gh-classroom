package list

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestListingClassrooms(t *testing.T) {
	defer gock.Off()
	t.Setenv("GITHUB_TOKEN", "999")

	gock.New("https://api.github.com"). 
	Get("/classrooms").
	MatchParam("page", "1").
	MatchParam("per_page", "30").
	Reply(200).
	JSON(`[{
		"id": 1,
		"name": "Classroom over api",
		"archived": false,
		"url": "https://classroom.github.com/classrooms/146-classroom-over-api"
	}]`)

	actual := new(bytes.Buffer)
	
	command := NewCmdList()
	command.SetOut(actual)
	command.SetErr(actual)

	err := command.Execute()
	assert.NoError(t, err, "Should not error")

	expected := "1 Classroom\nID  Name                Archived  URL\n1   Classroom over api  false     https://classroom.github.com/classrooms/146-classroom-over-api\n"

	assert.Equal(t, expected, actual.String(), "Actual output should match expected output")
}
