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

// GoogleMemorystoreInstanceInvalidModeRule checks the pattern is valid
type GoogleMemorystoreInstanceInvalidModeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleMemorystoreInstanceInvalidModeRule returns new rule with default attributes
func NewGoogleMemorystoreInstanceInvalidModeRule() *GoogleMemorystoreInstanceInvalidModeRule {
	return &GoogleMemorystoreInstanceInvalidModeRule{
		resourceType:  "google_memorystore_instance",
		attributeName: "mode",
	}
}

// Name returns the rule name
func (r *GoogleMemorystoreInstanceInvalidModeRule) Name() string {
	return "google_memorystore_instance_invalid_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleMemorystoreInstanceInvalidModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleMemorystoreInstanceInvalidModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleMemorystoreInstanceInvalidModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleMemorystoreInstanceInvalidModeRule) Check(runner tflint.Runner) error {
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
			validateFunc := validation.StringInSlice([]string{"CLUSTER", "CLUSTER_DISABLED", ""}, false)

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
