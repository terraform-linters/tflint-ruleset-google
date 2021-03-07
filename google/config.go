package google

import "github.com/hashicorp/hcl/v2"

// Config is the configuration for the ruleset.
type Config struct {
	DeepCheck bool `hcl:"deep_check,optional"`

	Remain hcl.Body `hcl:",remain"`
}
