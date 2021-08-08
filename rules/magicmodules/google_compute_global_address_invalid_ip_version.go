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
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeGlobalAddressInvalidIpVersionRule checks the pattern is valid
type GoogleComputeGlobalAddressInvalidIpVersionRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeGlobalAddressInvalidIpVersionRule returns new rule with default attributes
func NewGoogleComputeGlobalAddressInvalidIpVersionRule() *GoogleComputeGlobalAddressInvalidIpVersionRule {
	return &GoogleComputeGlobalAddressInvalidIpVersionRule{
		resourceType:  "google_compute_global_address",
		attributeName: "ip_version",
	}
}

// Name returns the rule name
func (r *GoogleComputeGlobalAddressInvalidIpVersionRule) Name() string {
	return "google_compute_global_address_invalid_ip_version"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeGlobalAddressInvalidIpVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeGlobalAddressInvalidIpVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeGlobalAddressInvalidIpVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeGlobalAddressInvalidIpVersionRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"IPV4", "IPV6", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
