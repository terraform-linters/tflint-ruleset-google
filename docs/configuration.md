# Configuration

This plugin can take advantage of additional features by configure the `plugin` block. Currently, this configuration is only available for [Deep Checking](deep_checking.md).

Here's an example:

```hcl
plugin "google" {
    // Plugin common attributes

    deep_check = false
}
```

## `deep_check`

Default: false

Enable [Deep Checking](deep_checking.md).
