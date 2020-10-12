package rules

import (
	"fmt"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleDataflowJobInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleDataflowJobInvalidMachineTypeRule struct {}

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
func (r *GoogleDataflowJobInvalidMachineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDataflowJobInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleDataflowJobInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("google_dataflow_job", "machine_type", func(attribute *hcl.Attribute) error {
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
