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

// GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule checks the pattern is valid
type GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleDataCatalogEntryInvalidUserSpecifiedTypeRule returns new rule with default attributes
func NewGoogleDataCatalogEntryInvalidUserSpecifiedTypeRule() *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule {
	return &GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule{
		resourceType:  "google_data_catalog_entry",
		attributeName: "user_specified_type",
	}
}

// Name returns the rule name
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Name() string {
	return "google_data_catalog_entry_invalid_user_specified_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDataCatalogEntryInvalidUserSpecifiedTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validateRegexp(`^[A-z_][A-z0-9_]{0,63}$`)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
