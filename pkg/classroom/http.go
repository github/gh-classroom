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

	if len(response) == 0 {
		return AssignmentList{}, nil
	}

	assignmentList := NewAssignmentList(response)

	return assignmentList, nil
}

func ListClassrooms(client api.RESTClient, page int, perPage int) ([]Classroom, error) {
	var response []Classroom

	err := client.Get(fmt.Sprintf("classrooms?page=%v&per_page=%v", page, perPage), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetAssignmentList(client api.RESTClient, assignmentID int, page int, perPage int) ([]AcceptedAssignment, error) {
	var response []AcceptedAssignment

	err := client.Get(fmt.Sprintf("assignments/%v/accepted_assignments?page=%v&per_page=%v", assignmentID, page, perPage), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetAssignment(client api.RESTClient, assignmentID int) (Assignment, error) {
	var response Assignment
	err := client.Get(fmt.Sprintf("assignments/%v", assignmentID), &response)
	if err != nil {
		return Assignment{}, err
	}
	return response, nil
}

func GetAssignmentGrades(client api.RESTClient, assignmentID int) ([]AssignmentGrade, error) {
	var response []AssignmentGrade
	err := client.Get(fmt.Sprintf("assignments/%v/grades", assignmentID), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func GetClassroom(client api.RESTClient, classroomID int) (Classroom, error) {
	var response Classroom

	err := client.Get(fmt.Sprintf("classrooms/%v", classroomID), &response)
	if err != nil {
		return Classroom{}, err
	}

	return response, nil
}
