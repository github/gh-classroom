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

func NewCmdStudentReposPull(f *cmdutil.Factory) *cobra.Command {
	var assignmentId int
	var directory string

	cmd := &cobra.Command{
		Use:   "student-repos",
		Short: "Run pull on an existing set of student repositories",
		Long: heredoc.Doc(`Given a previously cloned set of student repositories run a pull on each one to get any new commits.
		`),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			if assignmentId == 0 {
				classroom, err := shared.PromptForClassroom(client)

				if err != nil {
					log.Fatal(err)
				}
				classroomId := classroom.Id

				assignment, err := shared.PromptForAssignment(client, classroomId)
				if err != nil {
					log.Fatal(err)
				}
				assignmentId = assignment.Id
			}

			assignment, err := classroom.GetAssignment(client, assignmentId)

			if err != nil {
				log.Fatal(err)
			}

			if strings.HasPrefix(directory, "~") {
				dirname, _ := os.UserHomeDir()
				directory = filepath.Join(dirname, directory[1:])
			}

			fullPath, err := filepath.Abs(filepath.Join(directory, assignment.Slug+"-submissions"))
			if err != nil {
				log.Fatal(err)
			}
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				fmt.Println("Directories don't exist run gh classroom clone student-repos first")
				return
			}

			//Save off the cwd so we can restore when we run gh sync
			baseDir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			entries, err := os.ReadDir(fullPath)
			if err != nil {
				log.Fatal(err)
			}

			for _, r := range entries {
				if !r.IsDir() {
					continue
				}
				clonePath := filepath.Join(fullPath, r.Name)
				fmt.Printf("Pulling repo: %v\n", clonePath)
				err = os.Chdir(clonePath)
				if err != nil {
					log.Fatal(err)
				}
				_, _, err := gh.Exec("repo", "sync")
				if err != nil {
					//Don't bail on an error the repo could have changes preventing
					//a pull, continue with rest of repos
					fmt.Println(err)
				}
				err = os.Chdir(baseDir)
				if err != nil {
					log.Fatal(err)
				}
			}

		},
	}

	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory to that the repos live in")

	return cmd
}
