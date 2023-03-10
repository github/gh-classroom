package assignment

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"

	"github.com/cli/go-gh"
)


type Assignment struct {
	Id                          int    `json:"id"`
	PublicRepo                  bool   `json:"public_repo"`
	Title                       string `json:"title"`
	AssignmentType              string `json:"type"`
	InviteLink                  string `json:"invite_link"`
	InvitationsEnabled          bool   `json:"invitations_enabled"`
	Slug                        string `json:"slug"`
	StudentsAreRepoAdmins       bool   `json:"students_are_repo_admins"`
	FeedbackPullRequestsEnabled bool   `json:"feedback_pull_requests_enabled"`
	MaxTeams                    int    `json:"max_teams"`
	MaxMembers                  int    `json:"max_members"`
	Editor                      string `json:"editor"`
	Accepted                    int    `json:"accepted"`
	Submissions                 int    `json:"submissions"`
	Passing                     int    `json:"passing"`
	Language                    string `json:"language"`
	//Classroom                   ShortClassroom `json:"classroom"`
}

func NewCmdAssignment(f *cmdutil.Factory) *cobra.Command {
	var (
		web          bool
		assignmentID int
		response     Assignment
	)

	cmd := &cobra.Command{
		Use:   "assignment",
		Short: "Open an assignment",
		Long:  "Open an specific assignment for a classroom",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Show Assignment")

			if assignmentID == 0 {
				log.Fatal("No assignment ID provided")
			}

			client, err := gh.RESTClient(nil)

			if err != nil {
				log.Fatal(err)
			}

			err = client.Get(fmt.Sprintf("assignments/%d", assignmentID), &response)

			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open specified assignment in a web browser")
	cmd.Flags().IntVar(&assignmentID, "assignment-id", 0, "Assignment ID (required)")

	return cmd
}
