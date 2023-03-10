package view

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/tableprinter"
	"github.com/cli/go-gh/pkg/term"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

type Classroom struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func NewCmdView(f *cmdutil.Factory) *cobra.Command {
	var web bool
	var classroomID int

	cmd := &cobra.Command{
		Use:     "view",
		Example: `$ gh classroom view -c 4876 --web`,
		Short:   "Show the details of a Classroom",
		Long: `Display the classroom ID, classroom slug, title and other information about a classroom.
With "--web", open the classroom in a browser instead
For more information about output formatting flags, see "gh help"`,
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gh.RESTClient(nil)
			response, err := classroom.GetClassroom(client, classroomID)

			if classroomID == 0 {
				log.Fatal("Missing classroom ID. Try again")
			}

			if err != nil {
				log.Fatal(err)
			}

			if web {
				OpenInBrowser(response.Url)
				return
			}

			fmt.Println("CLASSROOM INFORMATION")
			t := tableprinter.New(cmd.OutOrStdout(), true, 14)
			PrintTable(response, t)

		},
	}

	cmd.Flags().BoolVarP(&web, "web", "w", false, "Open classroom in the browser")
	cmd.Flags().IntVarP(&classroomID, "classroom-id", "c", 0, "ID of the classroom")
	return cmd
}

func PrintTable(response classroom.LongClassroom, t tableprinter.TablePrinter) {
	t.AddField("ID", tableprinter.WithTruncate(nil))
	t.AddField("Name", tableprinter.WithTruncate(nil))
	t.AddField("Archived", tableprinter.WithTruncate(nil))
	t.AddField("URL", tableprinter.WithTruncate(nil))
	t.EndRow()

	t.AddField(strconv.Itoa(response.Id), tableprinter.WithTruncate(nil))
	t.AddField(response.Name, tableprinter.WithTruncate(nil))
	t.AddField(strconv.FormatBool(response.Archived))
	t.AddField(response.Url, tableprinter.WithTruncate(nil))
	t.EndRow()

	t.Render()
}

func OpenInBrowser(url string) {
	term := term.FromEnv()
	io := iostreams.System()

	if term.IsTerminalOutput() {
		fmt.Fprintln(io.ErrOut, "Opening classroom in your browser...")
	}
	browser := browser.New("", io.Out, io.ErrOut)
	browser.Browse(url)
	return
}
