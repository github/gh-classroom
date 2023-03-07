package clone

import (
	starter_repo "github.com/github/gh-classroom/cmd/gh-classroom/clone/starter-repo"
	student_repos "github.com/github/gh-classroom/cmd/gh-classroom/clone/student-repos"
	"github.com/spf13/cobra"
)

func NewCmdClone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone <starter-code|student-repos>",
		Short: "Clone starter code or a student's submissions",
	}

	cmd.AddCommand(starter_repo.NewCmdStarterRepo())
	cmd.AddCommand(student_repos.NewCmdStudentRepo())
	return cmd
}
