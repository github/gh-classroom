package list

import (
	"fmt"
	"log"
	"os"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/jsonpretty"

	"github.com/spf13/cobra"
)

type listOpts struct {
	page     int
}

func NewCmdList() *cobra.Command {
	opts := listOpts{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Classrooms.",
		Long: "List of Classrooms you own.",
		Example: `$ gh classroom list --page 1`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List Classrooms")
			for_api := []string{"api", "/classrooms"}
			result, _, err := gh.Exec(for_api...)
			if err != nil {
				log.Fatal(err)
			}

			err = jsonpretty.Format(os.Stdout, &result, " ", true)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().IntVarP(&opts.page, "page", "p", 1, "Search by page number.")
	return cmd
}
