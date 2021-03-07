# TFLint Ruleset for terraform-provider-google
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-google/workflows/build/badge.svg?branch=master)](https://github.com/terraform-linters/tflint-ruleset-google/actions)
[![GitHub release](https://img.shields.io/github/release/terraform-linters/tflint-ruleset-google.svg)](https://github.com/terraform-linters/tflint-ruleset-google/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

TFLint ruleset plugin for Terraform Google Cloud Platform provider

## Requirements

- TFLint v0.24+
- Go v1.16

## Installation

Download the plugin and place it in `~/.tflint.d/plugins/tflint-ruleset-google` (or `./.tflint.d/plugins/tflint-ruleset-google`). When using the plugin, configure as follows in `.tflint.hcl`:

```hcl
plugin "google" {
    enabled = true
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
