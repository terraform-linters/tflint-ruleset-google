# google_project_iam_member_invalid_member_format

Check iam member format.

## Example

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

## How To Fix

See the document. Probably you just need to add some prefixes.
