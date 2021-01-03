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

// GoogleStorageDefaultObjectAccessControlInvalidRoleRule checks the pattern is valid
type GoogleStorageDefaultObjectAccessControlInvalidRoleRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleStorageDefaultObjectAccessControlInvalidRoleRule returns new rule with default attributes
func NewGoogleStorageDefaultObjectAccessControlInvalidRoleRule() *GoogleStorageDefaultObjectAccessControlInvalidRoleRule {
	return &GoogleStorageDefaultObjectAccessControlInvalidRoleRule{
		resourceType:  "google_storage_default_object_access_control",
		attributeName: "role",
	}
}

// Name returns the rule name
func (r *GoogleStorageDefaultObjectAccessControlInvalidRoleRule) Name() string {
	return "google_storage_default_object_access_control_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleStorageDefaultObjectAccessControlInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleStorageDefaultObjectAccessControlInvalidRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleStorageDefaultObjectAccessControlInvalidRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleStorageDefaultObjectAccessControlInvalidRoleRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"OWNER", "READER"}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
