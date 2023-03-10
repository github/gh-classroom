package assignment

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdAssignment(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assignment",
		Short: "List your assignments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Show Assignment")
		},
	}

	return cmd
}
