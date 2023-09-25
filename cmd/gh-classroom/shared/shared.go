package shared

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh/pkg/api"
	"github.com/github/gh-classroom/pkg/classroom"
)

func PromptForClassroom(client api.RESTClient) (classroomId classroom.Classroom, err error) {
	if err != nil {
		return classroom.Classroom{}, err
	}

	classrooms, err := classroom.ListClassrooms(client, 1, 100)

	if err != nil {
		return classroom.Classroom{}, err
	}

	if len(classrooms) == 0 {
		return classroom.Classroom{}, errors.New("No classrooms found.")
	}

	optionMap := make(map[string]classroom.Classroom)
	options := make([]string, 0, len(classrooms))

	for _, classroom := range classrooms {
		optionMap[classroom.Name] = classroom
		options = append(options, classroom.Name)
	}

	var qs = []*survey.Question{
		{
			Name: "classroom",
			Prompt: &survey.Select{
				Message: "Select a classroom:",
				Options: options,
			},
		},
	}

	answer := struct {
		Classroom string
	}{}

	err = survey.Ask(qs, &answer)

	if err != nil {
		return classroom.Classroom{}, err
	}

	return optionMap[answer.Classroom], nil
}

func PromptForAssignment(client api.RESTClient, classroomId int) (assignment classroom.Assignment, err error) {
	assignmentList, err := classroom.ListAssignments(client, classroomId, 1, 100)
	if err != nil {
		return classroom.Assignment{}, err
	}

	optionMap := make(map[string]classroom.Assignment)
	options := make([]string, 0, len(assignmentList.Assignments))

	for _, assignment := range assignmentList.Assignments {
		optionMap[assignment.Title] = assignment
		options = append(options, assignment.Title)
	}

	if len(options) == 0 {
		return classroom.Assignment{}, errors.New("No assignments found for this classroom.")
	}

	var qs = []*survey.Question{
		{
			Name: "assignment",
			Prompt: &survey.Select{
				Message: "Select an assignment:",
				Options: options,
			},
		},
	}

	answer := struct {
		Assignment string
	}{}

	err = survey.Ask(qs, &answer)

	if err != nil {
		return classroom.Assignment{}, err
	}

	return optionMap[answer.Assignment], nil
}

func ListAcceptedAssignments(client api.RESTClient, assignmentID int, page int, perPage int) (classroom.AcceptedAssignmentList, error) {
	response, err := classroom.GetAssignmentList(client, assignmentID, page, perPage)
	if err != nil {
		return classroom.AcceptedAssignmentList{}, err
	}

	if len(response) == 0 {
		return classroom.AcceptedAssignmentList{}, nil
	}
	acceptedAssignmentList := classroom.NewAcceptedAssignmentList(response)

	return acceptedAssignmentList, nil
}

func ListAllAcceptedAssignments(client api.RESTClient, assignmentID int, perPage int) (classroom.AcceptedAssignmentList, error) {
	var page = 1
	response, err := classroom.GetAssignmentList(client, assignmentID, page, perPage)
	if err != nil {
		return classroom.AcceptedAssignmentList{}, err
	}

	if len(response) == 0 {
		return classroom.AcceptedAssignmentList{}, nil
	}

	//keep calling getAssignmentList until we get them all
	var nextList []classroom.AcceptedAssignment
	for hasNext := true; hasNext; {
		page += 1
		nextList, err = classroom.GetAssignmentList(client, assignmentID, page, perPage)
		if err != nil {
			return classroom.AcceptedAssignmentList{}, err
		}
		hasNext = len(nextList) > 0
		response = append(response, nextList...)
	}

	acceptedAssignmentList := classroom.NewAcceptedAssignmentList(response)

	return acceptedAssignmentList, nil
}
