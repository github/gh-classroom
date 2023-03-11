package accepted

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cli/cli/pkg/iostreams"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/tableprinter"
	"github.com/cli/go-gh/pkg/term"
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
			term := term.FromEnv()
			io := iostreams.System()
			cs := io.ColorScheme()

			client, err := gh.RESTClient(nil)
			var assignment classroom.Assignment
			if err != nil {
				log.Fatal(err)
			}

			if assignmentId == 0 {
				classroom, err := shared.PromptForClassroom(client)
				classroomId := classroom.Id
				if err != nil {
					log.Fatal(err)
				}

				assignment, err = shared.PromptForAssignment(client, classroomId)
				assignmentId = assignment.Id
				if err != nil {
					log.Fatal(err)
				}
			}

			if web {
				if term.IsTerminalOutput() {
					fmt.Fprintln(cmd.ErrOrStderr(), "Opening in your browser.")
				}
				browser := browser.New("", cmd.OutOrStdout(), cmd.OutOrStderr())
				browser.Browse(assignment.Url())
				return
			}
			acceptedAssignments, err := classroom.ListAcceptedAssignments(client, assignmentId, page, perPage)
			if err != nil {
				log.Fatal(err)
			}

			t := tableprinter.New(cmd.OutOrStdout(), term.IsTerminalOutput(), 14)
			t.AddField("ID", tableprinter.WithTruncate(nil))
			t.AddField("Submitted", tableprinter.WithTruncate(nil))
			t.AddField("Passing", tableprinter.WithTruncate(nil))
			t.AddField("Commit Count", tableprinter.WithTruncate(nil))
			t.AddField("Grade", tableprinter.WithTruncate(nil))
			t.AddField("Feedback Pull Request URL", tableprinter.WithTruncate(nil))
			if assignment.IsGroupAssignment() {
				t.AddField("Group Members", tableprinter.WithTruncate(nil))
			} else {
				t.AddField("Student", tableprinter.WithTruncate(nil))
			}
			t.AddField("Repository", tableprinter.WithTruncate(nil))
			t.EndRow()

			for _, acceptedAssignment := range acceptedAssignments.AcceptedAssignments {
				var students []string
				for _, student := range acceptedAssignment.Students {
					students = append(students, student.Login)
				}
				t.AddField(cs.Green(strconv.Itoa(acceptedAssignment.Id)), tableprinter.WithTruncate(nil))
				t.AddField(strconv.FormatBool(acceptedAssignment.Submitted), tableprinter.WithTruncate(nil))
				t.AddField(strconv.FormatBool(acceptedAssignment.Passing), tableprinter.WithTruncate(nil))
				t.AddField(strconv.Itoa(acceptedAssignment.CommitCount), tableprinter.WithTruncate(nil))
				t.AddField(acceptedAssignment.Grade, tableprinter.WithTruncate(nil))
				t.AddField(acceptedAssignment.FeedbackPullRequestUrl, tableprinter.WithTruncate(nil))
				t.AddField(strings.Join(students, ", "), tableprinter.WithTruncate(nil))
				t.AddField(acceptedAssignment.RepositoryUrl(), tableprinter.WithTruncate(nil))
				t.EndRow()
			}
			t.Render()
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open the assignment in a browser")
	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 30, "Number of accepted assignments per page")
	return cmd
}
