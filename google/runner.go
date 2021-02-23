package google

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Runner is a wrapper of RPC client for inserting custom actions for Google provider.
type Runner struct {
	tflint.Runner
	Client  *Client
	Project string
}

// NewRunner returns a custom Google runner.
func NewRunner(runner tflint.Runner, config *Config) (*Runner, error) {
	var client *Client
	var project string
	var err error
	if config.DeepCheck {
		client, err = NewClient()
		if err != nil {
			return nil, err
		}

		project, err = GetProject(runner)
		if err != nil {
			return nil, err
		}
	}

	return &Runner{
		Runner:  runner,
		Client:  client,
		Project: project,
	}, nil
}

// ParentProject returns a part of the API path about the parent project
func (r *Runner) ParentProject() string {
	return fmt.Sprintf("projects/%s", r.Project)
}
