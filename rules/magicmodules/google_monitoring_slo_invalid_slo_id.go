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

// GoogleMonitoringSloInvalidSloIdRule checks the pattern is valid
type GoogleMonitoringSloInvalidSloIdRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleMonitoringSloInvalidSloIdRule returns new rule with default attributes
func NewGoogleMonitoringSloInvalidSloIdRule() *GoogleMonitoringSloInvalidSloIdRule {
	return &GoogleMonitoringSloInvalidSloIdRule{
		resourceType:  "google_monitoring_slo",
		attributeName: "slo_id",
	}
}

// Name returns the rule name
func (r *GoogleMonitoringSloInvalidSloIdRule) Name() string {
	return "google_monitoring_slo_invalid_slo_id"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleMonitoringSloInvalidSloIdRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleMonitoringSloInvalidSloIdRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleMonitoringSloInvalidSloIdRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleMonitoringSloInvalidSloIdRule) Check(runner tflint.Runner) error {
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
			validateFunc := validateRegexp(`^[a-z0-9\-]+$`)

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
