package root

import (
	"github.com/spf13/cobra"

	acceptedAssignments "github.com/github/gh-classroom/cmd/gh-classroom/accepted-assignments"
	"github.com/github/gh-classroom/cmd/gh-classroom/assignment"
	"github.com/github/gh-classroom/cmd/gh-classroom/assignments"
	"github.com/github/gh-classroom/cmd/gh-classroom/clone"
	"github.com/github/gh-classroom/cmd/gh-classroom/list"
	"github.com/github/gh-classroom/cmd/gh-classroom/view"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classroom <command>",
		Short: "A GitHub Classroom CLI",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(view.NewCmdView())
	cmd.AddCommand(assignments.NewCmdAssignments())
	cmd.AddCommand(assignment.NewCmdAssignment())
	cmd.AddCommand(acceptedAssignments.NewCmdAcceptedAssignments())
	cmd.AddCommand(clone.NewCmdClone())

	return cmd
}
