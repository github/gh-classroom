package utils

import (
    "os"
    "testing"
    "bytes"
    "fmt"
    "errors"
    "path/filepath"
)

func TestCloneRepository(t *testing.T) {
    tmpDir, err := os.MkdirTemp("", "cloneTest")
    if err != nil {
        t.Fatalf("Failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tmpDir) // clean up
    // Test cases
    tests := []struct {
        name       string
        execMock   GitHubExec
        clonePath  string
        repoFullName string
        wantErr    bool
        errMsg     string
    }{
        {
            name: "successful clone",
            execMock: func(args ...string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
                    var stdoutBuf, stderrBuf bytes.Buffer
                    stdoutBuf.Write([]byte("your string here"))
                    return stdoutBuf, stderrBuf, nil
            },
            clonePath: "", // Will be set to a temp directory in the test
            repoFullName: "example/repo",
            wantErr:    false,
        },
        {
            name: "clone failure",
            execMock: func(args ...string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
                    var stdoutBuf, stderrBuf bytes.Buffer
                    return stdoutBuf, stderrBuf, errors.New("clone error")
            },
            clonePath:  filepath.Join(tmpDir, "repo"),
            repoFullName: "example/repo",
            wantErr:    true,
            errMsg:     "error cloning example/repo: clone error",
        },
        {
            name: "repository already exists",
            execMock: func(args ...string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
                    var stdoutBuf, stderrBuf bytes.Buffer
                    return stdoutBuf, stderrBuf, nil
            },
            clonePath:  "./repo", // Current directory always exists
            repoFullName: "example/repo",
            wantErr:    true,
            errMsg:     "repository already exists: .",
        },
    }


    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.name == "successful clone" {
                fmt.Println("Running successful clone test")
                tmpDir, err := os.MkdirTemp("", "cloneTest")
                if err != nil {
                    t.Fatalf("Failed to create temp directory: %v", err)
                }
                defer os.RemoveAll(tmpDir) // clean up
                tt.clonePath = filepath.Join(tmpDir, "repo")
            }


            fmt.Println("Running test", tt.name, "with clonePath", tt.clonePath, "and repoFullName", tt.repoFullName)
            err := CloneRepository(tt.clonePath, tt.repoFullName, tt.execMock)
            if err != nil && err.Error() != tt.errMsg {
                t.Errorf("CloneRepository() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}