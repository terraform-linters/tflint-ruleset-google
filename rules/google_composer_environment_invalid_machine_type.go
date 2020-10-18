package rules

import (
	"fmt"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComposerEnvironmentInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComposerEnvironmentInvalidMachineTypeRule struct{}

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
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComposerEnvironmentInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceBlocks("google_composer_environment", "config", func(block *hcl.Block) error {
		content, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Blocks: []hcl.BlockHeaderSchema{
				{Type: "node_config"},
			},
		})
		if diags.HasErrors() {
			return diags
		}

		for _, block := range content.Blocks {
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

				err = runner.EnsureNoError(err, func() error {
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

				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
