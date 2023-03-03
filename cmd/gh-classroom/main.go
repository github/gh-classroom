package main

import (
	"fmt"
	"os"
	"strings"
	// "github.com/github/gh-classroom/cmd/gh-classroom/root"
	"github.com/github/gh-classroom/cmd/gh-classroom/list"
	// "github.com/cli/cli/v2/pkg/cmd/factory"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:           "classroom",
		Short:         "Work with GitHub Projects.",
		Long:          "Work with GitHub Projects. Note that the token you are using must have 'project' scope, which is not set by default. You can verify your token scope by running 'gh auth status' and add the project scope by running 'gh auth refresh -s project'.",
		SilenceErrors: true,
	}

	// cmdFactory := factory.New("0.1.0")

	// rootCmd.AddCommand(cmdList.NewCmdList(cmdFactory, nil))

	if err := rootCmd.Execute(); err != nil {
		if strings.HasPrefix(err.Error(), "Message: Your token has not been granted the required scopes to execute this query") {
			fmt.Println("Your token has not been granted the required scopes to execute this query.\nRun 'gh auth refresh -s project' to add the 'project' scope.\nRun 'gh auth status' to see your current token scopes.")
			os.Exit(1)
		}
		fmt.Println(err)
		os.Exit(1)
	}
	list.list()
	// fmt.Println("hi world, this is the gh-classroom extension!")
	// client, err := gh.RESTClient(nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// response := struct{ Login string }{}
	// err = client.Get("user", &response)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("running as %s\n", response.Login)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
