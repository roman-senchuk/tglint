package cmd

import (
	"os"
	"path/filepath"
)

func findGitignore(startPath string) string {
	dir := startPath
	for {
		gitignorePath := filepath.Join(dir, ".gitignore")
		if _, err := os.Stat(gitignorePath); err == nil {
			return gitignorePath
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}
