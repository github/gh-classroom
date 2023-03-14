package classroom

import (
	"fmt"
	"strings"
)

type AssignmentList struct {
	Assignments []Assignment
	Classroom   Classroom
	Count       int
}

type Assignment struct {
	Id                          int              `json:"id"`
	PublicRepo                  bool             `json:"public_repo"`
	Title                       string           `json:"title"`
	AssignmentType              string           `json:"type"`
	InviteLink                  string           `json:"invite_link"`
	InvitationsEnabled          bool             `json:"invitations_enabled"`
	Slug                        string           `json:"slug"`
	StudentsAreRepoAdmins       bool             `json:"students_are_repo_admins"`
	FeedbackPullRequestsEnabled bool             `json:"feedback_pull_requests_enabled"`
	MaxTeams                    int              `json:"max_teams"`
	MaxMembers                  int              `json:"max_members"`
	Editor                      string           `json:"editor"`
	Accepted                    int              `json:"accepted"`
	Submissions                 int              `json:"submissions"`
	Passing                     int              `json:"passing"`
	Language                    string           `json:"language"`
	Classroom                   Classroom        `json:"classroom"`
	StarterCodeRepository       GithubRepository `json:"starter_code_repository"`
}

type Classroom struct {
	Id           int                `json:"id"`
	Name         string             `json:"name"`
	Archived     bool               `json:"archived"`
	Url          string             `json:"url"`
	Organization GitHubOrganization `json:"organization"`
}

type GithubRepository struct {
	Id            int    `json:"id"`
	FullName      string `json:"full_name"`
	HtmlUrl       string `json:"html_url"`
	NodeId        string `json:"node_id"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`
}

type Student struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	HtmlUrl   string `json:"html_url"`
}

type AcceptedAssignment struct {
	Id                     int              `json:"id"`
	Submitted              bool             `json:"submitted"`
	Passing                bool             `json:"passing"`
	CommitCount            int              `json:"commit_count"`
	Grade                  string           `json:"grade"`
	FeedbackPullRequestUrl string           `json:"feedback_pull_request_url"`
	Students               []Student        `json:"students"`
	Repository             GithubRepository `json:"repository"`
	Assignment             Assignment       `json:"assignment"`
}

type AcceptedAssignmentList struct {
	AcceptedAssignments []AcceptedAssignment
	Classroom           Classroom
	Assignment          Assignment
	Count               int
}

type GitHubOrganization struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	NodeID    string `json:"node_id"`
	HtmlUrl   string `json:"html_url"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

func NewAssignmentList(assignments []Assignment) AssignmentList {
	classroom := assignments[0].Classroom
	count := len(assignments)

	return AssignmentList{
		Assignments: assignments,
		Classroom:   classroom,
		Count:       count,
	}
}

func NewAcceptedAssignmentList(assignments []AcceptedAssignment) AcceptedAssignmentList {
	classroom := assignments[0].Assignment.Classroom
	assignment := assignments[0].Assignment
	count := len(assignments)

	return AcceptedAssignmentList{
		AcceptedAssignments: assignments,
		Classroom:           classroom,
		Assignment:          assignment,
		Count:               count,
	}
}

func (a AssignmentList) Url() string {
	return fmt.Sprintf(a.Classroom.Url)
}

func (a Assignment) Url() string {
	return fmt.Sprintf(a.Classroom.Url+"/assignments/%v", a.Slug)
}

func (a Assignment) IsGroupAssignment() bool {
	return a.AssignmentType == "group"
}

func (a AcceptedAssignment) RepositoryUrl() string {
	return a.Repository.HtmlUrl
}

func (gr GithubRepository) Name() string {
	return strings.Split(gr.FullName, "/")[1]
}
