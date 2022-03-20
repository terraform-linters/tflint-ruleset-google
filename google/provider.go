package google

import (
	"os"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleProviderBlockSchema is a schema of `google` provider block
var GoogleProviderBlockSchema = &hclext.BodySchema{
	Attributes: []hclext.AttributeSchema{
		{Name: "project"},
	},
}

// GetProject retrieves project_id from the "provider" block in the Terraform configuration and environment variables
func GetProject(runner tflint.Runner) (string, error) {
	providers, err := runner.GetModuleContent(
		&hclext.BodySchema{
			Blocks: []hclext.BlockSchema{
				{
					Type:       "provider",
					LabelNames: []string{"name"},
					Body:       GoogleProviderBlockSchema,
				},
			},
		},
		&tflint.GetModuleContentOption{ModuleCtx: tflint.RootModuleCtxType},
	)
	if err != nil {
		return "", err
	}

	for _, provider := range providers.Blocks {
		if provider.Labels[0] != "google" {
			continue
		}

		opts := &tflint.EvaluateExprOption{ModuleCtx: tflint.RootModuleCtxType}

		if attr, exists := provider.Body.Attributes["project"]; exists {
			var project string
			if err := runner.EvaluateExpr(attr.Expr, &project, opts); err != nil {
				return project, err
			}
		}
	}

	return getProjectFromEnv(), nil
}

func getProjectFromEnv() string {
	envs := []string{"GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT"}
	for _, env := range envs {
		if project := os.Getenv(env); project != "" {
			return project
		}
	}
	return ""
}
