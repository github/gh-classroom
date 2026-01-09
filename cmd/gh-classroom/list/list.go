package list

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/browser"
	"github.com/cli/go-gh/v2/pkg/tableprinter"
	"github.com/cli/go-gh/v2/pkg/term"
	"github.com/cli/go-gh/v2/pkg/text"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdList(f *cmdutil.Factory) *cobra.Command {
	var web bool
	var page int
	var perPage int

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List classrooms",
		Long:    "List of classrooms you own",
		Example: `$ gh classroom list --page 1`,
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			term := term.FromEnv()
			io := iostreams.System()
			cs := io.ColorScheme()

			client, err := api.DefaultRESTClient()
			if err != nil {
				log.Fatal(err)
			}

			response, err := classroom.ListClassrooms(client, page, perPage)
			if err != nil {
				log.Fatal(err)
			}

			count := len(response)

			if count == 0 {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "No classrooms found")
			} else {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%v\n\n", text.Pluralize(count, "Classroom"))
			}

			t := tableprinter.New(cmd.OutOrStdout(), true, 14)

			if web {
				if term.IsTerminalOutput() {
					_, _ = fmt.Fprintln(cmd.ErrOrStderr(), "Opening in your browser.")
				}
				browser := browser.New("", io.Out, io.ErrOut)
				err := browser.Browse("https://classroom.github.com/classrooms")
				if err != nil {
					log.Fatal(err)
				}
				return
			}

			t.AddField("ID", tableprinter.WithTruncate(nil))
			t.AddField("Name", tableprinter.WithTruncate(nil))
			t.AddField("URL", tableprinter.WithTruncate(nil))
			t.EndRow()
			for _, classroom := range response {
				t.AddField(cs.Green(strconv.Itoa(classroom.Id)), tableprinter.WithTruncate(nil))
				t.AddField(classroom.Name, tableprinter.WithTruncate(nil))
				t.AddField(classroom.Url, tableprinter.WithTruncate(nil))
				t.EndRow()
			}

			err = t.Render()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open the classroom list in a browser")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 30, "Number of classrooms per page")
	return cmd
}
