package starter_repo

import (
	"fmt"
	"log"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func promptForClassroom(client api.RESTClient) (classroomId int, err error) {
	if err != nil {
		return 0, err
	}

	classrooms, err := classroom.ListClassrooms(client, 1, 100)
	if err != nil {
		return 0, err
	}

	optionMap := make(map[string]int)
	options := make([]string, len(classrooms))

	for _, classroom := range classrooms {
		optionMap[classroom.Name] = classroom.Id
		options = append(options, classroom.Name)
	}

	// Prompt user to select a classrooms
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
		return 0, err
	}

	return optionMap[answer.Classroom], nil
}

func promptForAssignment(client api.RESTClient, classroomId int) (assignmentId int, err error) {
	assignmentList, err := classroom.ListAssignments(client, classroomId, 1, 100)
	if err != nil {
		return 0, err
	}

	optionMap := make(map[string]int)
	options := make([]string, len(assignmentList.Assignments))

	for _, assignment := range assignmentList.Assignments {
		optionMap[assignment.Title] = assignment.Id
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
		return 0, err
	}

	return optionMap[answer.Assignment], nil
}

func NewCmdStarterRepo() *cobra.Command {
	var assignmentId int
	var directory string

	cmd := &cobra.Command{
		Use:   "starter-repo",
		Short: "Clone starter code",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			if assignmentId == 0 {
				fmt.Println("Fetching ")

				classroomId, err := promptForClassroom(client)

				if err != nil {
					log.Fatal(err)
				}

				assignmentId, err = promptForAssignment(client, classroomId)

				if err != nil {
					log.Fatal(err)
				}
				fmt.Print("Assignment ID:")
				fmt.Println(assignmentId)
				fmt.Println("WOOOOOO")

				// Fetch list of assignments for that classroom
				// Prompt user to select an assignment
				// Fetch the starter code for that assignment
				// Clone the starter code into the current directory
			} else {
				// Fetch the starter code for that assignment
				// Clone the starter code into the current directory
			}
		},
	}
	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", "", "Directory to clone into")
	return cmd
}
