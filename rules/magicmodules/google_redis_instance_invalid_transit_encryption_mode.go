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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleRedisInstanceInvalidTransitEncryptionModeRule checks the pattern is valid
type GoogleRedisInstanceInvalidTransitEncryptionModeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleRedisInstanceInvalidTransitEncryptionModeRule returns new rule with default attributes
func NewGoogleRedisInstanceInvalidTransitEncryptionModeRule() *GoogleRedisInstanceInvalidTransitEncryptionModeRule {
	return &GoogleRedisInstanceInvalidTransitEncryptionModeRule{
		resourceType:  "google_redis_instance",
		attributeName: "transit_encryption_mode",
	}
}

// Name returns the rule name
func (r *GoogleRedisInstanceInvalidTransitEncryptionModeRule) Name() string {
	return "google_redis_instance_invalid_transit_encryption_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleRedisInstanceInvalidTransitEncryptionModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleRedisInstanceInvalidTransitEncryptionModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleRedisInstanceInvalidTransitEncryptionModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleRedisInstanceInvalidTransitEncryptionModeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"SERVER_AUTHENTICATION", "DISABLED", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
