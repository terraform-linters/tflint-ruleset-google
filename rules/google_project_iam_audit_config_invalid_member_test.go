package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleProjectIamAuditConfigInvalidMember(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid member",
			Content: `
resource "google_project_iam_audit_config" "iam_audit_config" {
	audit_log_config {
		exempted_members = [
			"jane@example.com",
		]
	}
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleProjectIamAuditConfigInvalidMemberRule(),
					Message: "jane@example.com is an invalid member format",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 22},
						End:      hcl.Pos{Line: 6, Column: 4},
					},
				},
			},
		},
		{
			Name: "valid member",
			Content: `
resource "google_project_iam_audit_config" "iam_audit_config" {
	audit_log_config {
		exempted_members = [
			"user:jane@example.com",
		]
	}
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "no member",
			Content: `
resource "google_project_iam_audit_config" "iam_audit_config" {
	audit_log_config {
	}
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleProjectIamAuditConfigInvalidMemberRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
