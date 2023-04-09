package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComposerEnvironmentInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComposerEnvironmentInvalidMachineTypeRule struct {
	tflint.DefaultRule
}

// NewGoogleComposerEnvironmentInvalidMachineTypeRule returns a new rule
func NewGoogleComposerEnvironmentInvalidMachineTypeRule() *GoogleComposerEnvironmentInvalidMachineTypeRule {
	return &GoogleComposerEnvironmentInvalidMachineTypeRule{}
}

// Name returns the rule name
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Name() string {
	return "google_composer_environment_invalid_machine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_composer_environment", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "config",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type: "node_config",
							Body: &hclext.BodySchema{
								Attributes: []hclext.AttributeSchema{{Name: "machine_type"}},
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, config := range resource.Body.Blocks {
			for _, nodeConfig := range config.Body.Blocks {
				attribute, exists := nodeConfig.Body.Attributes["machine_type"]
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
		}
	}

	return nil
}
