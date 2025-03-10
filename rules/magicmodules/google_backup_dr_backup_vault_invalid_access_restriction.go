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

// GoogleBackupDrBackupVaultInvalidAccessRestrictionRule checks the pattern is valid
type GoogleBackupDrBackupVaultInvalidAccessRestrictionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleBackupDrBackupVaultInvalidAccessRestrictionRule returns new rule with default attributes
func NewGoogleBackupDrBackupVaultInvalidAccessRestrictionRule() *GoogleBackupDrBackupVaultInvalidAccessRestrictionRule {
	return &GoogleBackupDrBackupVaultInvalidAccessRestrictionRule{
		resourceType:  "google_backup_dr_backup_vault",
		attributeName: "access_restriction",
	}
}

// Name returns the rule name
func (r *GoogleBackupDrBackupVaultInvalidAccessRestrictionRule) Name() string {
	return "google_backup_dr_backup_vault_invalid_access_restriction"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleBackupDrBackupVaultInvalidAccessRestrictionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleBackupDrBackupVaultInvalidAccessRestrictionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleBackupDrBackupVaultInvalidAccessRestrictionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleBackupDrBackupVaultInvalidAccessRestrictionRule) Check(runner tflint.Runner) error {
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
			validateFunc := validation.StringInSlice([]string{"ACCESS_RESTRICTION_UNSPECIFIED", "WITHIN_PROJECT", "WITHIN_ORGANIZATION", "UNRESTRICTED", "WITHIN_ORG_BUT_UNRESTRICTED_FOR_BA", ""}, false)

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
