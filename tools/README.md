# Automation tools

This directory contains scripts for automating various operations.

## Magic Modules

[Magic Modules](https://github.com/GoogleCloudPlatform/magic-modules) is a tool used to autogenerate Terraform Provider for GCP. We have [forked this project](https://github.com/terraform-linters/magic-modules) to allow the validation to automatically generate rules for TFLint.

See the README for how to use this module. You can generate rules with the command like the following:

```console
$ bundle exec compiler -a --force tflint -v ga -e terraform -o $GOPATH/src/github.com/terraform-linters/tflint-ruleset-google/rules/magicmodules
```
