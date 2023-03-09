package classroom

import (
	"fmt"

	"github.com/cli/go-gh/pkg/api"
)

func ListAssignments(client api.RESTClient, classroomID int, page int, perPage int) (AssignmentList, error) {
	var response []Assignment
	err := client.Get(fmt.Sprintf("classrooms/%v/assignments?page=%v&per_page=%v", classroomID, page, perPage), &response)
	if err != nil {
		return AssignmentList{}, err
	}

	assignmentList := NewAssignmentList(response)

	return assignmentList, nil
}
