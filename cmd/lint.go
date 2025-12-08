package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tglint/tglint/internal/linter"
	"github.com/tglint/tglint/internal/rules"
	"github.com/tglint/tglint/internal/walker"
)

var skipRulesFlag string

var lintCmd = &cobra.Command{
	Use:   "lint [path]",
	Short: "Lint terragrunt.hcl files",
	Long:  `Lint terragrunt.hcl files recursively and report violations.`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runLint,
}

func init() {
	lintCmd.Flags().StringVar(&skipRulesFlag, "skip-rules", "", "Comma-separated list of rule names to skip (e.g., remote_state_required,forbid_absolute_paths)")
}

func runLint(cmd *cobra.Command, args []string) error {
	rootPath := "."
	if len(args) > 0 {
		rootPath = args[0]
	}

	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	// Find .gitignore in root or parent directories
	gitignorePath := findGitignore(absPath)

	opts := walker.WalkOptions{
		RootPath:      absPath,
		GitignorePath: gitignorePath,
	}

	files, err := walker.Walk(opts)
	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	if len(files) == 0 {
		return nil
	}

	// Parse skipped rules
	var skippedRules map[string]bool
	if skipRulesFlag != "" {
		skippedRules = make(map[string]bool)
		for _, rule := range strings.Split(skipRulesFlag, ",") {
			skippedRules[strings.TrimSpace(rule)] = true
		}
	}

	l := linter.NewWithSkipRules(skippedRules)
	var allIssues []rules.Issue
	var hasError bool

	for _, file := range files {
		issues, err := l.LintFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", file, err)
			hasError = true
			continue
		}

		if len(issues) > 0 {
			allIssues = append(allIssues, issues...)
		}
	}

	// Print all issues
	for _, issue := range allIssues {
		relPath, _ := filepath.Rel(absPath, issue.File)
		fmt.Printf("%s:%d:%d: %s (%s)\n", relPath, issue.Line, issue.Column, issue.Message, issue.Rule)
	}

	if len(allIssues) > 0 {
		os.Exit(1)
	}

	if hasError {
		os.Exit(3)
	}

	return nil
}
