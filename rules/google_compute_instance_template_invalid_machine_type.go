package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeInstanceTemplateInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComputeInstanceTemplateInvalidMachineTypeRule struct {
	tflint.DefaultRule
}

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
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComputeInstanceTemplateInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_compute_instance_template", &hclext.BodySchema{
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

		err := runner.EvaluateExpr(attribute.Expr, func(machineType string) error {
			if validMachineTypes[machineType] || isCustomType(machineType) {
				return nil
			}

			return runner.EmitIssue(
				r,
				fmt.Sprintf(`"%s" is an invalid as machine type`, machineType),
				attribute.Expr.Range(),
			)
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
