package grades

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/term"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdAssignmentGrades(f *cmdutil.Factory) *cobra.Command {
	var (
		web               bool
		assignmentID      int
		filename          string
		isGroupAssignment bool
	)

	cmd := &cobra.Command{
		Use:     "assignment-grades",
		Example: `$ gh classroom assignment-grades -a 4876`,
		Short:   "Download a CSV of grades for an assignment in a classroom",
		Run: func(cmd *cobra.Command, args []string) {
			term := term.FromEnv()

			client, err := gh.RESTClient(nil)
			var assignment classroom.Assignment
			if err != nil {
				log.Fatal(err)
			}

			if assignmentID == 0 {
				classroom, err := shared.PromptForClassroom(client)
				classroomID := classroom.Id
				if err != nil {
					log.Fatal(err)
				}

				assignment, err = shared.PromptForAssignment(client, classroomID)
				assignmentID = assignment.Id
				if err != nil {
					log.Fatal(err)
				}
			}

			if web {
				if term.IsTerminalOutput() {
					fmt.Fprintln(cmd.ErrOrStderr(), "Opening in your browser.")
				}
				browser := browser.New("", cmd.OutOrStdout(), cmd.OutOrStderr())
				err := browser.Browse(assignment.Url())
				if err != nil {
					log.Fatal(err)
				}
				return
			}

			response, err := classroom.GetAssignmentGrades(client, assignmentID)
			if err != nil {
				log.Fatal(err)
			}

			if len(response) == 0 {
				log.Fatal("No grades were returned for assignment")
			}

			f, err := os.Create(filename)
			if err != nil {
				log.Fatalln("failed to open file", err)
			}
			defer f.Close()

			w := csv.NewWriter(f)
			defer w.Flush()

			for i, grade := range response {
				if len(grade.GroupName) != 0 {
					isGroupAssignment = true
				}

				if i == 0 {
					err := w.Write(gradeCSVHeaders(isGroupAssignment))
					if err != nil {
						log.Fatalln("error writing header to file", err)
					}
				}

				row := []string{
					grade.AssignmentName,
					grade.AssignmentURL,
					grade.StarterCodeURL,
					grade.GithubUsername,
					grade.RosterIdentifier,
					grade.StudentRepositoryName,
					grade.StudentRepositoryURL,
					grade.SubmissionTimestamp,
					grade.PointsAwarded,
					grade.PointsAvailable,
				}
				if isGroupAssignment {
					row = append(row, grade.GroupName)
				}

				err := w.Write(row)
				if err != nil {
					log.Fatalln("error writing row to file", err)
				}
			}
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open specified assignment in a web browser")
	cmd.Flags().IntVarP(&assignmentID, "assignment-id", "a", 0, "Assignment ID (optional)")
	cmd.Flags().StringVarP(&filename, "file-name", "f", "grades.csv", "File name (optional)")
	return cmd
}

func gradeCSVHeaders(isGroupAssignment bool) []string {
	headers := []string{
		"assignment_name",
		"assignment_url",
		"starter_code_url",
		"github_username",
		"roster_identifier",
		"student_repository_name",
		"student_repository_url",
		"submission_timestamp",
		"points_awarded",
		"points_available",
	}
	if isGroupAssignment {
		headers = append(headers, "group_name")
	}
	return headers
}
