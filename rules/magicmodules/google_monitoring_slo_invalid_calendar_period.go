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

// GoogleMonitoringSloInvalidCalendarPeriodRule checks the pattern is valid
type GoogleMonitoringSloInvalidCalendarPeriodRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleMonitoringSloInvalidCalendarPeriodRule returns new rule with default attributes
func NewGoogleMonitoringSloInvalidCalendarPeriodRule() *GoogleMonitoringSloInvalidCalendarPeriodRule {
	return &GoogleMonitoringSloInvalidCalendarPeriodRule{
		resourceType:  "google_monitoring_slo",
		attributeName: "calendar_period",
	}
}

// Name returns the rule name
func (r *GoogleMonitoringSloInvalidCalendarPeriodRule) Name() string {
	return "google_monitoring_slo_invalid_calendar_period"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleMonitoringSloInvalidCalendarPeriodRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleMonitoringSloInvalidCalendarPeriodRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleMonitoringSloInvalidCalendarPeriodRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleMonitoringSloInvalidCalendarPeriodRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"DAY", "WEEK", "FORTNIGHT", "MONTH", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
