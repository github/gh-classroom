package view

import (
	"bytes"
	"testing"

	"github.com/github/gh-classroom/pkg/classroom"
	"github.com/stretchr/testify/assert"
)

func TestPrintingOutput(t *testing.T) {
	model := classroom.Classroom{
		Id:          1,
		Name:        "My Classroom",
		Url:         "https://classroom.github.com/my-classroom",
		Organization: classroom.GitHubOrganization{
			Login: "org-login",
			HtmlUrl:  "https://github.com/github-org",
		},
	}
	actual := &bytes.Buffer{}

	RenderModel(model, actual)

	assert.Contains(t, actual.String(), "CLASSROOM INFORMATION")
	assert.Contains(t, actual.String(), "My Classroom")
	assert.Contains(t, actual.String(), "https://classroom.github.com/my-classroom")
	assert.Contains(t, actual.String(), "GITHUB INFORMATION")
	assert.Contains(t, actual.String(), "org-login")
	assert.Contains(t, actual.String(), "https://github.com/github-org")
}
