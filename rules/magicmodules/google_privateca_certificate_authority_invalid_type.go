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

// GooglePrivatecaCertificateAuthorityInvalidTypeRule checks the pattern is valid
type GooglePrivatecaCertificateAuthorityInvalidTypeRule struct {
	resourceType  string
	attributeName string
}

// NewGooglePrivatecaCertificateAuthorityInvalidTypeRule returns new rule with default attributes
func NewGooglePrivatecaCertificateAuthorityInvalidTypeRule() *GooglePrivatecaCertificateAuthorityInvalidTypeRule {
	return &GooglePrivatecaCertificateAuthorityInvalidTypeRule{
		resourceType:  "google_privateca_certificate_authority",
		attributeName: "type",
	}
}

// Name returns the rule name
func (r *GooglePrivatecaCertificateAuthorityInvalidTypeRule) Name() string {
	return "google_privateca_certificate_authority_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GooglePrivatecaCertificateAuthorityInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GooglePrivatecaCertificateAuthorityInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GooglePrivatecaCertificateAuthorityInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GooglePrivatecaCertificateAuthorityInvalidTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"SELF_SIGNED", "SUBORDINATE", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
