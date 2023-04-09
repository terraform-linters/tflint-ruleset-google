package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeResourcePolicyInvalidNameRule checks whether the name is invalid
type GoogleComputeResourcePolicyInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleComputeResourcePolicyInvalidNameRule returns new rule with default attributes
func NewGoogleComputeResourcePolicyInvalidNameRule() *GoogleComputeResourcePolicyInvalidNameRule {
	return &GoogleComputeResourcePolicyInvalidNameRule{
		resourceType:  "google_compute_resource_policy",
		attributeName: "name",
	}
}

// Name returns the rule name
func (r *GoogleComputeResourcePolicyInvalidNameRule) Name() string {
	return "google_compute_resource_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeResourcePolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeResourcePolicyInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeResourcePolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks whether the name is invalid
func (r *GoogleComputeResourcePolicyInvalidNameRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			validateFunc := validateRegexp(`^[a-z]([-a-z0-9]*[a-z0-9])$`)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
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
