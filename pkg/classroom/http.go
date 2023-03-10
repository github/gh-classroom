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

func ListClassrooms(client api.RESTClient, page int, perPage int) ([]ShortClassroom, error) {
	var response []ShortClassroom

	err := client.Get(fmt.Sprintf("classrooms?page=%v&per_page=%v", page, perPage), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func ListAcceptedAssignments(client api.RESTClient, assignmentID int, page int, perPage int) (AcceptedAssignmentList, error) {
	var response []AcceptedAssignment
	err := client.Get(fmt.Sprintf("assignments/%v/accepted_assignments?page=%v&per_page=%v", assignmentID, page, perPage), &response)
	if err != nil {
		return AcceptedAssignmentList{}, err
	}

	acceptedAssignmentList := NewAcceptedAssignmentList(response)

	return acceptedAssignmentList, nil
}

func GetAssignment(client api.RESTClient, assignmentID int) (Assignment, error) {
	var response Assignment
	err := client.Get(fmt.Sprintf("assignments/%v", assignmentID), &response)
	if err != nil {
		return Assignment{}, err
	}

	return response, nil
}

func GetClassroom(client api.RESTClient, classroomID int) (LongClassroom, error) {
	var response LongClassroom

	err := client.Get(fmt.Sprintf("classrooms/%v", classroomID), &response)
	if err != nil {
		return LongClassroom{}, err
	}

	return response, nil
}
