package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/google"
	"github.com/terraform-linters/tflint-ruleset-google/project"
	"github.com/terraform-linters/tflint-ruleset-google/rules"
	"github.com/terraform-linters/tflint-ruleset-google/rules/api"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &google.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "google",
				Version: project.Version,
				Rules:   rules.Rules,
			},
			APIRules: api.Rules,
		},
	})
}
