package assignments

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/tableprinter"
	"github.com/cli/go-gh/pkg/term"
	"github.com/cli/go-gh/pkg/text"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

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
			term := term.FromEnv()
			io := iostreams.System()
			cs := io.ColorScheme()

			if classroomID == 0 {
				log.Fatal("Missing classroom ID")
			}
			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			var response []classroom.Assignment
			err = client.Get(fmt.Sprintf("classrooms/%v/assignments", classroomID), &response)
			if err != nil {
				log.Fatal(err)
			}

			assignmentList := classroom.NewAssignmentList(response)

			fmt.Println(assignmentListSummary(assignmentList, cs))

			t := tableprinter.New(os.Stdout, term.IsTerminalOutput(), 14)

			if web {
				if term.IsTerminalOutput() {
					fmt.Fprintln(io.ErrOut, "Opening in your browser.")
				}
				browser := browser.New("", io.Out, io.ErrOut)
				browser.Browse(assignmentList.Url())
				return
			}

			for _, assignment := range response {
				t.AddField(cs.Green(strconv.Itoa(assignment.Id)), tableprinter.WithTruncate(nil))
				t.AddField(assignment.Title, tableprinter.WithTruncate(nil))
				t.AddField(cs.Gray(strconv.FormatBool(assignment.PublicRepo)), tableprinter.WithTruncate(nil))
				t.EndRow()
			}
			t.Render()
		},
	}
	cmd.Flags().BoolVar(&web, "web", false, "Open the assignment list in a browser")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 30, "Number of assignments per page")
	cmd.Flags().IntVarP(&classroomID, "classroom-id", "c", 0, "ID of the classroom")

	return cmd
}

func colorForVisibility(public bool) string {
	if public {
		return "green"
	}
	return "red"
}

func assignmentListSummary(a classroom.AssignmentList, cs *iostreams.ColorScheme) string {
	switch a.Count {
	case 0:
		return fmt.Sprintf("No assignments for %v\n", cs.Blue(a.Classroom.Name))
	case 1:
		return fmt.Sprintf("%v for %v (id: %v)\n", text.Pluralize(a.Count, "Assignment"), cs.Blue(a.Classroom.Name))
	default:
		return fmt.Sprintf("%v for %v\n", text.Pluralize(a.Count, "Assignment"), cs.Blue(a.Classroom.Name))
	}
}
