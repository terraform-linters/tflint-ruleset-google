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

// GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule checks the pattern is valid
type GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule returns new rule with default attributes
func NewGoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule() *GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule {
	return &GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule{
		resourceType:  "google_compute_node_template",
		attributeName: "cpu_overcommit_type",
	}
}

// Name returns the rule name
func (r *GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule) Name() string {
	return "google_compute_node_template_invalid_cpu_overcommit_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"ENABLED", "NONE", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
