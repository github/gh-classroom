package student_repos

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdStudentRepo(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "student-repos",
		Short: "Clone student repos for an assignment",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Clone student repos for an assignment")
		},
	}
	return cmd
}
