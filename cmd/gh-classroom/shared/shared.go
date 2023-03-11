package shared

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh/pkg/api"
	"github.com/github/gh-classroom/pkg/classroom"
)

func PromptForClassroom(client api.RESTClient) (classroomId classroom.ShortClassroom, err error) {
	if err != nil {
		return classroom.ShortClassroom{}, err
	}

	classrooms, err := classroom.ListClassrooms(client, 1, 100)
	if err != nil {
		return classroom.ShortClassroom{}, err
	}

	optionMap := make(map[string]classroom.ShortClassroom)
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
		return classroom.ShortClassroom{}, err
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
