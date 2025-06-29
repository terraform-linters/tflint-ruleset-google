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

// GoogleStorageBatchOperationsJobInvalidJobIdRule checks the pattern is valid
type GoogleStorageBatchOperationsJobInvalidJobIdRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleStorageBatchOperationsJobInvalidJobIdRule returns new rule with default attributes
func NewGoogleStorageBatchOperationsJobInvalidJobIdRule() *GoogleStorageBatchOperationsJobInvalidJobIdRule {
	return &GoogleStorageBatchOperationsJobInvalidJobIdRule{
		resourceType:  "google_storage_batch_operations_job",
		attributeName: "job_id",
	}
}

// Name returns the rule name
func (r *GoogleStorageBatchOperationsJobInvalidJobIdRule) Name() string {
	return "google_storage_batch_operations_job_invalid_job_id"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleStorageBatchOperationsJobInvalidJobIdRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleStorageBatchOperationsJobInvalidJobIdRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleStorageBatchOperationsJobInvalidJobIdRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleStorageBatchOperationsJobInvalidJobIdRule) Check(runner tflint.Runner) error {
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
			validateFunc := validateRegexp(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?.$`)

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
