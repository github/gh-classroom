package view

import (
	"fmt"
	"log"
	"strconv"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/cli/go-gh/pkg/term"
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
		Short:   "Show the details of a Classroom",
		Long: `Display the classroom ID, classroom slug, title and other information about a classroom.
With "--web", open the classroom in a browser instead
For more information about output formatting flags, see "gh help"`,
		Run: func(cmd *cobra.Command, args []string) {
			client, err := gh.RESTClient(nil)

			if err != nil {
				log.Fatal(err)
			}

			if classroomId == 0 {
				classroom, err := shared.PromptForClassroom(client)
				classroomId = classroom.Id

				if err != nil {
					log.Fatal(err)
				}
			}

			response, err := classroom.GetClassroom(client, classroomId)

			if web {
				OpenInBrowser(response.Url)
				return
			}

			fmt.Println()
			PrintClassroom(response)
			fmt.Println()
			PrintOrganizationInfo(response.Organization)
			return
		},
	}

	cmd.Flags().BoolVarP(&web, "web", "w", false, "Open classroom in the browser")
	cmd.Flags().IntVarP(&classroomId, "classroom-id", "c", 0, "ID of the classroom")
	return cmd
}

func OpenInBrowser(url string) {
	term := term.FromEnv()
	io := iostreams.System()
	c := iostreams.NewColorScheme(true, true, true)
	if term.IsTerminalOutput() {
		fmt.Fprintln(io.ErrOut, c.Yellow("\nOpening classroom in your browser...\n"))
	}
	browser := browser.New("", io.Out, io.ErrOut)
	browser.Browse(url)
	return
}

func PrintClassroom(response classroom.Classroom) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Println(c.Blue("CLASSROOM INFORMATION"))
	fmt.Println(c.Yellow("ID:"), c.Green(strconv.Itoa(response.Id)))
	fmt.Println(c.Yellow("Name:"), c.Green(response.Name))
	fmt.Println(c.Yellow("URL:"), c.Green(response.Url))
	return
}

func PrintOrganizationInfo(organization classroom.GitHubOrganizationInfo) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Println(c.Blue("GITHUB INFORMATION"))
	fmt.Println(c.Yellow("Login:"), c.Green(organization.Login))
	fmt.Println(c.Yellow("Organization URL"), c.Green(organization.HtmlUrl))
	return
}
