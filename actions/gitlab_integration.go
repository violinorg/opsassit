package actions

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"io/ioutil"
)

type GitLabClient struct {
	client *gitlab.Client
}

func NewGitLabClient(token string) (*GitLabClient, error) {
	client, err := gitlab.NewClient(token)
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
		// Draft field does not exist in go-gitlab package, use WIP prefix in title instead
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
