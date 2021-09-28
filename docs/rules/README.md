# Rules

This documentation describes a list of rules available by enabling this ruleset.

## Basic Rules

### Invalid machine types

These rules warn you if a machine type not listed at https://cloud.google.com/compute/docs/machine-types is being used. Please note that custom machine types cannot be detected correctly. These rules consider all machine types starting with `[e2|n2|n2d|n1]-custom-` to be valid.

|Name|Severity|Enabled|
| --- | --- | --- |
|google_composer_environment_invalid_machine_type|ERROR|✔|
|google_compute_instance_invalid_machine_type|ERROR|✔|
|google_compute_instance_template_invalid_machine_type|ERROR|✔|
|google_compute_reservation_invalid_machine_type|ERROR|✔|
|google_container_cluster_invalid_machine_type|ERROR|✔|
|google_container_node_pool_invalid_machine_type|ERROR|✔|
|google_dataflow_job_invalid_machine_type|ERROR|✔|

### Invalid member format

TODO

|Name|Severity|Enabled|
| --- | --- | --- |
|google_project_iam_member_invalid_member_format|ERROR|✔|

## Magic Modules Rules

These are the rules that warn against invalid values generated from [Magic Modules](https://github.com/terraform-linters/magic-modules). These rules are defined under the [`rules/magicmodules`](../../rules/magicmodules) directory.

See the [`tools`](../../tools) directory for how to generate these rules.

## Deep Checking Rules

By enabling [Deep Checking](../deep_checking.md), you can enable rules that invoke API to perform more strict checking.

|Name|Severity|Enabled|
| --- | --- | --- |
|google_disabled_api|ERROR|✔|
