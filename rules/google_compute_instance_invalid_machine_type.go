package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeInstanceInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComputeInstanceInvalidMachineTypeRule struct{}

// NewGoogleComputeInstanceInvalidMachineTypeRule returns a new rule
func NewGoogleComputeInstanceInvalidMachineTypeRule() *GoogleComputeInstanceInvalidMachineTypeRule {
	return &GoogleComputeInstanceInvalidMachineTypeRule{}
}

// Name returns the rule name
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Name() string {
	return "google_compute_instance_invalid_machine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("google_compute_instance", "machine_type", func(attribute *hcl.Attribute) error {
		var machineType string
		err := runner.EvaluateExpr(attribute.Expr, &machineType, nil)

		return runner.EnsureNoError(err, func() error {
			if validMachineTypes[machineType] || isCustomType(machineType) {
				return nil
			}

			return runner.EmitIssueOnExpr(
				r,
				fmt.Sprintf(`"%s" is an invalid as machine type`, machineType),
				attribute.Expr,
			)
		})
	})
}
