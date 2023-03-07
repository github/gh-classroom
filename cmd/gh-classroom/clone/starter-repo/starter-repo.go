package starter_repo

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdStarterRepo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "starter-repo",
		Short: "Clone starter code",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Clone starter code")
		},
	}
	return cmd
}
