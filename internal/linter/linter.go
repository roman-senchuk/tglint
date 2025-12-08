package linter

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/roman-senchuk/tglint/internal/rules"
)

// Linter runs linting rules on HCL files
type Linter struct {
	rules []rules.Rule
}

// New creates a new linter with default rules
func New() *Linter {
	return NewWithSkipRules(nil)
}

// NewWithSkipRules creates a new linter with default rules, skipping specified rules
func NewWithSkipRules(skippedRules map[string]bool) *Linter {
	allRules := []rules.Rule{
		rules.NewRemoteStateRequired(),
		rules.NewTerraformSourceRequired(),
		rules.NewForbidHardcodedAWSAccountID(),
		rules.NewDisallowEmptyInputs(),
		rules.NewForbidAbsolutePaths(),
	}

	var activeRules []rules.Rule
	for _, rule := range allRules {
		if skippedRules == nil || !skippedRules[rule.Name()] {
			activeRules = append(activeRules, rule)
		}
	}

	return &Linter{
		rules: activeRules,
	}
}

// LintFile lints a single file
func (l *Linter) LintFile(filePath string) ([]rules.Issue, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	file, diags := hclsyntax.ParseConfig(src, filePath, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL: %s", diags.Error())
	}

	var issues []rules.Issue

	for _, rule := range l.rules {
		ruleIssues, err := rule.Check(file, filePath)
		if err != nil {
			return nil, fmt.Errorf("rule %s failed: %w", rule.Name(), err)
		}
		issues = append(issues, ruleIssues...)
	}

	return issues, nil
}
