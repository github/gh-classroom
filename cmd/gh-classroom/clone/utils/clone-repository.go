// cmd/gh-classroom/clone/utils/clone-repository.go

package utils

import (
    "fmt"
    "os"
    "bytes"
)

// GitHubExec defines an interface for executing GitHub CLI commands.
// This abstraction allows for easier testing and decoupling from the actual CLI.
// Exec invokes a gh command in a subprocess and captures the output and error streams.
type GitHubExec interface {
    Exec(args ...string) (stdout, stderr bytes.Buffer, err error)
}

// CloneRepository attempts to clone a repository into the specified path.
// It returns an error if the cloning process fails.
func CloneRepository(clonePath string, repoFullName string, gh GitHubExec) error {
    if _, err := os.Stat(clonePath); os.IsNotExist(err) {
        fmt.Printf("Cloning into: %v\n", clonePath)
        _, _, err := gh.Exec("repo", "clone", repoFullName, "--", clonePath)
        if err != nil {
            fmt.Printf("error cloning %s: %v\n", repoFullName, err)
            return fmt.Errorf("error cloning %s: %v", repoFullName, err)
        }
        return nil // Success
    }
    fmt.Printf("Skip existing repo: %v use gh classroom pull to get new commits\n", clonePath)
    return fmt.Errorf("repository already exists: %s", clonePath)
}