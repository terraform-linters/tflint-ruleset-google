package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GoogleProjectIamMemberInvalidMemberRule checks whether member value is invalid
type GoogleProjectIamMemberInvalidMemberRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleProjectIamMemberInvalidMemberRule returns new rule with default attributes
func NewGoogleProjectIamMemberInvalidMemberRule() *GoogleProjectIamMemberInvalidMemberRule {
	return &GoogleProjectIamMemberInvalidMemberRule{
		resourceType:  "google_project_iam_member",
		attributeName: "member",
	}
}

// Name returns the rule name
func (r *GoogleProjectIamMemberInvalidMemberRule) Name() string {
	return "google_project_iam_member_invalid_member"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleProjectIamMemberInvalidMemberRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleProjectIamMemberInvalidMemberRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleProjectIamMemberInvalidMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether member format is invalid
func (r *GoogleProjectIamMemberInvalidMemberRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {

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
