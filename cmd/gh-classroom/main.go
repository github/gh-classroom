package main

import (
	"fmt"
	"os"
	"strings"
	// "github.com/github/gh-classroom/cmd/gh-classroom/root"
	"github.com/github/gh-classroom/cmd/gh-classroom/list"
	// "github.com/cli/cli/v2/pkg/cmd/factory"

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
