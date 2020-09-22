# TFLint Ruleset for terraform-provider-google
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-google/workflows/build/badge.svg?branch=master)](https://github.com/terraform-linters/tflint-ruleset-google/actions)

TFLint ruleset plugin for Terraform Google Cloud Platform provider

## Requirements

- TFLint v0.20+
- Go v1.15

## Installation

Download the plugin and place it in `~/.tflint.d/plugins/tflint-ruleset-google` (or `./.tflint.d/plugins/tflint-ruleset-google`). When using the plugin, configure as follows in `.tflint.hcl`:

```hcl
plugin "google" {
    enabled = true
}
```

## Rules

100+ rules are available. See the [documentation](docs/README.md).

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
