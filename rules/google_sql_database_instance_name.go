package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
)

// TODO: Write the rule's description here
// GoogleSQLDatabaseInstanceNameRule checks ...
type GoogleSQLDatabaseInstanceNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleSQLDatabaseInstanceNameRule returns new rule with default attributes
func NewGoogleSQLDatabaseInstanceNameRule() *GoogleSQLDatabaseInstanceNameRule {
	return &GoogleSQLDatabaseInstanceNameRule{
		resourceType:  "google_sql_database_instance",
		attributeName: "name",
	}
}

// Name returns the rule name
func (r *GoogleSQLDatabaseInstanceNameRule) Name() string {
	return "google_sql_database_instance_name"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleSQLDatabaseInstanceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleSQLDatabaseInstanceNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleSQLDatabaseInstanceNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *GoogleSQLDatabaseInstanceNameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			// regex to check if name contains only lowercase letters, numbers, and hyphens
			validateFunc := validateRegexp(`^[a-z0-9-]+$`)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
