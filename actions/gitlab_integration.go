package actions

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/xanzy/go-gitlab"
)

// CreateGitLabClient создает клиент GitLab
func CreateGitLabClient(gitlabURL, gitlabToken string) (*gitlab.Client, error) {
	client, err := gitlab.NewClient(gitlabToken, gitlab.WithBaseURL(gitlabURL))
	if err != nil {
		return nil, fmt.Errorf("error creating GitLab client: %v", err)
	}
	return client, nil
}

// ReadFileContent читает содержимое файла
func ReadFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(content), nil
}

// CreateBranch создает новую ветку в GitLab
func CreateBranch(client *gitlab.Client, projectID int, newBranch, baseBranch string) error {
	branch, resp, err := client.Branches.CreateBranch(projectID, &gitlab.CreateBranchOptions{
		Branch: &newBranch,
		Ref:    &baseBranch,
	})
	if err != nil && resp.StatusCode != 400 {
		return fmt.Errorf("error creating branch: %v", err)
	}
	if resp.StatusCode == 400 {
		color.New(color.FgBlue).Printf("Branch already exists: %s\n", newBranch)
	} else {
		color.New(color.FgGreen).Printf("Branch created: %s\n", branch.WebURL)
	}
	return nil
}

// FileExists проверяет, существует ли файл в GitLab
func FileExists(client *gitlab.Client, projectID int, filePath, branch string) (bool, error) {
	_, resp, err := client.RepositoryFiles.GetFile(projectID, filePath, &gitlab.GetFileOptions{Ref: gitlab.String(branch)})
	if err != nil && resp.StatusCode != 404 {
		return false, fmt.Errorf("error checking file existence: %v", err)
	}
	return resp.StatusCode == 200, nil
}

// CreateOrUpdateFile создает или обновляет файл в GitLab
func CreateOrUpdateFile(client *gitlab.Client, projectID int, filePath, content, branch, commitMessage string) error {
	fileExists, err := FileExists(client, projectID, filePath, branch)
	if err != nil {
		return err
	}

	if fileExists {
		_, _, err = client.RepositoryFiles.UpdateFile(projectID, filePath, &gitlab.UpdateFileOptions{
			Branch:        &branch,
			Content:       gitlab.String(content),
			CommitMessage: gitlab.String(commitMessage),
		})
	} else {
		_, _, err = client.RepositoryFiles.CreateFile(projectID, filePath, &gitlab.CreateFileOptions{
			Branch:        &branch,
			Content:       gitlab.String(content),
			CommitMessage: gitlab.String(commitMessage),
		})
	}
	if err != nil {
		return fmt.Errorf("error creating or updating file: %v", err)
	}
	return nil
}

// MergeRequestExists проверяет, существует ли merge request для данной ветки
func MergeRequestExists(client *gitlab.Client, projectID int, sourceBranch, targetBranch string) (bool, error) {
	opts := &gitlab.ListProjectMergeRequestsOptions{
		SourceBranch: &sourceBranch,
		TargetBranch: &targetBranch,
		State:        gitlab.String("opened"),
	}

	mrs, _, err := client.MergeRequests.ListProjectMergeRequests(projectID, opts)
	if err != nil {
		return false, fmt.Errorf("error checking merge request existence: %v", err)
	}

	return len(mrs) > 0, nil
}

// CreateMergeRequest создает merge request в GitLab
func CreateMergeRequest(client *gitlab.Client, projectID int, sourceBranch, targetBranch, title, description string) error {
	mrExists, err := MergeRequestExists(client, projectID, sourceBranch, targetBranch)
	if err != nil {
		return err
	}

	if mrExists {
		color.New(color.FgBlue).Printf("Merge request already exists for branch: %s\n", sourceBranch)
		return nil
	}

	mr, _, err := client.MergeRequests.CreateMergeRequest(projectID, &gitlab.CreateMergeRequestOptions{
		SourceBranch: &sourceBranch,
		TargetBranch: &targetBranch,
		Title:        gitlab.String(title),
		Description:  gitlab.String(description),
	})
	if err != nil {
		return fmt.Errorf("error creating merge request: %v", err)
	}
	color.New(color.FgGreen).Printf("Merge request created: %s\n", mr.WebURL)
	return nil
}

// HandleGitLabMergeRequest выполняет все действия для создания merge request в GitLab
func HandleGitLabMergeRequest(gitlabURL, gitlabToken, filePath, baseBranch, newBranch, targetBranch, projectID string) error {
	client, err := CreateGitLabClient(gitlabURL, gitlabToken)
	if err != nil {
		return err
	}

	projectIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	content, err := ReadFileContent(filePath)
	if err != nil {
		return err
	}

	err = CreateBranch(client, projectIDInt, newBranch, baseBranch)
	if err != nil {
		return err
	}

	err = CreateOrUpdateFile(client, projectIDInt, filePath, content, newBranch, "Drained keys from file2 to output file")
	if err != nil {
		return err
	}

	err = CreateMergeRequest(client, projectIDInt, newBranch, targetBranch, "Drained keys from file2 to output file", "This merge request contains the changes after draining keys from file2 to the output file.")
	if err != nil {
		return err
	}

	return nil
}
