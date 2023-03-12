package classroom

import (
	"testing"

	"github.com/cli/go-gh"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestListAssignments(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/classrooms/1/assignments").
		Reply(200).
		JSON(`[{"id": 1,
		"title": "Assignment 1",
		"description": "This is the first assignment",
		"due_date": "2018-01-01",
		"classroom": {
			"id": 1,
			"name":      "Classroom Name"
		},
		"starter_code_repository": {
			"id": 1,
			"full_name": "org1/starter-code-repo"
		}
	}]`)

	client, err := gh.RESTClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ListAssignments(client, 1, 1, 30)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, actual.Count)
	assert.Equal(t, "Assignment 1", actual.Assignments[0].Title)
	assert.Equal(t, "org1/starter-code-repo", actual.Assignments[0].StarterCodeRepository.FullName)
	assert.Equal(t, "Classroom Name", actual.Classroom.Name)
}
