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

func NewCmdStarterRepo(f *cmdutil.Factory) *cobra.Command {
	var assignmentId int
	var directory string

	cmd := &cobra.Command{
		Use:   "starter-repo",
		Short: "Clone starter code",
		Long: heredoc.Doc(`
		Clones starter code for an assignment into a directory

		By default, the starter code is cloned into the current directory. To clone into a different directory, use the --directory flag.

		If the directory does not exists, it will be created.
		`),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := api.DefaultRESTClient()
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

			fullPath, err := filepath.Abs(directory)

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

			clonePath := fullPath + "/" + assignment.Slug
			fmt.Printf("Cloning into: %v\n", clonePath)

			_, _, err = gh.Exec("repo", "clone", assignment.StarterCodeRepository.FullName, "--", clonePath)
			if err != nil {
				log.Fatal(err)
				return
			}
		},
	}
	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory to clone into")

	return cmd
}
