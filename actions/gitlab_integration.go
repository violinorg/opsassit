package actions

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

type GitLabClient struct {
	client *gitlab.Client
}

func NewGitLabClient(url, token string) (*GitLabClient, error) {
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		return nil, err
	}
	return &GitLabClient{client: client}, nil
}

func (g *GitLabClient) CreateBranch(projectID int, branchName, baseBranch string) error {
	branch, _, err := g.client.Branches.CreateBranch(projectID, &gitlab.CreateBranchOptions{
		Branch: gitlab.String(branchName),
		Ref:    gitlab.String(baseBranch),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created branch: %s\n", branch.Name)
	return nil
}

func (g *GitLabClient) CreateFile(projectID int, branchName, filePath, content string) error {
	commitAction := gitlab.CommitActionOptions{
		Action:   gitlab.FileAction(gitlab.FileCreate),
		FilePath: gitlab.String(filePath),
		Content:  gitlab.String(content),
	}

	commitMessage := "Add comparison result file"
	_, _, err := g.client.Commits.CreateCommit(projectID, &gitlab.CreateCommitOptions{
		Branch:        gitlab.String(branchName),
		CommitMessage: gitlab.String(commitMessage),
		Actions:       []*gitlab.CommitActionOptions{&commitAction},
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created file: %s\n", filePath)
	return nil
}

func (g *GitLabClient) CreateMergeRequest(projectID int, sourceBranch, targetBranch, title string) error {
	mr, _, err := g.client.MergeRequests.CreateMergeRequest(projectID, &gitlab.CreateMergeRequestOptions{
		SourceBranch: gitlab.String(sourceBranch),
		TargetBranch: gitlab.String(targetBranch),
		Title:        gitlab.String(title),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created Merge Request: %s\n", mr.Title)
	return nil
}
