package starter_repo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/cli/go-gh"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdStarterRepo() *cobra.Command {
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

				assignmentId, err = shared.PromptForAssignment(client, classroomId)

				if err != nil {
					log.Fatal(err)
				}

			}

			assignment, err := classroom.GetAssignment(client, assignmentId)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(assignment.Id)

			if assignment.StarterCodeRepository.FullName == "" {
				fmt.Println("No starter code repository found for this assignment.")
				return
			} else {
				fmt.Println(assignment.StarterCodeRepository.FullName)
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

			stdOut, _, err := gh.Exec("repo", "clone", assignment.StarterCodeRepository.FullName, "--", clonePath)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println(stdOut.String())
		},
	}
	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "ID of the assignment")
	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "Directory to clone into")

	return cmd
}
