package assignment

import (
	"fmt"
	"io"
	"strconv"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/github/gh-classroom/pkg/classroom"
)

func RenderModel(assignment classroom.Assignment, out io.Writer) {
	fmt.Println()
	printAssigment(assignment, out)
	fmt.Println()
}

func printAssigment(assignment classroom.Assignment, out io.Writer) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Fprintln(out, c.Blue("ASSIGNMENT INFORMATION"))
	fmt.Fprintln(out, c.Yellow("ID:"), c.Green(strconv.Itoa(assignment.Id)))
	fmt.Fprintln(out, c.Yellow("Title:"), c.Green(assignment.Title))
	fmt.Fprintln(out, c.Yellow("Invite Link:"), c.Green(assignment.InviteLink))
	fmt.Fprintln(out, c.Yellow("Accepted:"), c.Green(strconv.Itoa(assignment.Accepted)))
	fmt.Fprintln(out, c.Yellow("Submissions:"), c.Green(strconv.Itoa(assignment.Submissions)))
	fmt.Fprintln(out, c.Yellow("Passing:"), c.Green(strconv.Itoa(assignment.Passing)))
}
