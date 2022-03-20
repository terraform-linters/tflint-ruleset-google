package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeReservationInvalidMachineTypeRule checks whether the machine type is invalid
type GoogleComputeReservationInvalidMachineTypeRule struct {
	tflint.DefaultRule
}

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
func (r *GoogleComputeReservationInvalidMachineTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeReservationInvalidMachineTypeRule) Link() string {
	return ""
}

// Check checks whether the machine type is invalid
func (r *GoogleComputeReservationInvalidMachineTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("google_compute_reservation", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "specific_reservation",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type: "instance_properties",
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
		for _, reservations := range resource.Body.Blocks {
			for _, properties := range reservations.Body.Blocks {
				attribute, exists := properties.Body.Attributes["machine_type"]
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
	}

	return nil
}
