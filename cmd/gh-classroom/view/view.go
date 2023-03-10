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
			if len(args) != 0 {
				classroomID, err = strconv.Atoi(args[0]) //allows user to input classroom ID as an argument or flag
			}
			response, err := classroom.GetClassroom(client, classroomID)

			if err != nil {
				log.Fatal(err)
			}

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

func PrintClassroom(response classroom.LongClassroom) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Println(c.Blue("CLASSROOM INFORMATION"))
	fmt.Println(c.Yellow("ID:"), c.Green(strconv.Itoa(response.Id)))
	fmt.Println(c.Yellow("Name:"), c.Green(response.Name))
	fmt.Println(c.Yellow("Archived:"), c.Green(strconv.FormatBool(response.Archived)))
	fmt.Println(c.Yellow("URL:"), c.Green(response.Url))
	return
}

func PrintOrganizationInfo(organization classroom.GitHubOrganizationInfo){
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Println(c.Blue("GITHUB INFORMATION"))
	fmt.Println(c.Yellow("GitHub Organization ID:"), c.Green(strconv.Itoa(organization.Id)))
	fmt.Println(c.Yellow("Login:"), c.Green(organization.Login))
	fmt.Println(c.Yellow("Node_ID:"), c.Green(organization.NodeID))
	fmt.Println(c.Yellow("Html_Url:"), c.Green(organization.HtmlUrl))
	fmt.Println(c.Yellow("Name:"), c.Green(organization.Name))
	fmt.Println(c.Yellow("Avatar URL:"), c.Green(organization.AvatarUrl))
	return
}
