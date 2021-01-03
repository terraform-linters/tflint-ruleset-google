package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeReservationInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComputeReservationInvalidMachineTypeRule struct{}

// NewGoogleComputeReservationInvalidMachineTypeRule returns a new rule
func NewGoogleComputeReservationInvalidMachineTypeRule() *GoogleComputeReservationInvalidMachineTypeRule {
	return &GoogleComputeReservationInvalidMachineTypeRule{}
}

// Name returns the rule name
func (r *GoogleComputeReservationInvalidMachineTypeRule) Name() string {
	return "google_compute_reservation_invalid_machine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeReservationInvalidMachineTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeReservationInvalidMachineTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeReservationInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComputeReservationInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceBlocks("google_compute_reservation", "specific_reservation", func(block *hcl.Block) error {
		content, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Blocks: []hcl.BlockHeaderSchema{
				{Type: "instance_properties"},
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
				err := runner.EvaluateExpr(attribute.Expr, &machineType, nil)

				err = runner.EnsureNoError(err, func() error {
					if validMachineTypes[machineType] || isCustomType(machineType) {
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
