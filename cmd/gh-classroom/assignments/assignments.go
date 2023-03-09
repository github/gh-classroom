package assignments

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/tableprinter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

type Assignment struct {
	Id                          int            `json:"id" :"id"`
	PublicRepo                  bool           `json:"public_repo" :"public___repo"`
	Title                       string         `json:"title" :"title"`
	AssignmentType              string         `json:"type" :"assignment___type"`
	InviteLink                  string         `json:"invite_link" :"invite___link"`
	InvitationsEnabled          bool           `json:"invitations_enabled" :"invitations___enabled"`
	Slug                        string         `json:"slug" :"slug"`
	StudentsAreRepoAdmins       bool           `json:"students_are_repo_admins" :"students___are___repo___admins"`
	FeedbackPullRequestsEnabled bool           `json:"feedback_pull_requests_enabled" :"feedback___pull___requests___enabled"`
	MaxTeams                    int            `json:"max_teams" :"max___teams"`
	MaxMembers                  int            `json:"max_members" :"max___members"`
	Editor                      string         `json:"editor" :"editor"`
	Accepted                    int            `json:"accepted" :"accepted"`
	Submissions                 int            `json:"submissions" :"submissions"`
	Passing                     int            `json:"passing" :"passing"`
	Language                    string         `json:"language" :"language"`
	Classroom                   ShortClassroom `json:"classroom" :"classroom"`
}

type ShortClassroom struct {
	Id       int    `json:"id" :"id"`
	Name     string `json:"name" :"name"`
	archived bool   `json:"archived" :"archived"`
	url      string `json:"url" :"url"`
}

func NewCmdAssignments() *cobra.Command {
	var web bool
	var page int
	var perPage int
	var classroomID int

	cmd := &cobra.Command{
		Use:   "assignments",
		Short: "Display a list of assignments for a classroom.",
		Long:  "Display a list of assignments for a classroom.",
		Run: func(cmd *cobra.Command, args []string) {
			if classroomID == 0 {
				log.Fatal("Missing classroom ID")
			}
			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			var response []Assignment
			err = client.Get(fmt.Sprintf("classrooms/%v/assignments", classroomID), &response)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(classroomID)
			t := tableprinter.New(os.Stdout, true, 14)
			t.AddField("ID", tableprinter.WithTruncate(nil))
			t.AddField("Title", tableprinter.WithTruncate(nil))
			t.AddField("Repo Visibility", tableprinter.WithTruncate(nil))
			t.EndRow()
			for _, assignment := range response {
				t.AddField(strconv.Itoa(assignment.Id), tableprinter.WithTruncate(nil))
				t.AddField(assignment.Title, tableprinter.WithTruncate(nil))
				t.AddField(strconv.FormatBool(assignment.PublicRepo))
				t.EndRow()
			}
			t.Render()
		},
	}
	cmd.Flags().BoolVarP(&web, "web", "", false, "Open the assignment list in a browser")
	cmd.Flags().IntVarP(&page, "page", "", 1, "Page number")
	cmd.Flags().IntVarP(&perPage, "per-page", "", 30, "Number of assignments per page")
	cmd.Flags().IntVarP(&classroomID, "classroom-id", "", 0, "ID of the classroom")

	return cmd
}
