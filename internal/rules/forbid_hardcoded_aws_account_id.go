package rules

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// ForbidHardcodedAWSAccountID checks for hardcoded 12-digit AWS account IDs
type ForbidHardcodedAWSAccountID struct {
	accountIDRegex *regexp.Regexp
}

func NewForbidHardcodedAWSAccountID() *ForbidHardcodedAWSAccountID {
	return &ForbidHardcodedAWSAccountID{
		accountIDRegex: regexp.MustCompile(`\b\d{12}\b`),
	}
}

func (r *ForbidHardcodedAWSAccountID) Name() string {
	return "forbid_hardcoded_aws_account_id"
}

func (r *ForbidHardcodedAWSAccountID) Check(file *hcl.File, filePath string) ([]Issue, error) {
	var issues []Issue

	body := file.Body.(*hclsyntax.Body)

	// Check all attributes for hardcoded account IDs
	var checkExpr func(expr hclsyntax.Expression) error
	checkExpr = func(expr hclsyntax.Expression) error {
		switch e := expr.(type) {
		case *hclsyntax.TemplateExpr:
			for _, part := range e.Parts {
				if err := checkExpr(part); err != nil {
					return err
				}
			}
		case *hclsyntax.LiteralValueExpr:
			val := e.Val
			// Only check string values for AWS account IDs
			if val.Type() == cty.String {
				str := val.AsString()
				if r.accountIDRegex.MatchString(str) {
					// Verify it's actually a 12-digit number
					if num, err := strconv.ParseInt(str, 10, 64); err == nil && num >= 100000000000 && num <= 999999999999 {
						issues = append(issues, Issue{
							File:    filePath,
							Line:    e.SrcRange.Start.Line,
							Column:  e.SrcRange.Start.Column,
							Message: "hardcoded AWS account ID detected (12 digits)",
							Rule:    r.Name(),
						})
					}
				}
			}
		case *hclsyntax.ObjectConsExpr:
			for _, item := range e.Items {
				if err := checkExpr(item.ValueExpr); err != nil {
					return err
				}
			}
		case *hclsyntax.TupleConsExpr:
			for _, expr := range e.Exprs {
				if err := checkExpr(expr); err != nil {
					return err
				}
			}
		}
		return nil
	}

	// Check all blocks and attributes
	for _, block := range body.Blocks {
		for _, attr := range block.Body.Attributes {
			_ = checkExpr(attr.Expr)
		}
	}

	for _, attr := range body.Attributes {
		_ = checkExpr(attr.Expr)
	}

	return issues, nil
}
