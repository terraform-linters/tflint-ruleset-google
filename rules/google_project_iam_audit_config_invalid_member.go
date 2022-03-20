package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// GoogleProjectIamAuditConfigInvalidMemberRule checks whether member value is invalid
type GoogleProjectIamAuditConfigInvalidMemberRule struct {
	tflint.DefaultRule

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
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether member format is invalid
func (r *GoogleProjectIamAuditConfigInvalidMemberRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: r.blockName,
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, config := range resource.Body.Blocks {
			attribute, exists := config.Body.Attributes[r.attributeName]
			if !exists {
				continue
			}

			var members []string
			err := runner.EvaluateExpr(attribute.Expr, &members, nil)

			err = runner.EnsureNoError(err, func() error {
				for _, member := range members {
					if !isValidIAMMemberFormat(member) {
						return runner.EmitIssue(
							r,
							fmt.Sprintf("%s is an invalid member format", member),
							attribute.Expr.Range(),
						)
					}
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
