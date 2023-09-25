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

type assignmentList struct {
	assignments []classroom.AcceptedAssignment
	Error       error
}

func ListAllAcceptedAssignments(client api.RESTClient, assignmentID int, perPage int) (classroom.AcceptedAssignmentList, error) {

	//Calculate the number of go channels to create. We will assign 1 channel per page
	//so we don't put too much pressure on the API all at once
	theAssignment, err := classroom.GetAssignment(client, assignmentID)
	if err != nil {
		return classroom.AcceptedAssignmentList{}, err
	}
	numChannels := (theAssignment.Accepted / perPage)
	//See if we need one more page. This is probably always going to be true
	//But in the rare case it is not we can save one API call :)
	if theAssignment.Accepted%perPage > 0 {
		numChannels++
	}
	ch := make(chan assignmentList)
	for page := 1; page <= numChannels; page += 1 {
		go func(pg int) {
			response, err := classroom.GetAssignmentList(client, assignmentID, pg, perPage)
			ch <- assignmentList{
				assignments: response,
				Error:       err,
			}
		}(page)
	}
	assignments := make([]classroom.AcceptedAssignment, 0, theAssignment.Accepted)
	var hadErr error = nil
	for page := 1; page <= numChannels; page += 1 {
		result := <-ch
		if result.Error != nil {
			hadErr = result.Error
		}
		assignments = append(assignments, result.assignments...)
	}

	if hadErr != nil {
		return classroom.AcceptedAssignmentList{}, err
	}

	return classroom.NewAcceptedAssignmentList(assignments), nil
}
