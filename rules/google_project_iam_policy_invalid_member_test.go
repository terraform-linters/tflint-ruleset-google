package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleProjectIamPolicyInvalidMember(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid member",
			Content: `
resource "google_project_iam_policy" "iam_policy" {
	members = [
		"jane@example.com",
	]
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleProjectIamPolicyInvalidMemberRule(),
					Message: "jane@example.com is an invalid member format",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 5, Column: 3},
					},
				},
			},
		},
		{
			Name: "valid member",
			Content: `
resource "google_project_iam_policy" "iam_policy" {
	members = [
		"user:jane@example.com",
	]
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleProjectIamPolicyInvalidMemberRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
