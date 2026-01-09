package assignment

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

func NewCmdAssignment(f *cmdutil.Factory) *cobra.Command {
	var (
		web          bool
		assignmentId int
	)

	cmd := &cobra.Command{
		Use:     "assignment",
		Example: `$ gh classroom assignment -a 4876`,
		Short:   "Show the details of an assignment",
		Long: "Display the details of an assignment",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := api.DefaultRESTClient()

			if err != nil {
				log.Fatal(err)
			}

			if assignmentId == 0 {
				cr, err := shared.PromptForClassroom(client)
				classroomId := cr.Id

				if err != nil {
					log.Fatal(err)
				}

				assignment, err := shared.PromptForAssignment(client, classroomId)
				if err != nil {
					log.Fatal(err)
				}
				assignmentId = assignment.Id
			}

			response, err := classroom.GetAssignment(client, assignmentId)

			if err != nil {
				log.Fatal(err)
			}

			if web {
				OpenInBrowser(response.Url())
				return
			}

			RenderModel(response, cmd.OutOrStdout())
		},
	}

	cmd.Flags().BoolVar(&web, "web", false, "Open specified assignment in a web browser")
	cmd.Flags().IntVarP(&assignmentId, "assignment-id", "a", 0, "Assignment ID (required)")

	return cmd
}

func OpenInBrowser(url string) {
	term := term.FromEnv()
	io := iostreams.System()
	c := iostreams.NewColorScheme(true, true, true)
	if term.IsTerminalOutput() {
		fmt.Fprintln(io.ErrOut, c.Yellow("\nOpening assigment in your browser...\n"))
	}

	browser := browser.New("", io.Out, io.ErrOut)
	err := browser.Browse(url)
	if err != nil {
		log.Fatal(err)
	}
}
