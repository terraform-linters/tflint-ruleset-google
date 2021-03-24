package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleComputeResourcePolicyInvalidName(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "basic",
			Content: `
resource "google_compute_resource_policy" "snapshot_policy" {
  name = "snapshot_policy"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleComputeResourcePolicyInvalidNameRule(),
					Message: `"name" ("snapshot_policy") doesn't match regexp "^[a-z]([-a-z0-9]*[a-z0-9])$"`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 27},
					},
				},
			},
		},
	}

	rule := NewGoogleComputeResourcePolicyInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
