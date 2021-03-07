# google_disabled_api

Disallow to declare resources with disabled API.

## Example

```hcl
resource "google_compute_network" "vpc_network" {
  name                    = "terraform-network"
  auto_create_subnetworks = "true"
}
```

```
$ tflint
1 issue(s) found:

Error: Compute Engine API has not been used in foo-bar-baz before or it is disabled. (google_disabled_api_rule)

  on template.tf line 1:
   1: resource "google_compute_network" "vpc_network" {

```

[Service Usage API](https://cloud.google.com/service-usage/docs/reference/rest) must be enabled in order to use this rule.

## Why

`terraform apply` will fail when the resources refer to disabled APIs.

```
$ terraform apply

...

Error: Error creating Network: googleapi: Error 403: Compute Engine API has not been used in project 1234567890 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/compute.googleapis.com/overview?project=1234567890 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry., accessNotConfigured
```

Even more unfortunately, not all errors are returned. For example, if you have resources that depend on multiple disabled APIs, you will not see an error for the other APIs. This rule allows you to detect all disabled API errors in advance.

## How to Fix

Enable the API from the Cloud Console in your project or remove the resource.

https://cloud.google.com/endpoints/docs/openapi/enable-api

