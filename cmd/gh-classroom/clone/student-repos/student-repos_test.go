package student_repos

import (
    "testing"

    "github.com/spf13/cobra"
)

// Mock command setup for testing
func newTestCmd() *cobra.Command {
    var page, perPage int
    var getAll bool

    cmd := &cobra.Command{
        Use: "test-cmd",
        Run: func(cmd *cobra.Command, args []string) {
            getAll = !cmd.Flags().Changed("page")
        },
    }

    cmd.Flags().IntVar(&page, "page", 1, "Page number")
    cmd.Flags().IntVar(&perPage, "per-page", 15, "Number of items per page")
    cmd.Flags().BoolVar(&getAll, "all", true, "Get all items")

    return cmd
}

func TestGetAllFlag(t *testing.T) {
    tests := []struct {
        name     string
        args     []string
        expected bool
    }{
        {"Default", []string{}, true},
        {"PerPage Set", []string{"--per-page", "20"}, true},
        {"Page Set", []string{"--page", "2"}, false},
        {"Page and PerPage Set", []string{"--page", "2", "--per-page", "20"}, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cmd := newTestCmd()
            cmd.SetArgs(tt.args)
            if err := cmd.Execute(); err != nil {
                t.Fatalf("cmd.Execute() failed with %v", err)
            }

            getAllFlag, err := cmd.Flags().GetBool("all")
            if err != nil {
                t.Fatalf("Failed to get 'all' flag: %v", err)
            }

            if getAllFlag != tt.expected {
                t.Errorf("Expected getAll to be %v, got %v", tt.expected, getAllFlag)
            }
        })
    }
}