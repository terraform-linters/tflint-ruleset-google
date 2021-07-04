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

