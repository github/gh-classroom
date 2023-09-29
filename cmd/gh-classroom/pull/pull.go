package pull

import (
	"github.com/cli/cli/v2/pkg/cmdutil"
	starter_repo "github.com/github/gh-classroom/cmd/gh-classroom/pull/starter-repo"
	student_repos "github.com/github/gh-classroom/cmd/gh-classroom/pull/student-repos"
	"github.com/spf13/cobra"
)

func NewCmdPull(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pull <starter-code|student-repos>",
		Short:   "Pull starter code or a student's submissions",
		Example: "",
	}

	cmd.AddCommand(starter_repo.NewCmdStarterRepoPull(f))
	cmd.AddCommand(student_repos.NewCmdStudentReposPull(f))

	cmd.PersistentFlags().StringP("directory", "d", ".", "Directory to clone into")
	cmd.PersistentFlags().IntP("assignment-id", "a", 0, "ID of the assignment")
	return cmd
}
