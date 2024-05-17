package actions

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"io/ioutil"
)

type GitLabClient struct {
	client *gitlab.Client
}

func NewGitLabClient(token string) *GitLabClient {
	return &GitLabClient{
		client: gitlab.NewClient(nil, token),
	}
}

func (g *GitLabClient) CreateBranch(projectID int, branchName, baseBranch string) error {
	branch, _, err := g.client.Branches.CreateBranch(projectID, &gitlab.CreateBranchOptions{
		Branch: &branchName,
		Ref:    &baseBranch,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created branch: %s\n", branch.Name)
	return nil
}

func (g *GitLabClient) CreateFile(projectID int, branchName, filePath, content string) error {
	commitAction := gitlab.CommitActionOptions{
		Action:   gitlab.FileCreate,
		FilePath: filePath,
		Content:  content,
	}

	commitMessage := "Add comparison result file"
	_, _, err := g.client.Commits.CreateCommit(projectID, &gitlab.CreateCommitOptions{
		Branch:        &branchName,
		CommitMessage: &commitMessage,
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
		SourceBranch: &sourceBranch,
		TargetBranch: &targetBranch,
		Title:        &title,
		Draft:        gitlab.Bool(true),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created Merge Request: %s\n", mr.Title)
	return nil
}

func SaveComparisonResult(filePath string, keysOnlyInFile1, keysOnlyInFile2 []string) error {
	content := "Keys only in file1:\n"
	for _, key := range keysOnlyInFile1 {
		content += fmt.Sprintf("- %s\n", key)
	}
	content += "\nKeys only in file2:\n"
	for _, key := range keysOnlyInFile2 {
		content += fmt.Sprintf("- %s\n", key)
	}

	return ioutil.WriteFile(filePath, []byte(content), 0644)
}
