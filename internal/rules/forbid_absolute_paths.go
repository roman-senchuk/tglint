package rules

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ForbidAbsolutePaths checks for absolute paths in terraform.source
type ForbidAbsolutePaths struct{}

func NewForbidAbsolutePaths() *ForbidAbsolutePaths {
	return &ForbidAbsolutePaths{}
}

func (r *ForbidAbsolutePaths) Name() string {
	return "forbid_absolute_paths"
}

func (r *ForbidAbsolutePaths) Check(file *hcl.File, filePath string) ([]Issue, error) {
	var issues []Issue

	body := file.Body.(*hclsyntax.Body)

	// Find terraform block
	for _, block := range body.Blocks {
		if block.Type == "terraform" {
		// Check source attribute
		if sourceAttr, exists := block.Body.Attributes["source"]; exists {
			var checkExpr func(expr hclsyntax.Expression) bool
			checkExpr = func(expr hclsyntax.Expression) bool {
				switch e := expr.(type) {
				case *hclsyntax.TemplateExpr:
					// Check template parts
					for _, part := range e.Parts {
						if litExpr, ok := part.(*hclsyntax.LiteralValueExpr); ok {
							if litExpr.Val.Type() == cty.String {
								source := litExpr.Val.AsString()
								if strings.HasPrefix(source, "/") {
									return true
								}
							}
						}
					}
				case *hclsyntax.LiteralValueExpr:
					if e.Val.Type() == cty.String {
						source := e.Val.AsString()
						if strings.HasPrefix(source, "/") {
							return true
						}
					}
				}
				return false
			}

			if checkExpr(sourceAttr.Expr) {
				issues = append(issues, Issue{
					File:    filePath,
					Line:    sourceAttr.SrcRange.Start.Line,
					Column:  sourceAttr.SrcRange.Start.Column,
					Message: "absolute paths in terraform.source are not allowed",
					Rule:    r.Name(),
				})
			}
		}
			break
		}
	}

	return issues, nil
}
