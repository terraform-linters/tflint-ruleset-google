package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleProjectIamMemberInvalidMemberFormatRule checks whether member value is invalid
type GoogleProjectIamMemberInvalidMemberFormatRule struct{}

// NewGoogleProjectIamMemberInvalidMemberFormatRule returns new rule with default attributes
func NewGoogleProjectIamMemberInvalidMemberFormatRule() *GoogleProjectIamMemberInvalidMemberFormatRule {
	return &GoogleProjectIamMemberInvalidMemberFormatRule{}
}

// Name returns the rule name
func (r *GoogleProjectIamMemberInvalidMemberFormatRule) Name() string {
	return "google_project_iam_member_invalid_member_format"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleProjectIamMemberInvalidMemberFormatRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleProjectIamMemberInvalidMemberFormatRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleProjectIamMemberInvalidMemberFormatRule) Link() string {
	return iamMemberFormatDocumentLink
}

// Check checks whether member format is invalid
func (r *GoogleProjectIamMemberInvalidMemberFormatRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("google_project_iam_member", "member", func(attribute *hcl.Attribute) error {

		var member string
		err := runner.EvaluateExpr(attribute.Expr, &member, nil)

		return runner.EnsureNoError(err, func() error {
			if isValidIAMMemberFormat(member) {
				return nil
			}
			return runner.EmitIssueOnExpr(
				r,
				fmt.Sprintf("%s is an invalid member format", member),
				attribute.Expr,
			)
		})
	})
}
