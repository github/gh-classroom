package list

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/tableprinter"
	"github.com/cli/go-gh/pkg/term"
	"github.com/spf13/cobra"
)

type Classroom struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	Url      string `json:"url"`
}

func NewCmdList() *cobra.Command {
	var web bool
	var page int
	var perPage int

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List Classrooms",
		Long:    "List of Classrooms you own.",
		Example: `$ gh classroom list --page 1`,
		Run: func(cmd *cobra.Command, args []string) {
			term := term.FromEnv()
			io := iostreams.System()

			client, err := gh.RESTClient(nil)
			if err != nil {
				log.Fatal(err)
			}

			var response []Classroom
			err = client.Get(fmt.Sprintf("classrooms?page=%v&per_page=%v", page, perPage), &response)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Classrooms")

			t := tableprinter.New(os.Stdout, true, 14)

			if web {
				if term.IsTerminalOutput() {
					fmt.Fprintln(io.ErrOut, "Opening in your browser.")
				}
				browser := browser.New("", io.Out, io.ErrOut)
				browser.Browse("https://classroom.github.com/classrooms")
				return
			}

			t.AddField("ID", tableprinter.WithTruncate(nil))
			t.AddField("Name", tableprinter.WithTruncate(nil))
			t.AddField("Archived", tableprinter.WithTruncate(nil))
			t.AddField("URL", tableprinter.WithTruncate(nil))
			t.EndRow()
			for _, classroom := range response {
				t.AddField(strconv.Itoa(classroom.Id), tableprinter.WithTruncate(nil))
				t.AddField(classroom.Name, tableprinter.WithTruncate(nil))
				t.AddField(strconv.FormatBool(classroom.Archived))
				t.AddField(classroom.Url, tableprinter.WithTruncate(nil))
				t.EndRow()
			}
			t.Render()
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open the classroom list in a browser")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&perPage, "per-page", 30, "Number of classrooms per page")
	return cmd
}
