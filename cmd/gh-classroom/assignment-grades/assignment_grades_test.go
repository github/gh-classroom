package grades

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestAssignmentGradesFatalOnInvalidAPIResponse(t *testing.T) {
	// Run the crashing code when FLAG is set
	if os.Getenv("FLAG") == "1" {
		defer gock.Off()
		t.Setenv("GITHUB_TOKEN", "999")

		gock.New("https://api.github.com").
			Get("/assignments/1234/grades").
			Reply(200).
			JSON(`{ }`)

		actual := new(bytes.Buffer)

		f := &cmdutil.Factory{}
		command := NewCmdAssignmentGrades(f)
		command.SetOut(actual)
		command.SetErr(actual)
		command.SetArgs([]string{
			"-a1234",
		})

		command.Execute() //nolint:errcheck
		return
	}

	// Runs the test above in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestAssignmentGradesFatalOnInvalidAPIResponse")
	cmd.Env = append(os.Environ(), "FLAG=1")
	err := cmd.Run()

	// Gets a fatal error
	e, ok := err.(*exec.ExitError)
	expectedErrorString := "exit status 1"
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorString, e.Error())
}

func TestGettingGradesIndividualAssignment(t *testing.T) {
	t.Run("writes a csv when grades are returned from API", func(t *testing.T) {
		defer gock.Off()
		t.Setenv("GITHUB_TOKEN", "999")

		// given an api response with grades returned
		gock.New("https://api.github.com").
			Get("/assignments/1234/grades").
			Reply(200).
			JSON(`[{
				"assignment_name": "assignment",
				"assignment_url": "assignment.url",
				"github_username": "username",
				"points_available": "100",
				"points_awarded": "97",
				"roster_identifier": "username@example.com",
				"starter_code_url":"startercode.url",
				"student_repository_name": "repo",
				"student_repository_url": "repo.url",
				"submission_timestamp": "MM-DD-YYYY"
			}]`)

		actual := new(bytes.Buffer)
		outputFile := filepath.Join(t.TempDir(), "grades.csv")
		f := &cmdutil.Factory{}
		command := NewCmdAssignmentGrades(f)
		command.SetOut(actual)
		command.SetErr(actual)
		command.SetArgs([]string{
			"-a1234",
			"-f" + outputFile,
		})

		// When the command is executed
		err := command.Execute()

		// There should:
		// - be no error
		// - be a CSV written to the file passed in
		assert.NoError(t, err, "Should not error")

		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			t.Errorf("Expected persisted file at %s, did not find it: %s", outputFile, err)
		}
		b, err := os.ReadFile(outputFile)
		if err != nil {
			fmt.Print(err)
		}

		expected :=
			"assignment_name,assignment_url,starter_code_url,github_username,roster_identifier,student_repository_name,student_repository_url,submission_timestamp,points_awarded,points_available\n" +
				"assignment,assignment.url,startercode.url,username,username@example.com,repo,repo.url,MM-DD-YYYY,97,100\n"
		assert.Equal(t, string(b), expected)
	})
}

func TestGettingGradesGroupAssignment(t *testing.T) {
	t.Run("writes a csv when grades are returned from API", func(t *testing.T) {
		defer gock.Off()
		t.Setenv("GITHUB_TOKEN", "999")

		// given an api response with grades returned
		gock.New("https://api.github.com").
			Get("/assignments/1234/grades").
			Reply(200).
			JSON(`[{
				"assignment_name": "assignment",
				"assignment_url": "assignment.url",
				"github_username": "username",
				"points_available": "100",
				"points_awarded": "97",
				"roster_identifier": "username@example.com",
				"starter_code_url":"startercode.url",
				"student_repository_name": "repo",
				"student_repository_url": "repo.url",
				"submission_timestamp": "MM-DD-YYYY",
				"group_name": "group"
			}]`)

		actual := new(bytes.Buffer)
		outputFile := filepath.Join(t.TempDir(), "grades.csv")
		f := &cmdutil.Factory{}
		command := NewCmdAssignmentGrades(f)
		command.SetOut(actual)
		command.SetErr(actual)
		command.SetArgs([]string{
			"-a1234",
			"-f" + outputFile,
		})

		// When the command is executed
		err := command.Execute()

		// There should:
		// - be no error
		// - be a CSV written to the file passed in
		assert.NoError(t, err, "Should not error")

		if _, err := os.Stat(outputFile); os.IsNotExist(err) {
			t.Errorf("Expected persisted file at %s, did not find it: %s", outputFile, err)
		}
		b, err := os.ReadFile(outputFile)
		if err != nil {
			fmt.Print(err)
		}

		expected :=
			"assignment_name,assignment_url,starter_code_url,github_username,roster_identifier,student_repository_name,student_repository_url,submission_timestamp,points_awarded,points_available,group_name\n" +
				"assignment,assignment.url,startercode.url,username,username@example.com,repo,repo.url,MM-DD-YYYY,97,100,group\n"
		assert.Equal(t, string(b), expected)
	})
}
