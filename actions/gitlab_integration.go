package actions

import (
	"encoding/base64"
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
	if err != nil {
		if resp != nil && resp.StatusCode == 400 {
			_, _ = color.New(color.FgBlue).Printf("Branch already exists: %s\n", newBranch)
		} else {
			return fmt.Errorf("error creating branch: %v", err)
		}
	} else {
		_, _ = color.New(color.FgGreen).Printf("Branch created: %s\n", branch.WebURL)
	}
	return nil
}

// FileExists проверяет, существует ли файл в GitLab
func FileExists(client *gitlab.Client, projectID int, filePath, branch string) (bool, string, error) {
	file, resp, err := client.RepositoryFiles.GetFile(projectID, filePath, &gitlab.GetFileOptions{Ref: &branch})
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return false, "", nil
		}
		return false, "", fmt.Errorf("error checking file existence: %v", err)
	}
	if file == nil {
		return false, "", fmt.Errorf("unexpected nil file")
	}

	content, err := base64.StdEncoding.DecodeString(file.Content)
	if err != nil {
		return false, "", fmt.Errorf("error decoding file content: %v", err)
	}
	return true, string(content), nil
}

// CreateOrUpdateFile создает или обновляет файл в GitLab
func CreateOrUpdateFile(client *gitlab.Client, projectID int, filePath, content, branch, commitMessage string) error {
	fileExists, existingContent, err := FileExists(client, projectID, filePath, branch)
	if err != nil {
		return err
	}

	if fileExists {
		if content == existingContent {
			_, _ = color.New(color.FgBlue).Printf("File %s already exists in branch %s with the same content\n", filePath, branch)
			return nil
		}
		_, _, err = client.RepositoryFiles.UpdateFile(projectID, filePath, &gitlab.UpdateFileOptions{
			Branch:        &branch,
			Content:       &content,
			CommitMessage: &commitMessage,
		})
	} else {
		_, _, err = client.RepositoryFiles.CreateFile(projectID, filePath, &gitlab.CreateFileOptions{
			Branch:        &branch,
			Content:       &content,
			CommitMessage: &commitMessage,
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
func CreateMergeRequest(client *gitlab.Client, projectID int, sourceBranch, targetBranch, title, description string, squash, removeSourceBranch bool) error {
	mrExists, err := MergeRequestExists(client, projectID, sourceBranch, targetBranch)
	if err != nil {
		return err
	}

	if mrExists {
		_, _ = color.New(color.FgBlue).Printf("Merge request already exists for branch: %s\n", sourceBranch)
		return nil
	}

	mr, _, err := client.MergeRequests.CreateMergeRequest(projectID, &gitlab.CreateMergeRequestOptions{
		SourceBranch:       &sourceBranch,
		TargetBranch:       &targetBranch,
		Title:              &title,
		Description:        &description,
		Squash:             &squash,
		RemoveSourceBranch: &removeSourceBranch,
	})
	if err != nil {
		return fmt.Errorf("error creating merge request: %v", err)
	}
	_, _ = color.New(color.FgGreen).Printf("Merge request created: %s\n", mr.WebURL)
	return nil
}

// HandleGitLabMergeRequest выполняет все действия для создания merge request в GitLab
func HandleGitLabMergeRequest(gitlabURL, gitlabToken, srcFilePath, filePath, baseBranch, newBranch, targetBranch, projectID, mrTitle, mrDescription string, mrSquash, mrRemoveSourceBranch bool) error {
	client, err := CreateGitLabClient(gitlabURL, gitlabToken)
	if err != nil {
		return err
	}

	projectIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	content, err := ReadFileContent(srcFilePath)
	if err != nil {
		return err
	}

	// Создаем ветку перед созданием или обновлением файла
	err = CreateBranch(client, projectIDInt, newBranch, baseBranch)
	if err != nil {
		return err
	}

	err = CreateOrUpdateFile(client, projectIDInt, filePath, content, newBranch, "Drained keys from file2 to output file")
	if err != nil {
		return err
	}

	err = CreateMergeRequest(client, projectIDInt, newBranch, targetBranch, mrTitle, mrDescription, mrSquash, mrRemoveSourceBranch)
	if err != nil {
		return err
	}

	return nil
}
