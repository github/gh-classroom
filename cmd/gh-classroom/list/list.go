package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Classrooms.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List Classrooms")
		},
	}
	return cmd
}
