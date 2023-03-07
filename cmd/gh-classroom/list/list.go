package list

import (
	"fmt"
  "log"
	"os"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/jsonpretty"

	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Classrooms.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List Classrooms")
      args := []string{"api", "/classrooms"}
	    result, _, err := gh.Exec(args...)
	    if err != nil {
		    log.Fatal(err)
	    }

	    err = jsonpretty.Format(os.Stdout, &result, " ", true)
      if err != nil {
        log.Fatal(err);
      }
		},
	}
	return cmd
}
