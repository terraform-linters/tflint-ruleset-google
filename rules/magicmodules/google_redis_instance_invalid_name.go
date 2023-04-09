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

// GoogleRedisInstanceInvalidNameRule checks the pattern is valid
type GoogleRedisInstanceInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleRedisInstanceInvalidNameRule returns new rule with default attributes
func NewGoogleRedisInstanceInvalidNameRule() *GoogleRedisInstanceInvalidNameRule {
	return &GoogleRedisInstanceInvalidNameRule{
		resourceType:  "google_redis_instance",
		attributeName: "name",
	}
}

// Name returns the rule name
func (r *GoogleRedisInstanceInvalidNameRule) Name() string {
	return "google_redis_instance_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleRedisInstanceInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleRedisInstanceInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleRedisInstanceInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleRedisInstanceInvalidNameRule) Check(runner tflint.Runner) error {
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
			validateFunc := validateRegexp(`^[a-z][a-z0-9-]{0,39}[a-z0-9]$`)

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
