package magicmodules

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_generatedRegexRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid",
			Content: `
resource "google_compute_backend_bucket" "main" {
  name = "123"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleComputeBackendBucketInvalidNameRule(),
					Message: `"name" ("123") doesn't match regexp "^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$"`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 15},
					},
				},
			},
		},
		{
			Name: "valid",
			Content: `
resource "google_compute_backend_bucket" "main" {
  name = "valid"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleComputeBackendBucketInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func Test_generatedEnumRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid",
			Content: `
resource "google_cloud_asset_folder_feed" "main" {
  content_type = "INVALID"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleCloudAssetFolderFeedInvalidContentTypeRule(),
					Message: `expected content_type to be one of [CONTENT_TYPE_UNSPECIFIED RESOURCE IAM_POLICY ORG_POLICY ACCESS_POLICY ], got INVALID`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 18},
						End:      hcl.Pos{Line: 3, Column: 27},
					},
				},
			},
		},
		{
			Name: "valid",
			Content: `
resource "google_cloud_asset_folder_feed" "main" {
  name = "RESOURCE"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleCloudAssetFolderFeedInvalidContentTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
