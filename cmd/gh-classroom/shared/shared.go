package shared

import (
	"errors"
	"log"
	"math"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/github/gh-classroom/pkg/classroom"
)

func PromptForClassroom(client *api.RESTClient) (classroomId classroom.Classroom, err error) {
	if err != nil {
		return classroom.Classroom{}, err
	}

	classrooms, err := classroom.ListClassrooms(client, 1, 100)

	if err != nil {
		return classroom.Classroom{}, err
	}

	if len(classrooms) == 0 {
		return classroom.Classroom{}, errors.New("no classrooms found")
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

func PromptForAssignment(client *api.RESTClient, classroomId int) (assignment classroom.Assignment, err error) {
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
		return classroom.Assignment{}, errors.New("no assignments found for this classroom")
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

// Get the total number of accepted assignments and the number of pages necessary to get them all
// Calculate the number of pages necessary to get all of the assignments with a given perPage
func NumberOfAcceptedAssignmentsAndPages(client *api.RESTClient, assignmentID int, perPage int) (numPages, totalAccepted int) {
	assignment, err := classroom.GetAssignment(client, assignmentID)
	if err != nil {
		log.Fatal(err)
	}
	numPages = int(math.Ceil(float64(assignment.Accepted) / float64(perPage)))
	totalAccepted = assignment.Accepted
	return
}

func ListAcceptedAssignments(client *api.RESTClient, assignmentID int, page int, perPage int) (classroom.AcceptedAssignmentList, error) {
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

func ListAllAcceptedAssignments(client *api.RESTClient, assignmentID int, perPage int) (classroom.AcceptedAssignmentList, error) {

	numPages, totalAccepted := NumberOfAcceptedAssignmentsAndPages(client, assignmentID, perPage)

	ch := make(chan assignmentList)
	var wg sync.WaitGroup
	for page := 1; page <= numPages; page++ {
		wg.Add(1)
		go func(pg int) {
			defer wg.Done()
			response, err := classroom.GetAssignmentList(client, assignmentID, pg, perPage)
			ch <- assignmentList{
				assignments: response,
				Error:       err,
			}
		}(page)
	}

	var mu sync.Mutex
	assignments := make([]classroom.AcceptedAssignment, 0, totalAccepted)
	var hadErr error = nil
	for page := 1; page <= numPages; page++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := <-ch
			if result.Error != nil {
				hadErr = result.Error
			} else {
				mu.Lock()
				assignments = append(assignments, result.assignments...)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	close(ch)

	if hadErr != nil {
		return classroom.AcceptedAssignmentList{}, hadErr
	}

	return classroom.NewAcceptedAssignmentList(assignments), nil
}
