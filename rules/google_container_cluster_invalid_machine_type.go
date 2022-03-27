package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleContainerClusterInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleContainerClusterInvalidMachineTypeRule struct {
	tflint.DefaultRule
}

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
func (r *GoogleContainerClusterInvalidMachineTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleContainerClusterInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleContainerClusterInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_container_cluster", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "node_config",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: "machine_type"}},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, nodeConfig := range resource.Body.Blocks {
			attribute, exists := nodeConfig.Body.Attributes["machine_type"]
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
	}

	return nil
}
