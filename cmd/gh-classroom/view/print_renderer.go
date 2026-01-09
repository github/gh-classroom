package view

import (
	"fmt"
	"io"
	"strconv"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/github/gh-classroom/pkg/classroom"
)

func RenderModel(model classroom.Classroom, stdout io.Writer) {
	_, _ = fmt.Fprintln(stdout)
	printClassroom(model, stdout)
	_, _ = fmt.Fprintln(stdout)
	printOrganizationInfo(model.Organization, stdout)
}

func printClassroom(model classroom.Classroom, stdout io.Writer) {
	c := iostreams.System().ColorScheme()
	_, _ = fmt.Fprintln(stdout, c.Blue("CLASSROOM INFORMATION"))
	_, _ = fmt.Fprintln(stdout, c.Yellow("ID:"), c.Green(strconv.Itoa(model.Id)))
	_, _ = fmt.Fprintln(stdout, c.Yellow("Name:"), c.Green(model.Name))
	_, _ = fmt.Fprintln(stdout, c.Yellow("Classroom URL:"), c.Green(model.Url))
}

func printOrganizationInfo(organization classroom.GitHubOrganization, stdout io.Writer) {
	c := iostreams.System().ColorScheme()
	_, _ = fmt.Fprintln(stdout, c.Blue("GITHUB INFORMATION"))
	_, _ = fmt.Fprintln(stdout, c.Yellow("Login:"), c.Green(organization.Login))
	_, _ = fmt.Fprintln(stdout, c.Yellow("Organization URL"), c.Green(organization.HtmlUrl))
}
