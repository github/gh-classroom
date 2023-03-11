package assignment

import (
	"log"
	"fmt"
	"strconv"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/term"
	"github.com/github/gh-classroom/pkg/classroom"
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
	//Classroom                   Classroom `json:"classroom"`
}

func NewCmdAssignment(f *cmdutil.Factory) *cobra.Command {
	var (
		web          bool
		assignmentId int
	)

	cmd := &cobra.Command{
		Use:   "assignment",
		Short: "Open an assignment",
		Long:  "Open an specific assignment for a classroom",
		Run: func(cmd *cobra.Command, args []string) {
			c := iostreams.NewColorScheme(true, true, true)

			fmt.Println(c.Bold(c.Blue("Show Assignment")))

			client, err := gh.RESTClient(nil)
			response, err := classroom.GetAssignment(client, assignmentId)

			if assignmentId == 0 {
				log.Fatal("Assignment ID is required")
			}

			if web {
				if term.isTerminalOutput() { fmt.Println(io.ErrOut, "Opening assignment in a web browser")
				}
				// figure out how to format the url to open in the browser
				browser := browser.New("")
			}

			client, err := gh.RESTClient(nil)
			response, err := classroom.GetAssignment(client, assignmentId)

			if err != nil {
				log.Fatal(err)
			}

			if web {
				// // if term.isTerminalOutput() {
				// // 	fmt.Println(io.ErrOut, "Opening assignment in a web browser")
				// // }
				// // figure out how to format the url to open in the browser
				// browser := browser.New("")
			}

			fmt.Println()
			PrintAssigment(response)
			fmt.Println()
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open specified assignment in a web browser")
	cmd.Flags().IntVar(&assignmentId, "assignment-id", 0, "Assignment ID (required)")

	return cmd
}

func PrintAssigment(response classroom.Assignment) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Println(c.Blue("ASSIGNMENT INFORMATION"))
	fmt.Println(c.Yellow("ID:"), c.Green(strconv.Itoa(response.Id)))
	fmt.Println(c.Yellow("Title:"), c.Green(response.Title))
	fmt.Println(c.Yellow("Invite Link:"), c.Green(response.InviteLink))
	fmt.Println(c.Yellow("Accepted:"), c.Green(strconv.Itoa(response.Accepted)))
	fmt.Println(c.Yellow("Submissions:"), c.Green(strconv.Itoa(response.Submissions)))
	fmt.Println(c.Yellow("Passing:"), c.Green(strconv.Itoa(response.Passing)))
}

