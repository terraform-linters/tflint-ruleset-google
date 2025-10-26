package project

import "fmt"

// Version is ruleset version
const Version string = "0.37.1"

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/terraform-linters/tflint-ruleset-google/blob/v%s/docs/rules/%s.md", Version, name)
}
