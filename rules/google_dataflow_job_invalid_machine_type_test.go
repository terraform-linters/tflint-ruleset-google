package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleDataflowJobInvalidMachineType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid type",
			Content: `
resource "google_dataflow_job" "vm_instance" {
    machine_type = "n2-standard-1"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleDataflowJobInvalidMachineTypeRule(),
					Message: `"n2-standard-1" is an invalid as machine type`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 20},
						End:      hcl.Pos{Line: 3, Column: 35},
					},
				},
			},
		},
		{
			Name: "valid type",
			Content: `
resource "google_dataflow_job" "vm_instance" {
    machine_type = "n2-standard-2"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "custom type",
			Content: `
resource "google_dataflow_job" "vm_instance" {
    machine_type = "n2-custom-6-20480"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleDataflowJobInvalidMachineTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
