package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// DisallowEmptyInputs checks for empty inputs = {}
type DisallowEmptyInputs struct{}

func NewDisallowEmptyInputs() *DisallowEmptyInputs {
	return &DisallowEmptyInputs{}
}

func (r *DisallowEmptyInputs) Name() string {
	return "disallow_empty_inputs"
}

func (r *DisallowEmptyInputs) Check(file *hcl.File, filePath string) ([]Issue, error) {
	var issues []Issue

	body := file.Body.(*hclsyntax.Body)

	// Check inputs attribute
	if inputsAttr, exists := body.Attributes["inputs"]; exists {
		// Check if it's an empty object
		if objCons, ok := inputsAttr.Expr.(*hclsyntax.ObjectConsExpr); ok {
			if len(objCons.Items) == 0 {
				issues = append(issues, Issue{
					File:    filePath,
					Line:    inputsAttr.SrcRange.Start.Line,
					Column:  inputsAttr.SrcRange.Start.Column,
					Message: "empty inputs = {} is not allowed",
					Rule:    r.Name(),
				})
			}
		}
	}

	return issues, nil
}
