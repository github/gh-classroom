package view

import  (
	"fmt"
	"io"
	"strconv"
	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/cli/cli/v2/pkg/iostreams"
)

func RenderModel(model classroom.Classroom, stdout io.Writer) {
	fmt.Fprintln(stdout)
	printClassroom(model, stdout)
	fmt.Fprintln(stdout)
	printOrganizationInfo(model.Organization, stdout)
	return
}

func printClassroom(model classroom.Classroom, stdout io.Writer) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Fprintln(stdout, c.Blue("CLASSROOM INFORMATION"))
	fmt.Fprintln(stdout, c.Yellow("ID:"), c.Green(strconv.Itoa(model.Id)))
	fmt.Fprintln(stdout, c.Yellow("Name:"), c.Green(model.Name))
	fmt.Fprintln(stdout, c.Yellow("Classroom URL:"), c.Green(model.Url))
	return
}

func printOrganizationInfo(organization classroom.GitHubOrganization, stdout io.Writer) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Fprintln(stdout, c.Blue("GITHUB INFORMATION"))
	fmt.Fprintln(stdout, c.Yellow("Login:"), c.Green(organization.Login))
	fmt.Fprintln(stdout, c.Yellow("Organization URL"), c.Green(organization.HtmlUrl))
	return
}
