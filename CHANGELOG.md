## 0.28.0 (2024-05-05)

### Enhancements

- [#350](https://github.com/terraform-linters/tflint-ruleset-google/pull/350): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.18.0 to 0.20.0
  - This is required for provider-defined functions support
- [#352](https://github.com/terraform-linters/tflint-ruleset-google/pull/352): Update Magic Modules

### Chores

- [#337](https://github.com/terraform-linters/tflint-ruleset-google/pull/337): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.32.0 to 2.33.0
- [#349](https://github.com/terraform-linters/tflint-ruleset-google/pull/349): Bump google.golang.org/api from 0.165.0 to 0.177.0
- [#351](https://github.com/terraform-linters/tflint-ruleset-google/pull/351): deps: Go 1.22.2

## 0.27.1 (2024-02-24)

### Enhancements

- [#335](https://github.com/terraform-linters/tflint-ruleset-google/pull/335): feat: add H100 machine-type to ruleset

### Chores

- [#334](https://github.com/terraform-linters/tflint-ruleset-google/pull/334): Bump google.golang.org/api from 0.163.0 to 0.165.0

## 0.27.0 (2024-02-14)

### Breaking Changes

- [#332](https://github.com/terraform-linters/tflint-ruleset-google/pull/332): Update Magic Module
  - The following rules are renamed
    - `google_big_query_routine_invalid_determinism_level` -> `google_bigquery_routine_invalid_determinism_level`
    - `google_big_query_routine_invalid_language` -> `google_bigquery_routine_invalid_language`
    - `google_big_query_routine_invalid_routine_type` -> `google_bigquery_routine_invalid_routine_type`
    - `google_cloud_build_trigger_invalid_include_build_logs` -> `google_cloudbuild_trigger_invalid_include_build_logs`
    - `google_security_center_organization_custom_module_invalid_enablement_state` -> `google_scc_event_threat_detection_custom_module_invalid_enablement_state`
    - `google_security_center_folder_custom_module` -> `google_scc_folder_custom_module_invalid_enablement_state`
    - `google_security_center_project_custom_module_invalid_enablement_state` -> `google_scc_project_custom_module_invalid_enablement_state`
    - `google_security_center_source_invalid_display_name` -> `google_scc_source_invalid_display_name`
  - Removed `google_compute_router_peer_invalid_advertise_mode` rule

### Chores

- [#320](https://github.com/terraform-linters/tflint-ruleset-google/pull/320): Bump actions/setup-go from 4 to 5
- [#328](https://github.com/terraform-linters/tflint-ruleset-google/pull/328): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.29.0 to 2.32.0
- [#331](https://github.com/terraform-linters/tflint-ruleset-google/pull/331): Bump google.golang.org/api from 0.150.0 to 0.163.0
- [#333](https://github.com/terraform-linters/tflint-ruleset-google/pull/333): deps: Go 1.22

## 0.26.0 (2023-11-13)

### Breaking Changes

- [#315](https://github.com/terraform-linters/tflint-ruleset-google/pull/315): Update Magic Modules
  - Removed `google_cloudiot_registry_invalid_log_level` rule
  - Removed `google_cloud_iot_device_invalid_log_level` rule

### Enhancements

- [#313](https://github.com/terraform-linters/tflint-ruleset-google/pull/313): update machine types

### Chores

- [#298](https://github.com/terraform-linters/tflint-ruleset-google/pull/298): Bump goreleaser/goreleaser-action from 4 to 5
- [#304](https://github.com/terraform-linters/tflint-ruleset-google/pull/304): Bump golang.org/x/net from 0.14.0 to 0.17.0
- [#306](https://github.com/terraform-linters/tflint-ruleset-google/pull/306): Bump github.com/google/go-cmp from 0.5.9 to 0.6.0
- [#308](https://github.com/terraform-linters/tflint-ruleset-google/pull/308): Bump github.com/dave/dst from 0.27.2 to 0.27.3
- [#309](https://github.com/terraform-linters/tflint-ruleset-google/pull/309): Bump github.com/hashicorp/hcl/v2 from 2.18.0 to 2.19.1
- [#311](https://github.com/terraform-linters/tflint-ruleset-google/pull/311): Bump google.golang.org/grpc from 1.57.0 to 1.57.1
- [#314](https://github.com/terraform-linters/tflint-ruleset-google/pull/314): Bump google.golang.org/api from 0.139.0 to 0.150.0

## 0.25.0 (2023-09-11)

### Enhancements

- [#296](https://github.com/terraform-linters/tflint-ruleset-google/pull/296): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.26.1 to 2.29.0
- [#297](https://github.com/terraform-linters/tflint-ruleset-google/pull/297): Update Magic Modules

### Chores

- [#286](https://github.com/terraform-linters/tflint-ruleset-google/pull/286): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.17.0 to 0.18.0
- [#288](https://github.com/terraform-linters/tflint-ruleset-google/pull/288): Add raw binary entries to checksums.txt
- [#292](https://github.com/terraform-linters/tflint-ruleset-google/pull/292): Bump actions/checkout from 3 to 4
- [#293](https://github.com/terraform-linters/tflint-ruleset-google/pull/293): Bump github.com/hashicorp/hcl/v2 from 2.17.0 to 2.18.0
- [#294](https://github.com/terraform-linters/tflint-ruleset-google/pull/294): Go 1.21
- [#295](https://github.com/terraform-linters/tflint-ruleset-google/pull/295): Bump google.golang.org/api from 0.128.0 to 0.139.0

## 0.24.0 (2023-06-26)

### Breaking Changes

- [#279](https://github.com/terraform-linters/tflint-ruleset-google/pull/279): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.16.1 to 0.17.0
  - TFLint v0.40/v0.41 is no longer supported
- [#280](https://github.com/terraform-linters/tflint-ruleset-google/pull/280): Update Magic Modules
  - Remove `google_kms_crypto_key_invalid_purpose` rule

### Chores

- [#278](https://github.com/terraform-linters/tflint-ruleset-google/pull/278): Bump google.golang.org/api from 0.114.0 to 0.128.0

## 0.23.1 (2023-04-13)

### BugFixes

- [#267](https://github.com/terraform-linters/tflint-ruleset-google/pull/267): Bump tflint-plugin-sdk to v0.16.1

## 0.23.0 (2023-04-09)

### Breaking Changes

- [#262](https://github.com/terraform-linters/tflint-ruleset-google/pull/262): Update Magic Modules
  - Rename `google_bigquery_routine_invalid_determinism_level` to `google_big_query_routine_invalid_determinism_level`
  - Rename `google_bigquery_routine_invalid_language` to `google_big_query_routine_invalid_language`
  - Rename `google_bigquery_routine_invalid_routine_type` to `google_big_query_routine_invalid_routine_type`
  - Rename `google_cloudbuild_trigger_invalid_include_build_logs` to `google_cloud_build_trigger_invalid_include_build_logs`
  - Rename `google_cloudiot_device_invalid_log_level` to `google_cloud_iot_device_invalid_log_level`
  - Rename `google_scc_source_invalid_display_name` to `google_security_center_source_invalid_display_name`

### Enhancements

- [#260](https://github.com/terraform-linters/tflint-ruleset-google/pull/260): Add support for g2 instance types
- [#261](https://github.com/terraform-linters/tflint-ruleset-google/pull/261): Add a2-ultra and m2-hypermem machine types

### Chores

- [#239](https://github.com/terraform-linters/tflint-ruleset-google/pull/239): Use NewRunner hook
- [#247](https://github.com/terraform-linters/tflint-ruleset-google/pull/247): Bump golang.org/x/net from 0.3.0 to 0.7.0
- [#250](https://github.com/terraform-linters/tflint-ruleset-google/pull/250): Bump sigstore/cosign-installer from 2 to 3
- [#254](https://github.com/terraform-linters/tflint-ruleset-google/pull/254): Bump actions/setup-go from 3 to 4
- [#256](https://github.com/terraform-linters/tflint-ruleset-google/pull/256): Bump google.golang.org/api from 0.109.0 to 0.114.0
- [#257](https://github.com/terraform-linters/tflint-ruleset-google/pull/257): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.24.1 to 2.26.1
- [#258](https://github.com/terraform-linters/tflint-ruleset-google/pull/258): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.15.0 to 0.16.0
- [#263](https://github.com/terraform-linters/tflint-ruleset-google/pull/263): Follow up of the EnsureNoError deprecation
- [#264](https://github.com/terraform-linters/tflint-ruleset-google/pull/264): deps: Go 1.20

## 0.22.2 (2023-02-05)

### Enhancements

- [#244](https://github.com/terraform-linters/tflint-ruleset-google/pull/244): Update Magic Modules

### Chores

- [#237](https://github.com/terraform-linters/tflint-ruleset-google/pull/237): Pass `GITHUB_TOKEN` to e2e test workflow
- [#235](https://github.com/terraform-linters/tflint-ruleset-google/pull/235): Bump goreleaser/goreleaser-action from 3 to 4
- [#238](https://github.com/terraform-linters/tflint-ruleset-google/pull/238): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.14.0 to 0.15.0
- [#242](https://github.com/terraform-linters/tflint-ruleset-google/pull/242) [#245](https://github.com/terraform-linters/tflint-ruleset-google/pull/245): Bump google.golang.org/api from 0.103.0 to 0.109.0
- [#243](https://github.com/terraform-linters/tflint-ruleset-google/pull/243): Bump github.com/hashicorp/hcl/v2 from 2.15.0 to 2.16.0

## 0.22.1 (2022-12-05)

### Enhancements

- [#232](https://github.com/terraform-linters/tflint-ruleset-google/pull/232): Add T2A machine types to validMachineTypes
- [#233](https://github.com/terraform-linters/tflint-ruleset-google/pull/233): Add M3 machine types to validMachineTypes

## 0.22.0 (2022-12-05)

### Enhancements

- [#231](https://github.com/terraform-linters/tflint-ruleset-google/pull/231): Update Magic Modules

### Chores

- [#227](https://github.com/terraform-linters/tflint-ruleset-google/pull/227): Add signatures for keyless signing
- [#228](https://github.com/terraform-linters/tflint-ruleset-google/pull/228): Bump google.golang.org/api from 0.100.0 to 0.103.0
- [#229](https://github.com/terraform-linters/tflint-ruleset-google/pull/229): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.24.0 to 2.24.1
- [#230](https://github.com/terraform-linters/tflint-ruleset-google/pull/230): Bump github.com/hashicorp/hcl/v2 from 2.14.1 to 2.15.0

## 0.21.0 (2022-10-28)

### Enhancements

- [#222](https://github.com/terraform-linters/tflint-ruleset-google/pull/222): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.12.0 to 0.14.0
- [#224](https://github.com/terraform-linters/tflint-ruleset-google/pull/224): Update Magic Modules

### Chores

- [#220](https://github.com/terraform-linters/tflint-ruleset-google/pull/220): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.21.0 to 2.24.0
- [#221](https://github.com/terraform-linters/tflint-ruleset-google/pull/221): Bump github.com/dave/dst from 0.27.0 to 0.27.2
- [#223](https://github.com/terraform-linters/tflint-ruleset-google/pull/223): Bump google.golang.org/api from 0.95.0 to 0.100.0

## 0.20.0 (2022-09-08)

The minimum supported version of TFLint has changed in this version. TFLint v0.40.0+ is required for this plugin to work.

### Breaking Changes

- [#207](https://github.com/terraform-linters/tflint-ruleset-google/pull/207): Bump tflint-plugin-sdk to v0.12.0
- [#209](https://github.com/terraform-linters/tflint-ruleset-google/pull/209): Update Magic Modules
  - Removed `google_certificate_manager_certificate_invalid_scope` rule
  - Removed `google_certificate_manager_certificate_map_entry_invalid_matcher` rule

### Chores

- [#199](https://github.com/terraform-linters/tflint-ruleset-google/pull/199): go 1.19
- [#202](https://github.com/terraform-linters/tflint-ruleset-google/pull/202): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.20.0 to 2.21.0
- [#204](https://github.com/terraform-linters/tflint-ruleset-google/pull/204) [#208](https://github.com/terraform-linters/tflint-ruleset-google/pull/208): Bump google.golang.org/api from 0.88.0 to 0.95.0
- [#205](https://github.com/terraform-linters/tflint-ruleset-google/pull/205): Bump github.com/hashicorp/hcl/v2 from 2.13.0 to 2.14.0
- [#206](https://github.com/terraform-linters/tflint-ruleset-google/pull/206): build: Use `go-version-file` instead of `go-version`

## 0.19.0 (2022-07-31)

### Enhancements

- [#195](https://github.com/terraform-linters/tflint-ruleset-google/pull/195): Add c2d machines to validMachineTypes
- [#197](https://github.com/terraform-linters/tflint-ruleset-google/pull/197): Update Magic Modules

### Chores

- [#187](https://github.com/terraform-linters/tflint-ruleset-google/pull/187): Bump github.com/hashicorp/hcl/v2 from 2.12.0 to 2.13.0
- [#194](https://github.com/terraform-linters/tflint-ruleset-google/pull/194): Bump google.golang.org/api from 0.84.0 to 0.88.0
- [#196](https://github.com/terraform-linters/tflint-ruleset-google/pull/196): Bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.17.0 to 2.20.0

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

