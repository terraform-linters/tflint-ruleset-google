package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleComputeReservationInvalidMachineType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid type",
			Content: `
resource "google_compute_reservation" "gce_reservation" {
  name = "gce-reservation"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      min_cpu_platform = "Intel Cascade Lake"
      machine_type     = "n2_standard_2"
    }
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleComputeReservationInvalidMachineTypeRule(),
					Message: `"n2_standard_2" is an invalid as machine type`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 10, Column: 26},
						End:      hcl.Pos{Line: 10, Column: 41},
					},
				},
			},
		},
		{
			Name: "valid type",
			Content: `
resource "google_compute_reservation" "gce_reservation" {
  name = "gce-reservation"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      min_cpu_platform = "Intel Cascade Lake"
      machine_type     = "n2-standard-2"
    }
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "custom type",
			Content: `
resource "google_compute_reservation" "gce_reservation" {
  name = "gce-reservation"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      min_cpu_platform = "Intel Cascade Lake"
      machine_type     = "n2-custom-6-20480"
    }
  }
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleComputeReservationInvalidMachineTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
