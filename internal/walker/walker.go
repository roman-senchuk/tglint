package walker

import (
	"os"
	"path/filepath"
	"strings"

	gitignore "github.com/sabhiram/go-gitignore"
)

// WalkOptions configures file walking behavior
type WalkOptions struct {
	RootPath      string
	GitignorePath string
	TglintIgnore  string
}

// Walk finds all terragrunt.hcl and .tf files recursively
func Walk(opts WalkOptions) ([]string, error) {
	var files []string
	var gitignoreMatcher *gitignore.GitIgnore
	var tglintIgnoreMatcher *gitignore.GitIgnore

	// Load .gitignore if it exists
	if opts.GitignorePath != "" {
		if _, err := os.Stat(opts.GitignorePath); err == nil {
			gitignoreMatcher, _ = gitignore.CompileIgnoreFile(opts.GitignorePath)
		}
	}

	// Load .tglintignore if provided
	if opts.TglintIgnore != "" {
		tglintIgnoreMatcher = gitignore.CompileIgnoreLines(strings.Split(opts.TglintIgnore, "\n")...)
	}

	err := filepath.Walk(opts.RootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			// Skip .terraform and .terragrunt-cache directories (and any path containing them)
			base := filepath.Base(path)
			if base == ".terraform" || base == ".terragrunt-cache" {
				return filepath.SkipDir
			}
			// Also check if the path contains these directories anywhere
			if strings.Contains(path, "/.terraform/") || strings.Contains(path, "/.terragrunt-cache/") ||
				strings.Contains(path, "\\.terraform\\") || strings.Contains(path, "\\.terragrunt-cache\\") {
				return filepath.SkipDir
			}
			return nil
		}

		// Process terragrunt.hcl and .tf files
		name := info.Name()
		isTerragrunt := name == "terragrunt.hcl"
		isTerraform := strings.HasSuffix(name, ".tf")
		
		if !isTerragrunt && !isTerraform {
			return nil
		}

		relPath, err := filepath.Rel(opts.RootPath, path)
		if err != nil {
			return err
		}

		// Check .gitignore
		if gitignoreMatcher != nil {
			if gitignoreMatcher.MatchesPath(relPath) {
				return nil
			}
		}

		// Check .tglintignore
		if tglintIgnoreMatcher != nil {
			if tglintIgnoreMatcher.MatchesPath(relPath) {
				return nil
			}
		}

		files = append(files, path)
		return nil
	})

	return files, err
}
