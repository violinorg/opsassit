package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/violinorg/opsassit/actions"
)

func yamlCmd() *cli.Command {
	return &cli.Command{
		Name:  "yaml",
		Usage: "Commands for working with YAML files",
		Subcommands: []*cli.Command{
			diffCmd(),
		},
	}
}

func diffCmd() *cli.Command {
	return &cli.Command{
		Name:      "diff",
		Usage:     "Compare and update YAML files",
		ArgsUsage: "[file1] [file2]",
		Action:    diffAction,
		Flags: addGitLabFlags([]cli.Flag{
			&cli.BoolFlag{
				Name:  "approved",
				Usage: "Apply the changes to the file",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "Output file path",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "Output format (all, values, keys)",
			},
		}),
	}
}

func diffAction(c *cli.Context) error {
	file1Path := os.Getenv("FILE1_PATH")
	file2Path := os.Getenv("FILE2_PATH")
	outputPath := os.Getenv("OA_DRAIN_OUTPUT")

	if file1Path == "" || file2Path == "" {
		if c.NArg() != 2 {
			return fmt.Errorf("expected exactly 2 arguments")
		}
		file1Path = c.Args().Get(0)
		file2Path = c.Args().Get(1)
	}

	if outputPath == "" {
		outputPath = c.String("output")
	}

	format := c.String("format")
	if format == "" {
		format = "all"
	}

	approved := c.Bool("approved")

	// Read file1 to check for "# OpsAssist Verified"
	file1Content, err := os.ReadFile(file1Path)
	if err != nil {
		return fmt.Errorf("error reading file1: %v", err)
	}

	if strings.Contains(string(file1Content), "# OpsAssist Verified") {
		fmt.Println("The file is already tuned.")
		return nil
	}

	vars1, err := actions.LoadVariablesFromYAMLWithOrder(file1Path)
	if err != nil {
		return fmt.Errorf("error loading file1: %v", err)
	}

	vars2, err := actions.LoadVariablesFromYAMLWithOrder(file2Path)
	if err != nil {
		return fmt.Errorf("error loading file2: %v", err)
	}

	var updatedYAML string

	switch format {
	case "values":
		differences := actions.CompareValues(vars1, vars2)
		updatedYAML, err = actions.GenerateValuesComparisonYAML(differences)
		if err != nil {
			return fmt.Errorf("error generating values comparison YAML: %v", err)
		}
	case "keys":
		onlyInFile1, onlyInFile2 := actions.CompareKeys(vars1, vars2)
		updatedYAML, err = actions.GenerateKeysComparisonYAML(onlyInFile1, onlyInFile2)
		if err != nil {
			return fmt.Errorf("error generating keys comparison YAML: %v", err)
		}
	case "all":
		fallthrough
	default:
		updatedYAML, err = actions.GenerateUpdatedYAML(vars1, vars2)
		if err != nil {
			return fmt.Errorf("error generating updated YAML: %v", err)
		}
	}

	// Output the changes
	fmt.Println("Proposed changes:")
	fmt.Println(updatedYAML)

	if approved {
		updatedYAML += "\n# OpsAssist Verified\n"

		err = os.WriteFile(outputPath, []byte(updatedYAML), 0644)
		if err != nil {
			return fmt.Errorf("error writing updated YAML to output file: %v", err)
		}

		fmt.Println("Successfully drained keys from file2 to output file.")

		// Handle GitLab merge request if GitLab flags are set
		gitlabURL := c.String("gitlab-url")
		gitlabToken := c.String("gitlab-token")
		projectID := c.String("project-id")
		baseBranch := c.String("base-branch")
		newBranch := c.String("new-branch")
		targetBranch := c.String("target-branch")

		if gitlabURL != "" && gitlabToken != "" && projectID != "" && baseBranch != "" && newBranch != "" && targetBranch != "" {
			err = actions.HandleGitLabMergeRequest(gitlabURL, gitlabToken, outputPath, baseBranch, newBranch, targetBranch, projectID)
			if err != nil {
				return fmt.Errorf("error handling GitLab merge request: %v", err)
			}
		}

	} else {
		fmt.Println("Run the command with --approved to apply these changes.")
	}

	return nil
}
