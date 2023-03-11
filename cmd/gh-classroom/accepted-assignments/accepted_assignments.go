package accepted

import (
	"fmt"
	"log"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdAcceptedAssignments(f *cmdutil.Factory) *cobra.Command {
	var web bool
	var assignmentId int
	var page int
	var perPage int

	cmd := &cobra.Command{
		Use:   "accepted-assignments",
		Short: "List your student's accepted assignments",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			if assignmentId == 0 {
				classroom, err := shared.PromptForClassroom(client)
				classroomId := classroom.Id
				if err != nil {
					log.Fatal(err)
				}

				assignment, err := shared.PromptForAssignment(client, classroomId)
				assignmentId = assignment.Id
				if err != nil {
					log.Fatal(err)
				}
			}

			acceptedAssignments, err := classroom.ListAcceptedAssignments(client, assignmentId, page, perPage)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(acceptedAssignments)
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open the assignment in a browser")
	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 30, "Number of accepted assignments per page")
	return cmd
}
