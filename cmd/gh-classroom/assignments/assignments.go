package assignments

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/tableprinter"
	"github.com/cli/go-gh/pkg/term"
	"github.com/cli/go-gh/pkg/text"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdAssignments() *cobra.Command {
	var web bool
	var page int
	var perPage int
	var classroomId int

	cmd := &cobra.Command{
		Use:   "assignments",
		Short: "Display a list of assignments for a classroom.",
		Long:  "Display a list of assignments for a classroom.",
		Run: func(cmd *cobra.Command, args []string) {
			term := term.FromEnv()
			io := iostreams.System()
			cs := io.ColorScheme()

			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			if classroomId == 0 {
				classroomId, err = shared.PromptForClassroom(client)

				if err != nil {
					log.Fatal(err)
				}
			}


			assignmentList, err := classroom.ListAssignments(client, classroomId, page, perPage)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintln(cmd.OutOrStderr(), assignmentListSummary(assignmentList, cs))

			t := tableprinter.New(cmd.OutOrStdout(), term.IsTerminalOutput(), 14)

			if web {
				if term.IsTerminalOutput() {
					fmt.Fprintln(io.ErrOut, "Opening in your browser.")
				}
				browser := browser.New("", cmd.OutOrStdout(), cmd.OutOrStderr())
				browser.Browse(assignmentList.Url())
				return
			}
			t.AddField("ID", tableprinter.WithTruncate(nil))
			t.AddField("Title", tableprinter.WithTruncate(nil))
			t.AddField("Submission Public", tableprinter.WithTruncate(nil))
			t.AddField("Type", tableprinter.WithTruncate(nil))
			t.AddField("Editor", tableprinter.WithTruncate(nil))
			t.AddField("Invitation Link", tableprinter.WithTruncate(nil))
			t.AddField("Accepted", tableprinter.WithTruncate(nil))
			t.AddField("Submissions", tableprinter.WithTruncate(nil))
			t.AddField("Passing", tableprinter.WithTruncate(nil))
			t.EndRow()

			for _, assignment := range assignmentList.Assignments {
				t.AddField(cs.Green(strconv.Itoa(assignment.Id)), tableprinter.WithTruncate(nil))
				t.AddField(assignment.Title, tableprinter.WithTruncate(nil))
				t.AddField(cs.Gray(strconv.FormatBool(assignment.PublicRepo)), tableprinter.WithTruncate(nil))
				t.AddField(assignment.AssignmentType, tableprinter.WithTruncate(nil))
				t.AddField(assignment.Editor, tableprinter.WithTruncate(nil))
				t.AddField(assignment.InviteLink, tableprinter.WithTruncate(nil))
				t.AddField(strconv.Itoa(assignment.Accepted), tableprinter.WithTruncate(nil))
				t.AddField(strconv.Itoa(assignment.Submissions), tableprinter.WithTruncate(nil))
				t.AddField(strconv.Itoa(assignment.Passing), tableprinter.WithTruncate(nil))
				t.EndRow()
			}
			t.Render()
		},
	}
	cmd.Flags().BoolVar(&web, "web", false, "Open the assignment list in a browser")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 30, "Number of assignments per page")
	cmd.Flags().IntVarP(&classroomId, "classroom-id", "c", 0, "ID of the classroom")

	return cmd
}

func assignmentListSummary(a classroom.AssignmentList, cs *iostreams.ColorScheme) string {
	if a.Count == 0 {
		return fmt.Sprintf("No assignments for %v\n", cs.Blue(a.Classroom.Name))
	} else {
		return fmt.Sprintf("%v for %v\n", text.Pluralize(a.Count, "Assignment"), cs.Blue(a.Classroom.Name))
	}
}
