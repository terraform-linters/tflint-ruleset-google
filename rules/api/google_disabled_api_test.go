package api

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"google.golang.org/api/serviceusage/v1"
)

func Test_Check(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Response map[string]*serviceusage.GoogleApiServiceusageV1Service
		Expected helper.Issues
	}{
		{
			Name: "No resources",
			Content: `
terraform {}`,
			Response: map[string]*serviceusage.GoogleApiServiceusageV1Service{},
			Expected: helper.Issues{},
		},
		{
			Name: "compute.googleapis.com is not enabled",
			Content: `
resource "google_compute_network" "vpc_network" {
	name                    = "terraform-network"
	auto_create_subnetworks = "true"
}`,
			Response: map[string]*serviceusage.GoogleApiServiceusageV1Service{},
			Expected: helper.Issues{
				{
					Rule:    NewGoogleDisabledAPIRule(),
					Message: "compute.googleapis.com has not been used in foo-bar-baz before or it is disabled.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 48},
					},
				},
			},
		},
		{
			Name: "compute.googleapis.com is enabled",
			Content: `
resource "google_compute_network" "vpc_network" {
	name                    = "terraform-network"
	auto_create_subnetworks = "true"
}`,
			Response: map[string]*serviceusage.GoogleApiServiceusageV1Service{
				"compute.googleapis.com": {
					Config: &serviceusage.GoogleApiServiceusageV1ServiceConfig{
						Title: "Compute Engine API",
						Name:  "compute.googleapis.com",
					},
					Name:   "projects/123456789/services/compute.googleapis.com",
					Parent: "projects/123456789",
					State:  "ENABLED",
				},
			},
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleDisabledAPIRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := NewTestRunner(t, map[string]string{"resource.tf": tc.Content})

			rule.prepared = true
			rule.enabledAPIs = tc.Response

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}
			helper.AssertIssues(t, tc.Expected, runner.Runner.(*helper.Runner).Issues)
		})
	}
}
