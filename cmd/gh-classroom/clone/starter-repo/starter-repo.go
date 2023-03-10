package starter_repo

import (
	"fmt"
	"log"

	"github.com/cli/go-gh"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/spf13/cobra"
)


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
				classroomId, err := shared.PromptForClassroom(client)

				if err != nil {
					log.Fatal(err)
				}

				assignmentId, err = shared.PromptForAssignment(client, classroomId)

				if err != nil {
					log.Fatal(err)
				}

			}
			assignment, err := classroom.GetAssignment(client, assignmentId)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(assignment.StarterCodeRepository.FullName)

			if directory == "" {
				// verify directory exists
				// verify directory is empty
				// clone into directory

			}
		},
	}
	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory to clone into")
	return cmd
}
