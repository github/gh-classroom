package main

import (
	"fmt"

	//"github.com/cli/go-gh"
	"github.com/github/gh-classroom/cmd/gh-classroom/root"
)

func main() {
	// Create a new root command
	cmd := root.NewRootCmd()
	// execute the root command
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}

}

