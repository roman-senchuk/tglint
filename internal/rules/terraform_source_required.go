package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// TerraformSourceRequired checks that terraform.source is set
type TerraformSourceRequired struct{}

func NewTerraformSourceRequired() *TerraformSourceRequired {
	return &TerraformSourceRequired{}
}

func (r *TerraformSourceRequired) Name() string {
	return "terraform_source_required"
}

func (r *TerraformSourceRequired) Check(file *hcl.File, filePath string) ([]Issue, error) {
	var issues []Issue

	body := file.Body.(*hclsyntax.Body)

		// Find terraform block
		for _, block := range body.Blocks {
			if block.Type == "terraform" {
				// Check if source attribute exists
				_, exists := block.Body.Attributes["source"]
				if !exists {
					issues = append(issues, Issue{
						File:    filePath,
						Line:    block.TypeRange.Start.Line,
						Column:  block.TypeRange.Start.Column,
						Message: "terraform.source is required",
						Rule:    r.Name(),
					})
				}
				break
			}
		}

	return issues, nil
}
