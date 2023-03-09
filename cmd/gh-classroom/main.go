package main

import (
	"fmt"
	"github.com/spf13/cobra"
	//"github.com/cli/go-gh"
	"github.com/github/gh-classroom/cmd/gh-classroom/root"
)

func main() {
	// Create a new root command
	cmd := rootCmd()
	// execute the root command
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}

}

func rootCmd() *cobra.Command {
	return root.NewRootCmd()
}
