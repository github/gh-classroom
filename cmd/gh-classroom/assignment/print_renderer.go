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
	c := iostreams.System().ColorScheme()
	_, _ = fmt.Fprintln(stdout, c.Blue("CLASSROOM INFORMATION"))
	_, _ = fmt.Fprintln(stdout, c.Yellow("ID:"), c.Green(strconv.Itoa(model.Id)))
	_, _ = fmt.Fprintln(stdout, c.Yellow("Name:"), c.Green(model.Name))
	_, _ = fmt.Fprintln(stdout, c.Yellow("Classroom URL:"), c.Green(model.Url))
}

func printAssigment(assignment classroom.Assignment, out io.Writer) {
	c := iostreams.System().ColorScheme()
	_, _ = fmt.Fprintln(out, c.Blue("ASSIGNMENT INFORMATION"))
	_, _ = fmt.Fprintln(out, c.Yellow("ID:"), c.Green(strconv.Itoa(assignment.Id)))
	_, _ = fmt.Fprintln(out, c.Yellow("Title:"), c.Green(assignment.Title))
	_, _ = fmt.Fprintln(out, c.Yellow("Invite Link:"), c.Green(assignment.InviteLink))
	_, _ = fmt.Fprintln(out, c.Yellow("Starter Code Repo URL:"), c.Green(assignment.StarterCodeRepository.HtmlUrl))
	_, _ = fmt.Fprintln(out, c.Yellow("Type:"), c.Green(fmt.Sprintf("%v assignment", assignment.AssignmentType)))
	_, _ = fmt.Fprintln(out, c.Yellow("Deadline:"), c.Green(assignment.Deadline))
	_, _ = fmt.Fprintln(out, c.Yellow("Accepted:"), c.Green(strconv.Itoa(assignment.Accepted)))
	_, _ = fmt.Fprintln(out, c.Yellow("Submissions:"), c.Green(strconv.Itoa(assignment.Submissions)))
	_, _ = fmt.Fprintln(out, c.Yellow("Passing:"), c.Green(strconv.Itoa(assignment.Passing)))
}
