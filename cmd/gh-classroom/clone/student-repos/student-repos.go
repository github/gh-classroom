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
	"github.com/github/gh-classroom/cmd/gh-classroom/clone/utils"
)

func NewCmdStudentRepo(f *cmdutil.Factory) *cobra.Command {
	var assignmentId int
	var directory string
	var page int
	var perPage int
	var getAll bool
	var verbose bool

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
			
			// Default getAll to true unless page is set differently from default.
			getAll = !cmd.Flags().Changed("page")

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
			cloneErrors := []string{}
			for _, acceptAssignment := range acceptedAssignmentList.AcceptedAssignments {
				clonePath := filepath.Join(fullPath, acceptAssignment.Repository.Name)
				err := utils.CloneRepository(clonePath, acceptAssignment.Repository.FullName, gh.Exec)
				if err != nil {
						errMsg := fmt.Sprintf("Error cloning %s: %v", acceptAssignment.Repository.FullName, err)
						fmt.Println(errMsg)
						cloneErrors = append(cloneErrors, errMsg)
						continue // Continue with the next iteration
				}
				totalCloned++
			}
			if len(cloneErrors) > 0 {
				fmt.Println("Some repositories failed to clone.")
				if !verbose {
						fmt.Println("Run with --verbose flag to see more details")
				} else {
						for _, errMsg := range cloneErrors {
								fmt.Println(errMsg)
						}
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
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose error output")

	return cmd
}
