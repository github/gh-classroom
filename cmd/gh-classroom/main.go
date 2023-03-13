package main

import (
	"fmt"

	"github.com/cli/cli/v2/pkg/cmd/factory"
	"github.com/github/gh-classroom/cmd/gh-classroom/root"
)

func main() {
	cmdFactory := factory.New("0.1.0")

	cmd := root.NewRootCmd(cmdFactory)
	err := cmd.Execute()

	if err != nil {
		fmt.Println(err)
	}
}
