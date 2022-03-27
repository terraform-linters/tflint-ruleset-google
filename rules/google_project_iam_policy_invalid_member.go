package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GoogleProjectIamPolicyInvalidMemberRule checks whether member value is invalid
type GoogleProjectIamPolicyInvalidMemberRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleProjectIamPolicyInvalidMemberRule returns new rule with default attributes
func NewGoogleProjectIamPolicyInvalidMemberRule() *GoogleProjectIamPolicyInvalidMemberRule {
	return &GoogleProjectIamPolicyInvalidMemberRule{
		resourceType:  "google_project_iam_policy",
		attributeName: "members",
	}
}

// Name returns the rule name
func (r *GoogleProjectIamPolicyInvalidMemberRule) Name() string {
	return "google_project_iam_policy_invalid_member"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleProjectIamPolicyInvalidMemberRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleProjectIamPolicyInvalidMemberRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleProjectIamPolicyInvalidMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether member format is invalid
func (r *GoogleProjectIamPolicyInvalidMemberRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var members []string
		err := runner.EvaluateExpr(attribute.Expr, &members, nil)

		err = runner.EnsureNoError(err, func() error {
			for _, member := range members {
				if !isValidIAMMemberFormat(member) {
					return runner.EmitIssue(
						r,
						fmt.Sprintf("%s is an invalid member format", member),
						attribute.Expr.Range(),
					)
				}

			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
