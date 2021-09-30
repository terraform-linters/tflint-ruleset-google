package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GoogleProjectIamBindingInvalidMemberRule checks whether member value is invalid
type GoogleProjectIamBindingInvalidMemberRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleProjectIamBindingInvalidMemberRule returns new rule with default attributes
func NewGoogleProjectIamBindingInvalidMemberRule() *GoogleProjectIamBindingInvalidMemberRule {
	return &GoogleProjectIamBindingInvalidMemberRule{
		resourceType:  "google_project_iam_binding",
		attributeName: "members",
	}
}

// Name returns the rule name
func (r *GoogleProjectIamBindingInvalidMemberRule) Name() string {
	return "google_project_iam_binding_invalid_member"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleProjectIamBindingInvalidMemberRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleProjectIamBindingInvalidMemberRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleProjectIamBindingInvalidMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether member format is invalid
func (r *GoogleProjectIamBindingInvalidMemberRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var members []string
		err := runner.EvaluateExpr(attribute.Expr, &members, nil)

		return runner.EnsureNoError(err, func() error {
			for _, member := range members {
				if !isValidIAMMemberFormat(member) {
					return runner.EmitIssueOnExpr(
						r,
						fmt.Sprintf("%s is an invalid member format", member),
						attribute.Expr,
					)
				}
			}
			return nil
		})
	})
}
