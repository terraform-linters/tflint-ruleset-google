package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GoogleProjectIamAuditConfigInvalidMemberRule checks whether member value is invalid
type GoogleProjectIamAuditConfigInvalidMemberRule struct {
	resourceType  string
	blockName     string
	attributeName string
}

func NewGoogleProjectIamAuditConfigInvalidMemberRule() *GoogleProjectIamAuditConfigInvalidMemberRule {
	return &GoogleProjectIamAuditConfigInvalidMemberRule{
		resourceType:  "google_project_iam_audit_config",
		blockName:     "audit_log_config",
		attributeName: "exempted_members",
	}
}

// Name returns the rule name
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Name() string {
	return "google_project_iam_audit_config_invalid_member"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether member format is invalid
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceBlocks(r.resourceType, r.blockName, func(block *hcl.Block) error {
		content, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{Name: r.attributeName},
			},
		})
		if diags.HasErrors() {
			return diags
		}

		if attribute, exists := content.Attributes[r.attributeName]; exists {
			var members []string
			err := runner.EvaluateExpr(attribute.Expr, &members, nil)

			return runner.EnsureNoError(err, func() error {
				for _, member := range members {
					if !isValidIAMMemberFormat(member) {
						return runner.EmitIssueOnExpr(
							r,
							fmt.Sprintf("%s is an invalid member format", member),
							attribute.Expr,
						)
					}
				}
				return nil
			})
		}

		return nil
	})
}
