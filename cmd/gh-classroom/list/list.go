package main

import (
	"log"
	"os"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/jsonpretty"
)

func main() {
	args := []string{"api", "/classrooms"}
	result, _, err := gh.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}

	err = jsonpretty.Format(os.Stdout, &result, " ", true)
	if err != nil {
		log.Fatal(err);
	}
}




// package list

// import (
// 	// "fmt"
// 	"log"
// 	"github.com/cli/cli/v2/pkg/cmdutil"

// 	// "github.com/cli/browser"
// 	// "github.com/cli/go-gh/pkg/api"
// 	// "github.com/cli/go-gh/pkg/tableprinter"
// 	// "github.com/cli/go-gh/pkg/term"
// 	"github.com/cli/go-gh"
// 	"github.com/spf13/cobra"
// )

// func NewCmdList(f *cmdutil.Factory, runF func() error) *cobra.Command {
// 	listCmd := &cobra.Command{
// 		Short: "List the classrooms for an owner",
// 		Use:   "list",
// 		Example: `
// # list the current user's projects
// gh projects list
// # open projects for user monalisa in the browser
// gh projects list --user monalisa --web
// # list the projects for org github including closed projects
// gh projects list --org github --closed
// `,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			args := []string{"api", "/classrooms"}
// 			result, _, err := gh.Exec(args...)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			return result
// 		},
// 	}

// 	return listCmd
// }
