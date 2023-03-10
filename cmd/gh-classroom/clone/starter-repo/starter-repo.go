package starter_repo

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func NewCmdStarterRepo() *cobra.Command {
	var assignmentID int
	var directory string

	cmd := &cobra.Command{
		Use:   "starter-repo",
		Short: "Clone starter code",
		Run: func(cmd *cobra.Command, args []string) {
			if assignmentID == 0 {
				fmt.Println("Fetching ")
				// Fetch list of classrooms
				// Prompt user to select a classrooms
				var qs = []*survey.Question{
					{
						Name: "classroom",
						Prompt: &survey.Select{
							Message: "Select a classroom:",
							Options: []string{"Classroom 1", "Classroom 2", "Classroom 3"},
						},
					},
				}
				answers := struct {
					Classroom string // survey will match the question and field names
				}{}

				err := survey.Ask(qs, &answers)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
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
	cmd.Flags().IntVarP(&assignmentID, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", "", "Directory to clone into")
	return cmd
}
