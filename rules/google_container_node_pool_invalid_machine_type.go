package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleContainerNodePoolInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleContainerNodePoolInvalidMachineTypeRule struct{}

// NewGoogleContainerNodePoolInvalidMachineTypeRule returns a new rule
func NewGoogleContainerNodePoolInvalidMachineTypeRule() *GoogleContainerNodePoolInvalidMachineTypeRule {
	return &GoogleContainerNodePoolInvalidMachineTypeRule{}
}

// Name returns the rule name
func (r *GoogleContainerNodePoolInvalidMachineTypeRule) Name() string {
	return "google_container_node_pool_invalid_machine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleContainerNodePoolInvalidMachineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleContainerNodePoolInvalidMachineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleContainerNodePoolInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleContainerNodePoolInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceBlocks("google_container_node_pool", "node_config", func(block *hcl.Block) error {
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
