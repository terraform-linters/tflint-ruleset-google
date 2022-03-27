package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleDataflowJobInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleDataflowJobInvalidMachineTypeRule struct {
	tflint.DefaultRule
}

// NewGoogleDataflowJobInvalidMachineTypeRule returns a new rule
func NewGoogleDataflowJobInvalidMachineTypeRule() *GoogleDataflowJobInvalidMachineTypeRule {
	return &GoogleDataflowJobInvalidMachineTypeRule{}
}

// Name returns the rule name
func (r *GoogleDataflowJobInvalidMachineTypeRule) Name() string {
	return "google_dataflow_job_invalid_machine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDataflowJobInvalidMachineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDataflowJobInvalidMachineTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDataflowJobInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleDataflowJobInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_dataflow_job", &hclext.BodySchema{
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
