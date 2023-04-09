package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GoogleProjectIamBindingInvalidMemberRule checks whether member value is invalid
type GoogleProjectIamBindingInvalidMemberRule struct {
	tflint.DefaultRule

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
func (r *GoogleProjectIamBindingInvalidMemberRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleProjectIamBindingInvalidMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether member format is invalid
func (r *GoogleProjectIamBindingInvalidMemberRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(members []string) error {
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
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
