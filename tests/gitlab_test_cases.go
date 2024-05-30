package tests

type GitLabTestCase struct {
	Name     string
	Args     []string
	EnvVars  map[string]string
	Expected map[string]string
}

var GitLabTestCases = []GitLabTestCase{
	{
		Name: "all flags",
		Args: []string{
			"app", "gitlab", "auto-mr",
			"--gitlab-url=https://gitlab.com",
			"--gitlab-token=glpat-c8jsUhzPpQuic-abxXMX",
			"--gitlab-project-id=58164058",
			"--mr-base-branch=main",
			"--mr-new-branch=feature/oa-branch",
			"--mr-target-branch=main",
			"--mr-title=Example MR",
			"--mr-description=Gitlab flags test",
			"--mr-squash=true",
			"--mr-remove-source-branch=true",
			"--src-file-path=gitlab/data/output_format_all.yaml",
			"--file-path=configs/app_config.yaml",
		},
		EnvVars: map[string]string{},
		Expected: map[string]string{
			"gitlab-url":              "https://gitlab.com",
			"gitlab-token":            "glpat-c8jsUhzPpQuic-abxXMX",
			"gitlab-project-id":       "58164058",
			"mr-base-branch":          "main",
			"mr-new-branch":           "feature/oa-branch",
			"mr-target-branch":        "main",
			"mr-title":                "Example MR",
			"mr-description":          "Gitlab flags test",
			"mr-squash":               "true",
			"mr-remove-source-branch": "true",
			"src-file-path":           "gitlab/data/output_format_all.yaml",
			"file-path":               "configs/app_config.yaml",
		},
	},
	{
		Name: "env vars",
		Args: []string{"app", "gitlab", "auto-mr"},
		EnvVars: map[string]string{
			"OA_GITLAB_URL":                     "https://gitlab.com",
			"OA_GITLAB_TOKEN":                   "glpat-c8jsUhzPpQuic-abxXMX",
			"OA_GITLAB_PROJECT_ID":              "58164058",
			"OA_GITLAB_MR_BASE_BRANCH":          "main",
			"OA_GITLAB_MR_NEW_BRANCH":           "feature/oa-branch",
			"OA_GITLAB_MR_TARGET_BRANCH":        "main",
			"OA_GITLAB_MR_TITLE":                "Example MR",
			"OA_GITLAB_MR_DESCRIPTION":          "Gitlab flags test",
			"OA_GITLAB_MR_SQUASH":               "true",
			"OA_GITLAB_MR_REMOVE_SOURCE_BRANCH": "true",
			"OA_GITLAB_SRC_FILE_PATH":           "gitlab/data/output_format_all.yaml",
			"OA_GITLAB_FILE_PATH":               "configs/app_config.yaml",
		},
		Expected: map[string]string{
			"gitlab-url":              "https://gitlab.com",
			"gitlab-token":            "glpat-c8jsUhzPpQuic-abxXMX",
			"gitlab-project-id":       "58164058",
			"mr-base-branch":          "main",
			"mr-new-branch":           "feature/oa-branch",
			"mr-target-branch":        "main",
			"mr-title":                "Example MR",
			"mr-description":          "Gitlab flags test",
			"mr-squash":               "true",
			"mr-remove-source-branch": "true",
			"src-file-path":           "gitlab/data/output_format_all.yaml",
			"file-path":               "configs/app_config.yaml",
		},
	},
}
