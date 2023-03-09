package view

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdView(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show the details of a Classroom",
		Long: `Display the classroom ID, classroom slug, title and other information about a classroom.
With "--web", open the classroom in a browser instead
For more information about output formatting flags, see "gh help"`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("View Classroom")
		},
	}

	cmd.Flags().BoolP("web", "w", false, "Open classroom in the browser")
	cmd.Example = `gh classroom view 4 --web`
	return cmd
}
