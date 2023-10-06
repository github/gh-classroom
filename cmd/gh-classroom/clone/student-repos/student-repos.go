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
	var getAll bool

	cmd := &cobra.Command{
		Use:   "student-repos",
		Short: "Clone student repos for an assignment",
		Long: heredoc.Doc(`Clone student repos for an assignment into a directory.

		By default, the student repos are cloned into the current directory in a directory named after the assignment slug.
		To clone into a different directory, use the --directory flag.

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
			var acceptedAssignmentList classroom.AcceptedAssignmentList
			if getAll {
				acceptedAssignmentList, err = shared.ListAllAcceptedAssignments(client, assignmentId, perPage)
			} else {
				acceptedAssignmentList, err = shared.ListAcceptedAssignments(client, assignmentId, page, perPage)
			}

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

			totalCloned := 0
			for _, acceptAssignment := range acceptedAssignmentList.AcceptedAssignments {
				clonePath := filepath.Join(fullPath, acceptAssignment.Repository.Name())
				if _, err := os.Stat(clonePath); os.IsNotExist(err) {
					fmt.Printf("Cloning into: %v\n", clonePath)
					_, _, err := gh.Exec("repo", "clone", acceptAssignment.Repository.FullName, "--", clonePath)
					totalCloned++
					if err != nil {
						log.Fatal(err)
						return
					}
				} else {
					fmt.Printf("Skip existing repo: %v use gh classroom pull to get new commits\n", clonePath)
				}
			}
			if getAll {
				fmt.Printf("Cloned %v repos.\n", totalCloned)
			} else {
				numPages, _ := shared.NumberOfAcceptedAssignmentsAndPages(client, assignmentId, perPage)
				fmt.Printf("Cloned %v repos. There are %v more pages of repos to clone.\n", totalCloned, numPages-page)
			}
		},
	}

	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory to clone into")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 15, "Number of accepted assignments per page")
	cmd.Flags().BoolVar(&getAll, "all", true, "Clone All assignments by default")

	return cmd
}
