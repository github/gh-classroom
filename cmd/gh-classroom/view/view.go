package view

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdView() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show the details of a Classroom",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("View Classroom")
		},
	}
	return cmd
}
