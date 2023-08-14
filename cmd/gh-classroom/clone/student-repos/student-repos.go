package student_repos

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdStudentRepo(f *cmdutil.Factory) *cobra.Command {
	var assignmentId int
	var directory string
	var page int
	var perPage int

	cmd := &cobra.Command{
		Use:   "student-repos",
		Short: "Clone student repos for an assignment",
		Long: heredoc.Doc(`Clone student repos for an assignment into a directory.

		By default, the student repos are cloned into the current directory a directory named after the assignment slug. To clone into a different directory, use the --directory flag.

		If the directory does not exists, it will be created.
		`),
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

			acceptedAssignmentList, err := classroom.ListAcceptedAssignments(client, assignmentId, page, perPage)

			if err != nil {
				log.Fatal(err)
			}

			assignment := acceptedAssignmentList.Assignment

			if strings.HasPrefix(directory, "~") {
				dirname, _ := os.UserHomeDir()
				directory = filepath.Join(dirname, directory[1:])
			}

			fullPath, err := filepath.Abs(filepath.Join(directory, assignment.Slug+"-submissions"))

			if err != nil {
				fmt.Println("Error getting absolute path for directory: ", err)
				return
			}

			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				fmt.Println("Creating directory: ", fullPath)
				err = os.MkdirAll(fullPath, 0755)

				if err != nil {
					log.Fatal(err)
					return
				}
			}

			for _, acceptAssignment := range acceptedAssignmentList.AcceptedAssignments {
				clonePath := filepath.Join(fullPath, acceptAssignment.Repository.Name())
				fmt.Printf("Cloning into: %v\n", clonePath)
				_, _, err = gh.Exec("repo", "clone", acceptAssignment.Repository.FullName, "--", clonePath)
				if err != nil {
					log.Fatal(err)
					return
				}
			}
		},
	}

	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory to clone into")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 30, "Number of accepted assignments per page")

	return cmd
}
