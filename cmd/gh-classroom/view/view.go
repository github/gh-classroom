package view

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"log"
	"github.com/spf13/cobra"
	"github.com/cli/go-gh"
    "github.com/cli/go-gh/pkg/term"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/github/gh-classroom/pkg/classroom"
	// "github.com/cli/go-gh/pkg/tableprinter"
	// "github.com/spf13/cobra"
)

type Classroom struct {
	ID int `json:"id"`
	Slug string `json:"slug"`
	Title string `json:"title"`
	Url string `json:"url"`
}

func NewCmdView(f *cmdutil.Factory) *cobra.Command {
	var web bool
	var classroomID int
	term := term.FromEnv()
	io := iostreams.System()
	cmd := &cobra.Command{
		Use:   "view",
		Example: `$ gh classroom view 4 --web`,
		Short: "Show the details of a Classroom",
		Long: `Display the classroom ID, classroom slug, title and other information about a classroom.
With "--web", open the classroom in a browser instead
For more information about output formatting flags, see "gh help"`,
		Run: func(cmd *cobra.Command, args []string){
			client, err := gh.RESTClient(nil)
			if classroomID == 0 {
				log.Fatal("Missing classroom ID. Try again")
			}

			if err != nil {
				log.Fatal(err)
			}

			response, err := classroom.GetClassroom(client, classroomID)

			
			// err = client.Get(fmt.Sprintf("classrooms/%v", classroomID), &response)

			fmt.Println("Classroom View")
		
			if web {
				if term.IsTerminalOutput() {
					fmt.Fprintln(io.ErrOut, "Opening classroom your browser...")
				}
				browser := browser.New("", io.Out, io.ErrOut)
				browser.Browse(response.Url)
				return
			}
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open classroom in the browser")
	cmd.Flags().IntVarP(&classroomID, "classroom-id", "c", 0, "ID of the classroom")
	return cmd
}
