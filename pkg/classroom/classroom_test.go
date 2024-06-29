package classroom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAssignmentList(t *testing.T) {
	t.Run("should create a new assignment list", func(t *testing.T) {
		assignments := []Assignment{
			{Title: "Assignment 1"},
			{Title: "Assignment 2"},
			{Title: "Assignment 3"},
		}

		assignmentList := NewAssignmentList(assignments)

		assert.Equal(t, assignmentList.Assignments, assignments)
	})

	t.Run("counts the number of assignments", func(t *testing.T) {
		assignments := []Assignment{
			{Title: "Assignment 1"},
			{Title: "Assignment 2"},
			{Title: "Assignment 3"},
		}

		assignmentList := NewAssignmentList(assignments)

		assert.Equal(t, assignmentList.Count, 3)
	})

	t.Run("uses the classroom from the first assignment", func(t *testing.T) {
		assignments := []Assignment{
			{Title: "Assignment 1", Classroom: Classroom{Name: "The Classroom"}},
			{Title: "Assignment 2"},
			{Title: "Assignment 3"},
		}

		assignmentList := NewAssignmentList(assignments)

		assert.Equal(t, assignmentList.Classroom.Name, "The Classroom")
	})
}

func TestNewAcceptedAssignmentList(t *testing.T) {
	t.Run("stores the list of accepted assignments", func(t *testing.T) {
		acceptedAssignments := []AcceptedAssignment{
			{Id: 1},
			{Id: 2},
			{Id: 3},
		}

		assignmentList := NewAcceptedAssignmentList(acceptedAssignments)

		assert.Equal(t, assignmentList.AcceptedAssignments, acceptedAssignments)
	})

	t.Run("counts accepted assignments", func(t *testing.T) {
		acceptedAssignments := []AcceptedAssignment{
			{Id: 1},
			{Id: 2},
			{Id: 3},
		}

		assignmentList := NewAcceptedAssignmentList(acceptedAssignments)

		assert.Equal(t, assignmentList.Count, 3)
	})

	t.Run("uses the classroom from the first accepted assignment", func(t *testing.T) {
		acceptedAssignments := []AcceptedAssignment{
			{Id: 1, Assignment: Assignment{Classroom: Classroom{Name: "The Classroom"}}},
			{Id: 2},
			{Id: 3},
		}

		assignmentList := NewAcceptedAssignmentList(acceptedAssignments)

		assert.Equal(t, assignmentList.Classroom.Name, "The Classroom")
	})

	t.Run("uses the first accepted assignment's assignment", func(t *testing.T) {
		acceptedAssignments := []AcceptedAssignment{
			{Id: 1, Assignment: Assignment{Title: "The Assignment"}},
			{Id: 2},
			{Id: 3},
		}

		assignmentList := NewAcceptedAssignmentList(acceptedAssignments)

		assert.Equal(t, assignmentList.Assignment.Title, "The Assignment")
	})
}

func TestAssignmentLists(t *testing.T) {
	t.Run("Returns the Classroom URL for the assignment list", func(t *testing.T) {
		assignmentList := AssignmentList{Classroom: Classroom{Url: "https://classroom.github.com/url-here"}}
		assert.Equal(t, assignmentList.Url(), "https://classroom.github.com/url-here")
	})
}

func TestAssignments(t *testing.T) {
	t.Run("Uses classroom url and slug for url", func(t *testing.T) {
		assignment := Assignment{Classroom: Classroom{Url: "https://classroom.github.com/url-here"}, Slug: "the-slug"}
		assert.Equal(t, assignment.Url(), "https://classroom.github.com/url-here/assignments/the-slug")
	})

	t.Run("returns if assignment is a group assignment", func(t *testing.T) {
		assignment := Assignment{AssignmentType: "individual"}
		assert.False(t, assignment.IsGroupAssignment())

		assignment = Assignment{AssignmentType: "group"}
		assert.True(t, assignment.IsGroupAssignment())
	})
}

func TestAcceptedAssignments(t *testing.T) {
	t.Run("Uses repository url for RepositoryUrl()", func(t *testing.T) {
		acceptedAssignment := AcceptedAssignment{Repository: GithubRepository{HtmlUrl: "https://github.com/owner/repo"}}
		assert.Equal(t, acceptedAssignment.RepositoryUrl(), "https://github.com/owner/repo")
	})
}