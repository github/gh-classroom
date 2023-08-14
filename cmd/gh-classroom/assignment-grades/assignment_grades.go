package assignmentgrades

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdAssignments(f *cmdutil.Factory) *cobra.Command {
	var (
		web          bool
		assignmentID int
		filename     string
	)

	cmd := &cobra.Command{
		Use:     "assignment-grades",
		Example: `$ gh classroom assignment-grades -a 4876`,
		Short:   "Download a CSV of grades for an assignment in a classroom",
		Long:    "Download a CSV of grades for an assignment in a classroom",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			if assignmentID == 0 {
				cr, err := shared.PromptForClassroom(client)
				classroomID := cr.Id

				if err != nil {
					log.Fatal(err)
				}

				assignment, err := shared.PromptForAssignment(client, classroomID)
				if err != nil {
					log.Fatal(err)
				}
				assignmentID = assignment.Id
			}

			response, err := classroom.GetAssignment(client, assignmentID)
			if err != nil {
				log.Fatal(err)
			}
			grades := response.Grades
			if len(grades) == 0 {
				log.Fatal("No grades were returned for assignment")
			}

			f, err := os.Create(filename)
			if err != nil {
				log.Fatalln("failed to open file", err)
			}
			defer f.Close()

			w := csv.NewWriter(f)
			err = w.WriteAll(grades)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.Flags().BoolVar(&web, "web", false, "Open specified assignment in a web browser")
	cmd.Flags().IntVarP(&assignmentID, "assignment-id", "a", 0, "Assignment ID (optional)")
	cmd.Flags().StringVarP(&filename, "file-name", "f", "grades.csv", "File name (optional)")

	return cmd
}
