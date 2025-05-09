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

// GoogleNetworkServicesGrpcRouteInvalidLocationRule checks the pattern is valid
type GoogleNetworkServicesGrpcRouteInvalidLocationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleNetworkServicesGrpcRouteInvalidLocationRule returns new rule with default attributes
func NewGoogleNetworkServicesGrpcRouteInvalidLocationRule() *GoogleNetworkServicesGrpcRouteInvalidLocationRule {
	return &GoogleNetworkServicesGrpcRouteInvalidLocationRule{
		resourceType:  "google_network_services_grpc_route",
		attributeName: "location",
	}
}

// Name returns the rule name
func (r *GoogleNetworkServicesGrpcRouteInvalidLocationRule) Name() string {
	return "google_network_services_grpc_route_invalid_location"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleNetworkServicesGrpcRouteInvalidLocationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleNetworkServicesGrpcRouteInvalidLocationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleNetworkServicesGrpcRouteInvalidLocationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleNetworkServicesGrpcRouteInvalidLocationRule) Check(runner tflint.Runner) error {
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
			validateFunc := validateRegexp(`^global$`)

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
