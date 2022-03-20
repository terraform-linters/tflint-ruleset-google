package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/rules/api"
	"github.com/terraform-linters/tflint-ruleset-google/rules/magicmodules"
)

var rules = [][]tflint.Rule{
	{
		NewGoogleComposerEnvironmentInvalidMachineTypeRule(),
		NewGoogleComputeInstanceInvalidMachineTypeRule(),
		NewGoogleComputeReservationInvalidMachineTypeRule(),
		NewGoogleComputeInstanceTemplateInvalidMachineTypeRule(),
		NewGoogleContainerClusterInvalidMachineTypeRule(),
		NewGoogleContainerNodePoolInvalidMachineTypeRule(),
		NewGoogleDataflowJobInvalidMachineTypeRule(),
		NewGoogleComputeResourcePolicyInvalidNameRule(),
		NewGoogleProjectIamMemberInvalidMemberRule(),
		NewGoogleProjectIamAuditConfigInvalidMemberRule(),
		NewGoogleProjectIamBindingInvalidMemberRule(),
		NewGoogleProjectIamPolicyInvalidMemberRule(),
	},
	magicmodules.Rules,
	api.Rules,
}

// Rules is a list of all rules
var Rules []tflint.Rule

func init() {
	for _, r := range rules {
		Rules = append(Rules, r...)
	}
}
