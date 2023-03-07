package assignment

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdAssignment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assignment",
		Short: "List your assignments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Show Assignment")
		},
	}

	return cmd
}
