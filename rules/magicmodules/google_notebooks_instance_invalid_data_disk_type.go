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

// GoogleNotebooksInstanceInvalidDataDiskTypeRule checks the pattern is valid
type GoogleNotebooksInstanceInvalidDataDiskTypeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleNotebooksInstanceInvalidDataDiskTypeRule returns new rule with default attributes
func NewGoogleNotebooksInstanceInvalidDataDiskTypeRule() *GoogleNotebooksInstanceInvalidDataDiskTypeRule {
	return &GoogleNotebooksInstanceInvalidDataDiskTypeRule{
		resourceType:  "google_notebooks_instance",
		attributeName: "data_disk_type",
	}
}

// Name returns the rule name
func (r *GoogleNotebooksInstanceInvalidDataDiskTypeRule) Name() string {
	return "google_notebooks_instance_invalid_data_disk_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleNotebooksInstanceInvalidDataDiskTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleNotebooksInstanceInvalidDataDiskTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleNotebooksInstanceInvalidDataDiskTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleNotebooksInstanceInvalidDataDiskTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
