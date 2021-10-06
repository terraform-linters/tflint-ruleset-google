# google_project_iam_invalid_member

Check IAM member format against IAM policy for projects:

- `google_project_iam_policy`
- `google_project_iam_binding`
- `google_project_iam_member`
- `google_project_iam_audit_config`

## Example

Here is an example of `google_project_iam_member`.

```hcl
resource "google_project_iam_member" "iam_member" {
	member = "jane@example.com"
}
```

```
$ tflint
1 issue(s) found:

Error: jane@example.com is an invalid member format. (google_project_iam_member)

  on template.tf line 1:
   1: resource "google_project_iam_member" "iam_member" {

```

## Why

A member value you wrote is invalid. `terraform apply` will result in failure.
In particular, don't forget to add a prefix to the value. Since the principal type is identified with a prefix, a value without a prefix cannot be interpreted correctly. There are two exceptions; `allUsers` and `allAuthenticatedUsers` are valid without a prefix. In conclusion, the valid member values are:

- `allUsers`
- `allAuthenticatedUsers`
- `user:{emailid}`
- `serviceAccount:{emailid}`
- `group:{emailid}`
- `domain:{domain}`

See these documents in detail:

 - https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam#member/members
 - https://cloud.google.com/iam/docs/overview#concepts_related_identity
 - https://cloud.google.com/iam/docs/overview#cloud-iam-policy

## How To Fix

See the document. Probably you just need to add some prefixes. The example can be fixed as below:

```hcl
resource "google_project_iam_member" "iam_member" {
	member = "user:jane@example.com"
}
```
