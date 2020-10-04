package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/rules/magicmodules"
)

// Rules is a list of all rules
var Rules = append([]tflint.Rule{
	NewGoogleComposerEnvironmentInvalidMachineTypeRule(),
	NewGoogleComputeInstanceInvalidMachineTypeRule(),
	NewGoogleComputeReservationInvalidMachineTypeRule(),
	NewGoogleComputeInstanceTemplateInvalidMachineTypeRule(),
	NewGoogleContainerClusterInvalidMachineTypeRule(),
	NewGoogleContainerNodePoolInvalidMachineTypeRule(),
	NewGoogleDataflowJobInvalidMachineTypeRule(),
}, magicmodules.Rules...)
