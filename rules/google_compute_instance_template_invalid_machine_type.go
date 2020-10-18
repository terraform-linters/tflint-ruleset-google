package rules

import (
	"fmt"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeInstanceTemplateInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComputeInstanceTemplateInvalidMachineTypeRule struct {}

// NewGoogleComputeInstanceTemplateInvalidMachineTypeRule returns a new rule
func NewGoogleComputeInstanceTemplateInvalidMachineTypeRule() *GoogleComputeInstanceTemplateInvalidMachineTypeRule {
	return &GoogleComputeInstanceTemplateInvalidMachineTypeRule{}
}

// Name returns the rule name
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Name() string {
	return "google_compute_instance_template_invalid_machine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("google_compute_instance_template", "machine_type", func(attribute *hcl.Attribute) error {
		var machineType string
		err := runner.EvaluateExpr(attribute.Expr, &machineType)

		return runner.EnsureNoError(err, func() error {
			if validMachineTypes[machineType] ||
				strings.HasPrefix(machineType, "e2-custom-") ||
				strings.HasPrefix(machineType, "n2-custom-") ||
				strings.HasPrefix(machineType, "n2d-custom-") ||
				strings.HasPrefix(machineType, "n1-custom-") {
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
