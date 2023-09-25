package shared

import (
	"testing"

	"github.com/cli/go-gh"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestListAcceptedAssignments(t *testing.T) {
	t.Setenv("GITHUB_TOKEN", "999")
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/assignments/1/accepted_assignments").
		Reply(200).
		JSON(`[{"id": 1,
		"assignment": {
			"id": 1,
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
		},
		"students": [{
			"id": 1,
			"login": "student1"
		}],
		"repository": {
			"id": 1,
			"full_name": "org1/student1-repo"
		}
	}]`)

	gock.New("https://api.github.com").
		Path("/assignments/1/accepted_assignments").
		Reply(200).
		JSON(`[{"id": 2,
		"assignment": {
			"id": 2,
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
		},
		"students": [{
			"id": 2,
			"login": "student2"
		}],
		"repository": {
			"id": 2,
			"full_name": "org1/student2-repo"
		}
	}]`)

	client, err := gh.RESTClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	//Ask for page 1 and 1 per page
	actual, err := ListAcceptedAssignments(client, 1, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, actual.Count)
	assert.Equal(t, 1, actual.AcceptedAssignments[0].Id)
	assert.Equal(t, "org1/student1-repo", actual.AcceptedAssignments[0].Repository.FullName)
	assert.Equal(t, "student1", actual.AcceptedAssignments[0].Students[0].Login)

	//Ask for page 2 and 1 per page
	actual, err = ListAcceptedAssignments(client, 1, 2, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, actual.Count)
	assert.Equal(t, 2, actual.AcceptedAssignments[0].Id)
	assert.Equal(t, "org1/student2-repo", actual.AcceptedAssignments[0].Repository.FullName)
	assert.Equal(t, "student2", actual.AcceptedAssignments[0].Students[0].Login)

}
