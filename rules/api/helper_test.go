package api

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"github.com/terraform-linters/tflint-ruleset-google/google"
)

func NewTestRunner(t *testing.T, files map[string]string) *google.Runner {
	return &google.Runner{
		Runner:  helper.TestRunner(t, files),
		Client:  &google.Client{},
		Project: "foo-bar-baz",
	}
}
