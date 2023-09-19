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

func getAssignmentList(client api.RESTClient, assignmentID int, page int, perPage int) ([]AcceptedAssignment, error) {
	var response []AcceptedAssignment
	err := client.Get(fmt.Sprintf("assignments/%v/accepted_assignments?page=%v&per_page=%v", assignmentID, page, perPage), &response)
	return response, err
}

func ListAcceptedAssignments(client api.RESTClient, assignmentID int, page int, perPage int) (AcceptedAssignmentList, error) {
	response, err := getAssignmentList(client, assignmentID, page, perPage)
	if err != nil {
		return AcceptedAssignmentList{}, err
	}

	if len(response) == 0 {
		return AcceptedAssignmentList{}, nil
	}
	acceptedAssignmentList := NewAcceptedAssignmentList(response)

	return acceptedAssignmentList, nil
}

func ListAllAcceptedAssignments(client api.RESTClient, assignmentID int, perPage int) (AcceptedAssignmentList, error) {
	var page = 1
	response, err := getAssignmentList(client, assignmentID, page, perPage)
	if err != nil {
		return AcceptedAssignmentList{}, err
	}

	if len(response) == 0 {
		return AcceptedAssignmentList{}, nil
	}

	//keep calling getAssignmentList until we get them all
	var nextList []AcceptedAssignment
	for hasNext := true; hasNext; {
		page += 1
		nextList, err = getAssignmentList(client, assignmentID, page, perPage)
		if err != nil {
			return AcceptedAssignmentList{}, err
		}
		hasNext = len(nextList) > 0
		response = append(response, nextList...)
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
