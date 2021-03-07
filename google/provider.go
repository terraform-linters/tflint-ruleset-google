package google

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleProviderBlockSchema is a schema of `google` provider block
var GoogleProviderBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "project"},
	},
}

// ProviderData represents a provider block with an eval context (runner)
type ProviderData struct {
	provider   *configs.Provider
	runner     tflint.Runner
	attributes hcl.Attributes
	blocks     hcl.Blocks
}

// GetProject retrieves project_id from the "provider" block in the Terraform configuration and environment variables
func GetProject(runner tflint.Runner) (string, error) {
	provider, err := runner.RootProvider("google")
	if err != nil {
		return "", err
	}
	if provider == nil {
		return getProjectFromEnv(), nil
	}

	d, err := newProviderData(provider, runner)
	if err != nil {
		return "", err
	}

	project, exists, err := d.Get("project")
	if exists {
		return project, err
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

func newProviderData(provider *configs.Provider, runner tflint.Runner) (*ProviderData, error) {
	providerData := &ProviderData{
		provider:   provider,
		runner:     runner,
		attributes: map[string]*hcl.Attribute{},
		blocks:     []*hcl.Block{},
	}

	if provider != nil {
		content, _, diags := provider.Config.PartialContent(GoogleProviderBlockSchema)
		if diags.HasErrors() {
			return nil, diags
		}

		providerData.attributes = content.Attributes
		providerData.blocks = content.Blocks
	}

	return providerData, nil
}

// Get returns a value corresponding to the given key
// It should be noted that the value is evaluated if it is evaluable
// The second return value is a flag that determines whether a value exists
// We assume the provider has only simple attributes, so it just returns string
func (d *ProviderData) Get(key string) (string, bool, error) {
	attribute, exists := d.attributes[key]
	if !exists {
		return "", false, nil
	}

	var val string
	err := d.runner.EvaluateExprOnRootCtx(attribute.Expr, &val, nil)

	err = d.runner.EnsureNoError(err, func() error { return nil })
	if err != nil {
		return "", true, err
	}
	return val, true, nil
}
