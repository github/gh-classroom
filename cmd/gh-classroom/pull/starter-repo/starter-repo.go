package starter_repo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdStarterRepoPull(f *cmdutil.Factory) *cobra.Command {
	var assignmentId int
	var directory string

	cmd := &cobra.Command{
		Use:   "starter-repo",
		Short: "Run pull on an existing starter repository",
		Long: heredoc.Doc(`Given a starter repository that was previously cloned run a pull to get any new commits.
		`),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := api.DefaultRESTClient()
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

			if assignment.StarterCodeRepository.FullName == "" {
				fmt.Println("No starter code repository found for this assignment.")
				return
			}

			if strings.HasPrefix(directory, "~") {
				dirname, _ := os.UserHomeDir()
				directory = filepath.Join(dirname, directory[1:])
			}

			fullPath, err := filepath.Abs(filepath.Join(directory, assignment.Slug))
			if err != nil {
				log.Fatal(err)
			}

			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				fmt.Println("Starter code don't exist run gh classroom clone starter-code first")
				return
			}
			fmt.Println("Pulling: ", fullPath)
			err = os.Chdir(fullPath)
			if err != nil {
				log.Fatal(err)
			}

			_, se, err := gh.Exec("repo", "sync")
			if err != nil {
				fmt.Println(se.String())
				return
			}
		},
	}

	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory to that the repos live in")

	return cmd
}
