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
			"name": "Classroom Name"
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
	assert.Equal(t, "Classroom Name", actual.Classroom.Name)
}

func TestListClassrooms(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/classrooms").
		Reply(200).
		JSON(`[{"id": 1,
		"name": "Classroom Name"
	}]`)

	client, err := gh.RESTClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ListClassrooms(client, 1, 30)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(actual))
	assert.Equal(t, "Classroom Name", actual[0].Name)
}

func TestListAcceptedAssignments(t *testing.T) {
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

	client, err := gh.RESTClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ListAcceptedAssignments(client, 1, 1, 30)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, actual.Count)
	assert.Equal(t, 1, actual.AcceptedAssignments[0].Id)
	assert.Equal(t, "org1/student1-repo", actual.AcceptedAssignments[0].Repository.FullName)
	assert.Equal(t, "student1", actual.AcceptedAssignments[0].Students[0].Login)
}

func TestGetAssignment(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/assignments/1").
		Reply(200).
		JSON(`{"id": 1,
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
	}`)

	client, err := gh.RESTClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := GetAssignment(client, 1)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, actual.Id)
	assert.Equal(t, "Assignment 1", actual.Title)
	assert.Equal(t, "org1/starter-code-repo", actual.StarterCodeRepository.FullName)
	assert.Equal(t, "Classroom Name", actual.Classroom.Name)
}

func TestGetClassroom(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/classrooms/1").
		Reply(200).
		JSON(`{"id": 1,
		"name": "Classroom Name"
	}`)

	client, err := gh.RESTClient(nil)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := GetClassroom(client, 1)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, actual.Id)
	assert.Equal(t, "Classroom Name", actual.Name)
}
