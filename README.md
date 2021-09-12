# TFLint Ruleset for terraform-provider-google
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-google/workflows/build/badge.svg?branch=master)](https://github.com/terraform-linters/tflint-ruleset-google/actions)
[![GitHub release](https://img.shields.io/github/release/terraform-linters/tflint-ruleset-google.svg)](https://github.com/terraform-linters/tflint-ruleset-google/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

TFLint ruleset plugin for Terraform Google Cloud Platform provider

## Requirements

- TFLint v0.30+
- Go v1.17

## Installation

You can install the plugin by adding a config to `.tflint.hcl` and running `tflint --init`:

```hcl
plugin "google" {
    enabled = true
    version = "0.12.1"
    source  = "github.com/terraform-linters/tflint-ruleset-google"
}
```

For more configuration about the plugin, see [Plugin Configuration](docs/configuration.md).

## Rules

100+ rules are available. See the [documentation](docs/rules/README.md).

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

Note that if you install the plugin with `make install`, you must omit the `version` and `source` attributes in `.tflint.hcl`:

```hcl
plugin "google" {
    enabled = true
}
```

## Add a new rule

If you are interested in adding a new rule to this ruleset, you can use the generator. Run the following command:

```
$ go run ./rules/generator
```

Follow the instructions to edit the generated files and open a new pull request.
