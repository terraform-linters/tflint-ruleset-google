package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_GoogleComposerEnvironmentInvalidMachineType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid type",
			Content: `
resource "google_composer_environment" "test" {
  name   = "mycomposer"
  region = "us-central1"
  config {
    node_count = 4

    node_config {
      zone         = "us-central1-a"
      machine_type = "n1-standard-11"

      network    = google_compute_network.test.id
      subnetwork = google_compute_subnetwork.test.id

      service_account = google_service_account.test.name
    }
  }

  depends_on = [google_project_iam_member.composer-worker]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewGoogleComposerEnvironmentInvalidMachineTypeRule(),
					Message: `"n1-standard-11" is an invalid as machine type`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 10, Column: 22},
						End:      hcl.Pos{Line: 10, Column: 38},
					},
				},
			},
		},
		{
			Name: "valid type",
			Content: `
resource "google_composer_environment" "test" {
  name   = "mycomposer"
  region = "us-central1"
  config {
    node_count = 4

    node_config {
      zone         = "us-central1-a"
      machine_type = "n2-standard-2"

      network    = google_compute_network.test.id
      subnetwork = google_compute_subnetwork.test.id

      service_account = google_service_account.test.name
    }
  }

  depends_on = [google_project_iam_member.composer-worker]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "custom type",
			Content: `
resource "google_composer_environment" "test" {
  name   = "mycomposer"
  region = "us-central1"
  config {
    node_count = 4

    node_config {
      zone         = "us-central1-a"
      machine_type = "custom-6-20480"

      network    = google_compute_network.test.id
      subnetwork = google_compute_subnetwork.test.id

      service_account = google_service_account.test.name
    }
  }

  depends_on = [google_project_iam_member.composer-worker]
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewGoogleComposerEnvironmentInvalidMachineTypeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
