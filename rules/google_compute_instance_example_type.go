package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeInstanceExampleTypeRule checks whether ...
type GoogleComputeInstanceExampleTypeRule struct{}

// NewGoogleComputeInstanceExampleTypeRule returns a new rule
func NewGoogleComputeInstanceExampleTypeRule() *GoogleComputeInstanceExampleTypeRule {
	return &GoogleComputeInstanceExampleTypeRule{}
}

// Name returns the rule name
func (r *GoogleComputeInstanceExampleTypeRule) Name() string {
	return "google_compute_instance_example_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeInstanceExampleTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeInstanceExampleTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeInstanceExampleTypeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *GoogleComputeInstanceExampleTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("google_compute_instance", "machine_type", func(attribute *hcl.Attribute) error {
		var machineType string
		err := runner.EvaluateExpr(attribute.Expr, &machineType)

		return runner.EnsureNoError(err, func() error {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("machine type is %s", machineType),
				attribute.Expr.Range(),
				tflint.Metadata{Expr: attribute.Expr},
			)
		})
	})
}
