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
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule checks the pattern is valid
type GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeBackendServiceSignedUrlKeyInvalidNameRule returns new rule with default attributes
func NewGoogleComputeBackendServiceSignedUrlKeyInvalidNameRule() *GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule {
	return &GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule{
		resourceType:  "google_compute_backend_service_signed_url_key",
		attributeName: "name",
	}
}

// Name returns the rule name
func (r *GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule) Name() string {
	return "google_compute_backend_service_signed_url_key_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeBackendServiceSignedUrlKeyInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validateRegexp(`^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$`)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
