# Deep Checking

Deep Checking uses your provider's credentials to perform a more strict inspection.

For example, you can confirm that the API you are trying to use in the project is enabled. See also the [google_disabled_api](rules/google_disabled_api.md) rule.

```console
$ tflint
1 issue(s) found:

Error: Compute Engine API has not been used in foo-bar-baz before or it is disabled. (google_disabled_api)

  on template.tf line 25:
  25: resource "google_compute_network" "vpc_network" {

```

You can enable Deep Checking by changing the plugin configuration.

```hcl
plugin "google" {
  enabled = true

  deep_check = true
}
```

## Credentials

Currently, credentials, regions, etc. declared inside the "google" provider block are not considered except for the `project` attribute. You need to pass the credentials to TFLint using environment variables and so on.
