package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleContainerClusterInvalidMachineType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid type",
			Content: `
resource "google_container_cluster" "primary_preemptible_nodes" {
  name               = "my-node-pool"
  location           = "us-central1"
  initial_node_count = 1

  node_config {
    preemptible  = true
    machine_type = "n2-medium"

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleContainerClusterInvalidMachineTypeRule(),
					Message: `"n2-medium" is an invalid as machine type`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 9, Column: 20},
						End:      hcl.Pos{Line: 9, Column: 31},
					},
				},
			},
		},
		{
			Name: "valid type",
			Content: `
resource "google_container_cluster" "primary_preemptible_nodes" {
  name               = "my-node-pool"
  location           = "us-central1"
  initial_node_count = 1

  node_config {
    preemptible  = true
    machine_type = "e2-medium"

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "custom type",
			Content: `
resource "google_container_cluster" "primary_preemptible_nodes" {
  name               = "my-node-pool"
  location           = "us-central1"
  initial_node_count = 1

  node_config {
    preemptible  = true
    machine_type = "custom-6-20480"

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleContainerClusterInvalidMachineTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
