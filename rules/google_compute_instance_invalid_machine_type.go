package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeInstanceInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComputeInstanceInvalidMachineTypeRule struct {
	tflint.DefaultRule
}

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
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComputeInstanceInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_compute_instance", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: "machine_type"}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes["machine_type"]
		if !exists {
			continue
		}

		var machineType string
		err := runner.EvaluateExpr(attribute.Expr, &machineType, nil)

		err = runner.EnsureNoError(err, func() error {
			if validMachineTypes[machineType] || isCustomType(machineType) {
				return nil
			}

			return runner.EmitIssue(
				r,
				fmt.Sprintf(`"%s" is an invalid as machine type`, machineType),
				attribute.Expr.Range(),
			)
		})
		if err != nil {
			return err
		}
	}

	return nil
}
