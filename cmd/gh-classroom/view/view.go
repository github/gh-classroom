package view

import (
	"fmt"
	"log"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/browser"
	"github.com/cli/go-gh/v2/pkg/term"
	"github.com/github/gh-classroom/cmd/gh-classroom/shared"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/spf13/cobra"
)

func NewCmdView(f *cmdutil.Factory) *cobra.Command {
	var web bool
	var classroomId int

	cmd := &cobra.Command{
		Use:     "view",
		Example: `$ gh classroom view -c 4876 --web`,
		Short:   "Show the details of a classroom",
		Long: `Display the classroom ID, classroom slug, title and other information about a classroom.
With "--web", open the classroom in a browser instead.`,
		Run: func(cmd *cobra.Command, args []string) {
			client, err := api.DefaultRESTClient()

			if err != nil {
				log.Fatal(err)
				return
			}

			if classroomId == 0 {
				classroom, err := shared.PromptForClassroom(client)
				classroomId = classroom.Id

				if err != nil {
					log.Fatal(err)
					return
				}
			}

			response, err := classroom.GetClassroom(client, classroomId)
			if err != nil {
				log.Fatal(err)
			}

			if web {
				OpenInBrowser(response.Url)
				return
			}

			RenderModel(response, cmd.OutOrStdout())
		},
	}

	cmd.Flags().BoolVarP(&web, "web", "w", false, "Open classroom in the browser")
	cmd.Flags().IntVarP(&classroomId, "classroom-id", "c", 0, "ID of the classroom")
	return cmd
}

func OpenInBrowser(url string) {
	term := term.FromEnv()
	io := iostreams.System()
	c := io.ColorScheme()
	if term.IsTerminalOutput() {
		fmt.Fprintln(io.ErrOut, c.Yellow("\nOpening classroom in your browser...\n"))
	}
	browser := browser.New("", io.Out, io.ErrOut)
	err := browser.Browse(url)
	if err != nil {
		log.Fatal(err)
	}
}
