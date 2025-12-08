package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tglint/tglint/internal/formatter"
	"github.com/tglint/tglint/internal/walker"
)

var checkFlag bool

var fmtCmd = &cobra.Command{
	Use:   "fmt [path]",
	Short: "Format terragrunt.hcl files",
	Long:  `Format terragrunt.hcl files recursively. Use --check to verify formatting without modifying files.`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runFmt,
}

func init() {
	fmtCmd.Flags().BoolVar(&checkFlag, "check", false, "Check if files are formatted (exit 2 if not)")
}

func runFmt(cmd *cobra.Command, args []string) error {
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
		if !checkFlag {
			fmt.Printf("No terragrunt.hcl files found in %s\n", absPath)
		}
		return nil
	}

	var unformatted []string
	var hasError bool
	var formattedCount int

	for _, file := range files {
		if checkFlag {
			formatted, err := formatter.FormatCheck(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", file, err)
				hasError = true
				continue
			}
			if !formatted {
				relPath, _ := filepath.Rel(absPath, file)
				fmt.Printf("ERROR: %s is not formatted\n", relPath)
				unformatted = append(unformatted, file)
			}
		} else {
			changed, err := formatter.FormatFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", file, err)
				hasError = true
				continue
			}
			if changed {
				relPath, _ := filepath.Rel(absPath, file)
				fmt.Printf("FORMAT %s\n", relPath)
				formattedCount++
			}
		}
	}

	// Print summary
	if !checkFlag && !hasError {
		if formattedCount > 0 {
			fmt.Printf("\nFormatted %d file(s) recursively\n", formattedCount)
		} else {
			fmt.Printf("\nAll %d file(s) are already formatted\n", len(files))
		}
	} else if checkFlag {
		if len(unformatted) == 0 {
			fmt.Printf("\nAll %d file(s) are formatted\n", len(files))
		} else {
			fmt.Printf("\n%d file(s) need formatting\n", len(unformatted))
		}
	}

	if checkFlag && len(unformatted) > 0 {
		os.Exit(2)
	}

	if hasError {
		os.Exit(3)
	}

	return nil
}
