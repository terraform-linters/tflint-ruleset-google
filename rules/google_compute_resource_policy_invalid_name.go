package rules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeResourcePolicyInvalidNameRule checks whether the name is invalid
type GoogleComputeResourcePolicyInvalidNameRule struct {
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
func (r *GoogleComputeResourcePolicyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeResourcePolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks whether the name is invalid
func (r *GoogleComputeResourcePolicyInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validateRegexp(`^[a-z]([-a-z0-9]*[a-z0-9])$`)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}