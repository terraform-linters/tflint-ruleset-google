package api

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

// Rules is a list of rules with invoking APIs
var Rules = []tflint.Rule{
	NewGoogleDisabledAPIRule(),
}
