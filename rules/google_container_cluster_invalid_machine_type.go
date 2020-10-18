package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleContainerClusterInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleContainerClusterInvalidMachineTypeRule struct{}

// NewGoogleContainerClusterInvalidMachineTypeRule returns a new rule
func NewGoogleContainerClusterInvalidMachineTypeRule() *GoogleContainerClusterInvalidMachineTypeRule {
	return &GoogleContainerClusterInvalidMachineTypeRule{}
}

// Name returns the rule name
func (r *GoogleContainerClusterInvalidMachineTypeRule) Name() string {
	return "google_container_cluster_invalid_machine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleContainerClusterInvalidMachineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleContainerClusterInvalidMachineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleContainerClusterInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleContainerClusterInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceBlocks("google_container_cluster", "node_config", func(block *hcl.Block) error {
		content, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{Name: "machine_type"},
			},
		})
		if diags.HasErrors() {
			return diags
		}

		if attribute, exists := content.Attributes["machine_type"]; exists {
			var machineType string
			err := runner.EvaluateExpr(attribute.Expr, &machineType)

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
		}

		return nil
	})
}
