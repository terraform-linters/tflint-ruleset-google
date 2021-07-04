// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleApigeeInstanceInvalidPeeringCidrRangeRule checks the pattern is valid
type GoogleApigeeInstanceInvalidPeeringCidrRangeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleApigeeInstanceInvalidPeeringCidrRangeRule returns new rule with default attributes
func NewGoogleApigeeInstanceInvalidPeeringCidrRangeRule() *GoogleApigeeInstanceInvalidPeeringCidrRangeRule {
	return &GoogleApigeeInstanceInvalidPeeringCidrRangeRule{
		resourceType:  "google_apigee_instance",
		attributeName: "peering_cidr_range",
	}
}

// Name returns the rule name
func (r *GoogleApigeeInstanceInvalidPeeringCidrRangeRule) Name() string {
	return "google_apigee_instance_invalid_peering_cidr_range"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleApigeeInstanceInvalidPeeringCidrRangeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleApigeeInstanceInvalidPeeringCidrRangeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleApigeeInstanceInvalidPeeringCidrRangeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleApigeeInstanceInvalidPeeringCidrRangeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"SLASH_16", "SLASH_20", "SLASH_23", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
