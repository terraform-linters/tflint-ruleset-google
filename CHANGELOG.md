## 0.18.0 (2022-06-19)

### Enhancements

- [#185](https://github.com/terraform-linters/tflint-ruleset-google/pull/185): Update Magic Modules

### Chores

- [#177](https://github.com/terraform-linters/tflint-ruleset-google/pull/177): Bump goreleaser/goreleaser-action from 2 to 3
- [#180](https://github.com/terraform-linters/tflint-ruleset-google/pull/180): Bump github.com/dave/dst from 0.26.2 to 0.27.0
- [#183](https://github.com/terraform-linters/tflint-ruleset-google/pull/183): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.15.0 to 2.17.0
- [#186](https://github.com/terraform-linters/tflint-ruleset-google/pull/186): Bump google.golang.org/api from 0.78.0 to 0.84.0

## 0.17.0 (2022-05-05)

### Enhancements

- [#174](https://github.com/terraform-linters/tflint-ruleset-google/pull/174): Update Magic Modules

### Chores

- [#160](https://github.com/terraform-linters/tflint-ruleset-google/pull/160): chores: Remove snaker
- [#161](https://github.com/terraform-linters/tflint-ruleset-google/pull/161): Fix rule template for rule generator
- [#165](https://github.com/terraform-linters/tflint-ruleset-google/pull/165): Bump actions/setup-go from 2 to 3
- [#169](https://github.com/terraform-linters/tflint-ruleset-google/pull/169) [#173](https://github.com/terraform-linters/tflint-ruleset-google/pull/173): Bump google.golang.org/api from 0.73.0 to 0.78.0
- [#170](https://github.com/terraform-linters/tflint-ruleset-google/pull/170): Bump github.com/google/go-cmp from 0.5.7 to 0.5.8
- [#171](https://github.com/terraform-linters/tflint-ruleset-google/pull/171): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.10.0 to 0.11.0
- [#172](https://github.com/terraform-linters/tflint-ruleset-google/pull/172): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.12.0 to 2.15.0

## 0.16.1 (2022-03-31)

### BugFixes

- [#159](https://github.com/terraform-linters/tflint-ruleset-google/pull/159): Suppress unevaluable/unknown/null errors on provider block eval

## 0.16.0 (2022-03-27)

The minimum supported version of TFLint has changed in this version. TFLint v0.35.0+ is required for this plugin to work.

### Breaking Changes

- [#155](https://github.com/terraform-linters/tflint-ruleset-google/pull/155): Bump tflint-plugin-sdk for gRPC-based new plugin system

### Enhancements

- [#158](https://github.com/terraform-linters/tflint-ruleset-google/pull/158): Update Magic Modules

### Chores

- [#151](https://github.com/terraform-linters/tflint-ruleset-google/pull/151): Bump actions/checkout from 2 to 3
- [#154](https://github.com/terraform-linters/tflint-ruleset-google/pull/154): go 1.18
- [#156](https://github.com/terraform-linters/tflint-ruleset-google/pull/156): Bump google.golang.org/api from 0.60.0 to 0.73.0
- [#157](https://github.com/terraform-linters/tflint-ruleset-google/pull/157): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.9.0 to 2.12.0

## 0.15.0 (2021-12-07)

### Breaking Changes

- [#136](https://github.com/terraform-linters/tflint-ruleset-google/pull/136): Update Magic Modules
  - Removed `google_apigee_instance_invalid_peering_cidr_range` rule
  - Removed `google_filestore_instance_invalid_tier` rule

### Chores

- [#135](https://github.com/terraform-linters/tflint-ruleset-google/pull/135): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.8.0 to 2.9.0

## 0.14.0 (2021-11-07)

### Breaking Changes

- [#133](https://github.com/terraform-linters/tflint-ruleset-google/pull/133): build: Remove unsupported build targets

### Chores

- [#131](https://github.com/terraform-linters/tflint-ruleset-google/pull/131): Bump google.golang.org/api from 0.58.0 to 0.60.0
- [#132](https://github.com/terraform-linters/tflint-ruleset-google/pull/132): Update Magic Modules

## 0.13.2 (2021-11-02)

### BugFixes

- [#130](https://github.com/terraform-linters/tflint-ruleset-google/pull/130): rules: Use WalkResources to avoid a bug in JSON syntax

## 0.13.1 (2021-10-24)

### Enhancements

- [#125](https://github.com/terraform-linters/tflint-ruleset-google/pull/125): rules: Add a2 variants and information about machine type documentation
- [#126](https://github.com/terraform-linters/tflint-ruleset-google/pull/126): rules: Add larger N2 machine types (Ice Lake) and Tau T2D types

## 0.13.0 (2021-10-09)

### Breaking Changes

- [#124](https://github.com/terraform-linters/tflint-ruleset-google/pull/124): Update Magic Modules
  - Removed `google_dns_record_set_invalid_type` rule

### Enhancements

- [#121](https://github.com/terraform-linters/tflint-ruleset-google/pull/121): rules: Add `google_project_iam_member_invalid_member` rule
- [#122](https://github.com/terraform-linters/tflint-ruleset-google/pull/122): rules: Add more `google_project_iam_*_invalid_member` rules

### Chores

- [#120](https://github.com/terraform-linters/tflint-ruleset-google/pull/120): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.7.1 to 2.8.0
- [#123](https://github.com/terraform-linters/tflint-ruleset-google/pull/123): Bump google.golang.org/api from 0.56.0 to 0.58.0

## 0.12.1 (2021-09-12)

### BugFixes

- [#118](https://github.com/terraform-linters/tflint-ruleset-google/pull/118): build: Update GoReleaser version
  - v0.12.0 release doesn't include darwin/arm64 build. This change fixes the issue.

## 0.12.0 (2021-09-12)

### Enhancements

- [#117](https://github.com/terraform-linters/tflint-ruleset-google/pull/117): Update Magic Modules

### Chores

- [#113](https://github.com/terraform-linters/tflint-ruleset-google/pull/113): Bump actions/setup-go from 2.1.3 to 2.1.4
- [#114](https://github.com/terraform-linters/tflint-ruleset-google/pull/114): Bump google.golang.org/api from 0.52.0 to 0.56.0
- [#115](https://github.com/terraform-linters/tflint-ruleset-google/pull/115): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.7.0 to 2.7.1
- [#116](https://github.com/terraform-linters/tflint-ruleset-google/pull/116): build: Go 1.17

## 0.11.0 (2021-08-08)

### Breaking Changes

- [#111](https://github.com/terraform-linters/tflint-ruleset-google/pull/111): Update Magic Modules
  - Remove `google_privateca_certificate_authority_invalid_tier` rule

### Chores

- [#107](https://github.com/terraform-linters/tflint-ruleset-google/pull/107): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.9.0 to 0.9.1
- [#108](https://github.com/terraform-linters/tflint-ruleset-google/pull/108): Bump github.com/hashicorp/hcl/v2 from 2.10.0 to 2.10.1
- [#110](https://github.com/terraform-linters/tflint-ruleset-google/pull/110): Bump google.golang.org/api from 0.49.0 to 0.52.0

## 0.10.0 (2021-07-05)

The minimum supported version of TFLint has changed in this version. TFLint v0.30.0+ is required for this plugin to work.

### Breaking Changes

- [#104](https://github.com/terraform-linters/tflint-ruleset-google/pull/104): Bump tflint-plugin-sdk to v0.9.0

### Enhancements

- [#105](https://github.com/terraform-linters/tflint-ruleset-google/pull/105): Update Magic Modules

### Chores

- [#102](https://github.com/terraform-linters/tflint-ruleset-google/pull/102): Bump google.golang.org/api from 0.47.0 to 0.49.0
- [#103](https://github.com/terraform-linters/tflint-ruleset-google/pull/103): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.6.1 to 2.7.0

## 0.9.1 (2021-06-12)

### Chores

- [#100](https://github.com/terraform-linters/tflint-ruleset-google/pull/100): build: Add support for darwin/arm64 build

## 0.9.0 (2021-06-06)

### Enhancements

- [#98](https://github.com/terraform-linters/tflint-ruleset-google/pull/98): Update Magic Modules

### Chores

- [#96](https://github.com/terraform-linters/tflint-ruleset-google/pull/96): Bump google.golang.org/api from 0.45.0 to 0.47.0
- [#97](https://github.com/terraform-linters/tflint-ruleset-google/pull/97): Bump github.com/google/go-cmp from 0.5.5 to 0.5.6
- [#99](https://github.com/terraform-linters/tflint-ruleset-google/pull/99): Add notes about auto installation

## 0.8.0 (2021-05-02)

### Breaking Changes

- [#94](https://github.com/terraform-linters/tflint-ruleset-google/pull/94): Update Magic Modules
  - `google_compute_address_invalid_purpose` rule and `google_compute_global_address_invalid_purpose` rule are removed

### Chores

- [#91](https://github.com/terraform-linters/tflint-ruleset-google/pull/91): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.5.0 to 2.6.1
- [#92](https://github.com/terraform-linters/tflint-ruleset-google/pull/92): Bump google.golang.org/api from 0.43.0 to 0.45.0
- [#93](https://github.com/terraform-linters/tflint-ruleset-google/pull/93): Bump github.com/hashicorp/hcl/v2 from 2.9.1 to 2.10.0

## 0.7.0 (2021-04-04)

This release adds support for deep checking. `google_disabled_api` rule finds resource declarations that depend on APIs that are not enabled. See [Deep Checking](docs/deep_checking.md) for details.

### Enhancements

- [#75](https://github.com/terraform-linters/tflint-ruleset-google/pull/75): rule: Add google_disabled_api rule
- [#84](https://github.com/terraform-linters/tflint-ruleset-google/pull/84): rule: Add google_compute_resource_policy_invalid_name rule
- [#88](https://github.com/terraform-linters/tflint-ruleset-google/pull/88): Update Magic Modules

### Chores

- [#80](https://github.com/terraform-linters/tflint-ruleset-google/pull/80): Bump github.com/google/go-cmp from 0.5.4 to 0.5.5
- [#81](https://github.com/terraform-linters/tflint-ruleset-google/pull/81): Bump github.com/hashicorp/hcl/v2 from 2.9.0 to 2.9.1
- [#83](https://github.com/terraform-linters/tflint-ruleset-google/pull/83): Add rule generator
- [#86](https://github.com/terraform-linters/tflint-ruleset-google/pull/86): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.4.4 to 2.5.0
- [#87](https://github.com/terraform-linters/tflint-ruleset-google/pull/87): Bump google.golang.org/api from 0.40.0 to 0.43.0
- [#89](https://github.com/terraform-linters/tflint-ruleset-google/pull/89): Bump tflint-plugin-sdk

## 0.6.0 (2021-03-07)

### Enhancements

- [#79](https://github.com/terraform-linters/tflint-ruleset-google/pull/79): Update Magic Modules

### Chores

- [#76](https://github.com/terraform-linters/tflint-ruleset-google/pull/76): Upgrade to Go 1.16
- [#77](https://github.com/terraform-linters/tflint-ruleset-google/pull/77): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.4.2 to 2.4.4
- [#78](https://github.com/terraform-linters/tflint-ruleset-google/pull/78): Bump github.com/hashicorp/hcl/v2 from 2.8.2 to 2.9.0

## 0.5.1 (2021-02-02)

### BugFixes

- [#73](https://github.com/terraform-linters/tflint-ruleset-google/pull/73): Bump tflint-plugin-sdk to v0.8.1

### Chores

- [#72](https://github.com/terraform-linters/tflint-ruleset-google/pull/72): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.4.1 to 2.4.2

## 0.5.0 (2021-02-01)

The minimum supported version of TFLint has changed in this version. TFLint v0.24.0+ is required for this plugin to work.

### Breaking Changes

- [#70](https://github.com/terraform-linters/tflint-ruleset-google/pull/70): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.7.0 to 0.8.0

### Enhancements

- [#71](https://github.com/terraform-linters/tflint-ruleset-google/pull/71): Update Magic Modules

### Chores

- [#69](https://github.com/terraform-linters/tflint-ruleset-google/pull/69): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.4.0 to 2.4.1

## 0.4.0 (2021-01-04)

The minimum supported version of TFLint has changed in this version. TFLint v0.23.0+ is required for this plugin to work.

### Breaking Changes

- [#65](https://github.com/terraform-linters/tflint-ruleset-google/pull/65): Bump tflint-plugin-sdk to v0.7.0

### Enhancements

- [#64](https://github.com/terraform-linters/tflint-ruleset-google/pull/64): Update Magic Modules

### Chores

- [#59](https://github.com/terraform-linters/tflint-ruleset-google/pull/59): Bump github.com/google/go-cmp from 0.5.3 to 0.5.4
- [#62](https://github.com/terraform-linters/tflint-ruleset-google/pull/62): Bump github.com/hashicorp/hcl/v2 from 2.7.1 to 2.8.1
- [#63](https://github.com/terraform-linters/tflint-ruleset-google/pull/63): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.3.0 to 2.4.0

## 0.3.0 (2020-11-24)

The minimum supported version of TFLint has changed in this version. TFLint v0.21.0+ is required for this plugin to work.

### Breaking Changes

- [#57](https://github.com/terraform-linters/tflint-ruleset-google/pull/57): Bump tflint-plugin-sdk to v0.6.0
  - Added support for JSON configuration syntax

### Enhancements

- [#56](https://github.com/terraform-linters/tflint-ruleset-google/pull/56): Update Magic Modules

### Chores

- [#52](https://github.com/terraform-linters/tflint-ruleset-google/pull/52): Bump github.com/google/go-cmp from 0.5.2 to 0.5.3
- [#53](https://github.com/terraform-linters/tflint-ruleset-google/pull/53): Stop using set-env commands
- [#54](https://github.com/terraform-linters/tflint-ruleset-google/pull/54): Bump github.com/hashicorp/hcl/v2 from 2.6.0 to 2.7.1
- [#55](https://github.com/terraform-linters/tflint-ruleset-google/pull/55): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.0.4 to 2.3.0

## 0.2.0 (2020-10-19)

### Enhancements

- [#43](https://github.com/terraform-linters/tflint-ruleset-google/pull/43) [#45](https://github.com/terraform-linters/tflint-ruleset-google/pull/45) [#48](https://github.com/terraform-linters/tflint-ruleset-google/pull/48): Add invalid machine type rules
- [#47](https://github.com/terraform-linters/tflint-ruleset-google/pull/47): Update Magic Modules

### Chores

- [#44](https://github.com/terraform-linters/tflint-ruleset-google/pull/44): Bump actions/setup-go from v2.1.2 to v2.1.3
- [#46](https://github.com/terraform-linters/tflint-ruleset-google/pull/46): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.0.3 to 2.0.4

## 0.1.0 (2020-09-22)

Initial release 🎉

