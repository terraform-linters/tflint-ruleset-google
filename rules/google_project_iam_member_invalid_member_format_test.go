package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleProjectIamMemberInvalidMemberFormat(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid member",
			Content: `
resource "google_project_iam_member" "iam_member" {
	member = "jane@example.com"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleProjectIamMemberInvalidMemberFormatRule(),
					Message: "jane@example.com is an invalid member format",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 11},
						End:      hcl.Pos{Line: 3, Column: 29},
					},
				},
			},
		},
		{
			Name: "valid member",
			Content: `
resource "google_project_iam_member" "iam_member" {
	member = "user:jane@example.com"
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleProjectIamMemberInvalidMemberFormatRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
