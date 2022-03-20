// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule checks the pattern is valid
type GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule returns new rule with default attributes
func NewGoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule() *GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule {
	return &GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule{
		resourceType:  "google_compute_external_vpn_gateway",
		attributeName: "redundancy_type",
	}
}

// Name returns the rule name
func (r *GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule) Name() string {
	return "google_compute_external_vpn_gateway_invalid_redundancy_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"FOUR_IPS_REDUNDANCY", "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY", ""}, false)

		err = runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssue(r, err.Error(), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
