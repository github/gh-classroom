package assignments

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdAssignments() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assignments",
		Short: "List your assignments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List Assignments")
		},
	}

	return cmd
}
