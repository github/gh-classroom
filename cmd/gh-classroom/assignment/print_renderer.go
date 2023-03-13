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
	printClassroom(assignment.Classroom, out)
	fmt.Println()
	printAssigment(assignment, out)
	fmt.Println()
}

func printClassroom(model classroom.Classroom, stdout io.Writer) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Fprintln(stdout, c.Blue("CLASSROOM INFORMATION"))
	fmt.Fprintln(stdout, c.Yellow("ID:"), c.Green(strconv.Itoa(model.Id)))
	fmt.Fprintln(stdout, c.Yellow("Name:"), c.Green(model.Name))
	fmt.Fprintln(stdout, c.Yellow("Classroom URL:"), c.Green(model.Url))
}

func printAssigment(assignment classroom.Assignment, out io.Writer) {
	c := iostreams.NewColorScheme(true, true, true)
	fmt.Fprintln(out, c.Blue("ASSIGNMENT INFORMATION"))
	fmt.Fprintln(out, c.Yellow("ID:"), c.Green(strconv.Itoa(assignment.Id)))
	fmt.Fprintln(out, c.Yellow("Title:"), c.Green(assignment.Title))
	fmt.Fprintln(out, c.Yellow("Invite Link:"), c.Green(assignment.InviteLink))
	fmt.Fprintln(out, c.Yellow("Starter Code Repo URL:"), c.Green(assignment.StarterCodeRepository.HtmlUrl))
	fmt.Fprintln(out, c.Yellow("Type:"), c.Green(fmt.Sprintf("%v assignment", assignment.AssignmentType)))
	fmt.Fprintln(out, c.Yellow("Accepted:"), c.Green(strconv.Itoa(assignment.Accepted)))
	fmt.Fprintln(out, c.Yellow("Submissions:"), c.Green(strconv.Itoa(assignment.Submissions)))
	fmt.Fprintln(out, c.Yellow("Passing:"), c.Green(strconv.Itoa(assignment.Passing)))
}
