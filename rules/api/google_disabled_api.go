package api

import (
	"context"
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/google"
	"github.com/terraform-linters/tflint-ruleset-google/rules/magicmodules"
	"google.golang.org/api/serviceusage/v1"
)

// GoogleDisabledAPIRule checks whether the API required by resources is disabled
type GoogleDisabledAPIRule struct {
	prepared    bool
	enabledAPIs map[string]*serviceusage.GoogleApiServiceusageV1Service
}

// NewGoogleDisabledAPIRule returns a new rule
func NewGoogleDisabledAPIRule() *GoogleDisabledAPIRule {
	return &GoogleDisabledAPIRule{
		prepared:    false,
		enabledAPIs: map[string]*serviceusage.GoogleApiServiceusageV1Service{},
	}
}

// Name returns the rule name
func (r *GoogleDisabledAPIRule) Name() string {
	return "google_disabled_api_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDisabledAPIRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDisabledAPIRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDisabledAPIRule) Link() string {
	return ""
}

// Check checks whether the API required by resources is disabled
func (r *GoogleDisabledAPIRule) Check(rr tflint.Runner) error {
	runner := rr.(*google.Runner)

	if runner.Project == "" {
		return nil
	}

	if !r.prepared {
		ctx := context.TODO()
		err := runner.Client.ServiceUsage.Services.List(runner.ParentProject()).Filter("state:ENABLED").Pages(ctx, func(resp *serviceusage.ListServicesResponse) error {
			for _, service := range resp.Services {
				r.enabledAPIs[service.Config.Title] = service
			}
			return nil
		})
		if err != nil {
			err := &tflint.Error{
				Code:    tflint.ExternalAPIError,
				Level:   tflint.ErrorLevel,
				Message: "An error occurred while invoking ServiceUsage.List",
				Cause:   err,
			}
			return err
		}
	}

	for resource, product := range magicmodules.Products {
		if len(product.APIsRequired) == 0 {
			continue
		}

		err := runner.WalkResources(resource, func(resource *configs.Resource) error {
			for _, ref := range product.APIsRequired {
				if _, ok := r.enabledAPIs[ref.Name]; !ok {
					err := runner.EmitIssue(
						r,
						fmt.Sprintf("%s has not been used in %s before or it is disabled.", ref.Name, runner.Project),
						resource.DeclRange,
					)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
