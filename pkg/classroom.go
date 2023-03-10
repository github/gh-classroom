package classroom

import (
	"fmt"
)

type AssignmentList struct {
	Assignments []Assignment
	Classroom   ShortClassroom
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
	Classroom                   ShortClassroom   `json:"classroom"`
	StarterCodeRepository       GithubRepository `json:"starter_code_repository"`
}

type ShortClassroom struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	Url      string `json:"url"`
}

type GithubRepository struct {
	Id            int    `json:"id"`
	FullName      string `json:"full_name"`
	HtmlUrl       string `json:"html_url"`
	NodeId        string `json:"node_id"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`
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

func (a AssignmentList) Url() string {
	return fmt.Sprintf(a.Classroom.Url)
}

func (a Assignment) Url() string {
	return fmt.Sprintf(a.Classroom.Url+"/assignments/%v", a.Slug)
}