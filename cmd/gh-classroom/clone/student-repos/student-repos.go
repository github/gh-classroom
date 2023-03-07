package student_repos

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdStudentRepo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "student-repos",
		Short: "Clone student repos for an assignment",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Clone student repos for an assignment")
		},
	}
	return cmd
}
