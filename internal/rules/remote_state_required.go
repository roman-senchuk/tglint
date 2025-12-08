package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// RemoteStateRequired checks that remote_state block exists
type RemoteStateRequired struct{}

func NewRemoteStateRequired() *RemoteStateRequired {
	return &RemoteStateRequired{}
}

func (r *RemoteStateRequired) Name() string {
	return "remote_state_required"
}

func (r *RemoteStateRequired) Check(file *hcl.File, filePath string) ([]Issue, error) {
	var issues []Issue

	body := file.Body.(*hclsyntax.Body)
	found := false

	for _, block := range body.Blocks {
		if block.Type == "remote_state" {
			found = true
			break
		}
	}

	if !found {
		issues = append(issues, Issue{
			File:    filePath,
			Line:    1,
			Column:  1,
			Message: "remote_state block is required",
			Rule:    r.Name(),
		})
	}

	return issues, nil
}
