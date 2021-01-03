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
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule checks the pattern is valid
type GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleDataCatalogTagTemplateInvalidTagTemplateIdRule returns new rule with default attributes
func NewGoogleDataCatalogTagTemplateInvalidTagTemplateIdRule() *GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule {
	return &GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule{
		resourceType:  "google_data_catalog_tag_template",
		attributeName: "tag_template_id",
	}
}

// Name returns the rule name
func (r *GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule) Name() string {
	return "google_data_catalog_tag_template_invalid_tag_template_id"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDataCatalogTagTemplateInvalidTagTemplateIdRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validateRegexp(`^[a-z_][a-z0-9_]{0,63}$`)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
