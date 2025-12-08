package rules

import (
	"github.com/hashicorp/hcl/v2"
)

// Rule defines a linting rule
type Rule interface {
	Name() string
	Check(file *hcl.File, filePath string) ([]Issue, error)
}
