package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleSQLDatabaseInstanceName(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "basic",
			Content: `
resource "google_sql_database_instance" "instance" {
  name = "my_instance"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleSQLDatabaseInstanceNameRule(),
					Message: `"name" ("instance") doesn't match regexp "^[a-z0-9-]+$"`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 0, Column: 0},
						End:      hcl.Pos{Line: 0, Column: 0},
					},
				},
			},
		},
	}

	rule := NewGoogleSQLDatabaseInstanceNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
