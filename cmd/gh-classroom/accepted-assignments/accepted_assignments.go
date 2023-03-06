package accepted

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdAcceptedAssignments() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accepted-assignments",
		Short: "List your student's accepted assignments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List Accepted Assignments")
		},
	}

	return cmd
}
