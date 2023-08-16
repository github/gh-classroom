package root

import (
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"

	acceptedAssignments "github.com/github/gh-classroom/cmd/gh-classroom/accepted-assignments"
	"github.com/github/gh-classroom/cmd/gh-classroom/assignment"
	assignmentgrades "github.com/github/gh-classroom/cmd/gh-classroom/assignment-grades"
	"github.com/github/gh-classroom/cmd/gh-classroom/assignments"
	"github.com/github/gh-classroom/cmd/gh-classroom/clone"
	"github.com/github/gh-classroom/cmd/gh-classroom/list"
	"github.com/github/gh-classroom/cmd/gh-classroom/view"
)

func NewRootCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classroom <command>",
		Short: "A GitHub Classroom CLI",
	}

	cmd.AddCommand(list.NewCmdList(f))
	cmd.AddCommand(view.NewCmdView(f))
	cmd.AddCommand(assignments.NewCmdAssignments(f))
	cmd.AddCommand(assignment.NewCmdAssignment(f))
	cmd.AddCommand(acceptedAssignments.NewCmdAcceptedAssignments(f))
	cmd.AddCommand(clone.NewCmdClone(f))
	cmd.AddCommand(assignmentgrades.NewCmdAssignmentGrades(f))

	return cmd
}
