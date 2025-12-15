package formatter

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Format formats a Terraform/Terragrunt HCL file
func Format(filePath string) ([]byte, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the file
	f, diags := hclwrite.ParseConfig(src, filePath, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL: %s", diags.Error())
	}

	// Format returns canonical HCL
	formatted := f.Bytes()
	
	// Ensure file ends with a newline
	if len(formatted) > 0 && formatted[len(formatted)-1] != '\n' {
		formatted = append(formatted, '\n')
	}
	
	return formatted, nil
}

// FormatFile formats a file and writes it back if changed
func FormatFile(filePath string) (bool, error) {
	formatted, err := Format(filePath)
	if err != nil {
		return false, err
	}

	// Read original file to compare
	original, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to read file: %w", err)
	}

	// Check if content changed
	changed := string(original) != string(formatted)

	if changed {
		if err := os.WriteFile(filePath, formatted, 0644); err != nil {
			return false, fmt.Errorf("failed to write file: %w", err)
		}
	}

	return changed, nil
}

// FormatCheck checks if a file is formatted without modifying it
func FormatCheck(filePath string) (bool, error) {
	formatted, err := Format(filePath)
	if err != nil {
		return false, err
	}

	original, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to read file: %w", err)
	}

	return string(original) == string(formatted), nil
}
