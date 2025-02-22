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
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleSccV2OrganizationSourceInvalidDisplayNameRule checks the pattern is valid
type GoogleSccV2OrganizationSourceInvalidDisplayNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleSccV2OrganizationSourceInvalidDisplayNameRule returns new rule with default attributes
func NewGoogleSccV2OrganizationSourceInvalidDisplayNameRule() *GoogleSccV2OrganizationSourceInvalidDisplayNameRule {
	return &GoogleSccV2OrganizationSourceInvalidDisplayNameRule{
		resourceType:  "google_scc_v2_organization_source",
		attributeName: "display_name",
	}
}

// Name returns the rule name
func (r *GoogleSccV2OrganizationSourceInvalidDisplayNameRule) Name() string {
	return "google_scc_v2_organization_source_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleSccV2OrganizationSourceInvalidDisplayNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleSccV2OrganizationSourceInvalidDisplayNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleSccV2OrganizationSourceInvalidDisplayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleSccV2OrganizationSourceInvalidDisplayNameRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			validateFunc := validateRegexp(`[\p{L}\p{N}]({\p{L}\p{N}_- ]{0,30}[\p{L}\p{N}])?`)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
