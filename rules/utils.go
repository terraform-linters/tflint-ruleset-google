package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var validMachineTypes = map[string]bool{
	"e2-standard-2":    true,
	"e2-standard-4":    true,
	"e2-standard-8":    true,
	"e2-standard-16":   true,
	"e2-standard-32":   true,
	"e2-highmem-2":     true,
	"e2-highmem-4":     true,
	"e2-highmem-8":     true,
	"e2-highmem-16":    true,
	"e2-highcpu-2":     true,
	"e2-highcpu-4":     true,
	"e2-highcpu-8":     true,
	"e2-highcpu-16":    true,
	"e2-highcpu-32":    true,
	"n2-standard-2":    true,
	"n2-standard-4":    true,
	"n2-standard-8":    true,
	"n2-standard-16":   true,
	"n2-standard-32":   true,
	"n2-standard-48":   true,
	"n2-standard-64":   true,
	"n2-standard-80":   true,
	"n2-highmem-2":     true,
	"n2-highmem-4":     true,
	"n2-highmem-8":     true,
	"n2-highmem-16":    true,
	"n2-highmem-32":    true,
	"n2-highmem-48":    true,
	"n2-highmem-64":    true,
	"n2-highmem-80":    true,
	"n2-highcpu-2":     true,
	"n2-highcpu-4":     true,
	"n2-highcpu-8":     true,
	"n2-highcpu-16":    true,
	"n2-highcpu-32":    true,
	"n2-highcpu-48":    true,
	"n2-highcpu-64":    true,
	"n2-highcpu-80":    true,
	"n2d-standard-2":   true,
	"n2d-standard-4":   true,
	"n2d-standard-8":   true,
	"n2d-standard-16":  true,
	"n2d-standard-32":  true,
	"n2d-standard-48":  true,
	"n2d-standard-64":  true,
	"n2d-standard-80":  true,
	"n2d-standard-96":  true,
	"n2d-standard-128": true,
	"n2d-standard-224": true,
	"n2d-highmem-2":    true,
	"n2d-highmem-4":    true,
	"n2d-highmem-8":    true,
	"n2d-highmem-16":   true,
	"n2d-highmem-32":   true,
	"n2d-highmem-48":   true,
	"n2d-highmem-64":   true,
	"n2d-highmem-80":   true,
	"n2d-highmem-96":   true,
	"n2d-highcpu-2":    true,
	"n2d-highcpu-4":    true,
	"n2d-highcpu-8":    true,
	"n2d-highcpu-16":   true,
	"n2d-highcpu-32":   true,
	"n2d-highcpu-48":   true,
	"n2d-highcpu-64":   true,
	"n2d-highcpu-80":   true,
	"n2d-highcpu-96":   true,
	"n2d-highcpu-128":  true,
	"n2d-highcpu-224":  true,
	"n1-standard-1":    true,
	"n1-standard-2":    true,
	"n1-standard-4":    true,
	"n1-standard-8":    true,
	"n1-standard-16":   true,
	"n1-standard-32":   true,
	"n1-standard-64":   true,
	"n1-standard-96":   true,
	"n1-highmem-2":     true,
	"n1-highmem-4":     true,
	"n1-highmem-8":     true,
	"n1-highmem-16":    true,
	"n1-highmem-32":    true,
	"n1-highmem-64":    true,
	"n1-highmem-96":    true,
	"n1-highcpu-2":     true,
	"n1-highcpu-4":     true,
	"n1-highcpu-8":     true,
	"n1-highcpu-16":    true,
	"n1-highcpu-32":    true,
	"n1-highcpu-64":    true,
	"n1-highcpu-96":    true,
	"c2-standard-4":    true,
	"c2-standard-8":    true,
	"c2-standard-16":   true,
	"c2-standard-30":   true,
	"c2-standard-60":   true,
	"m2-ultramem-208":  true,
	"m2-ultramem-416":  true,
	"m2-megamem-416":   true,
	"m1-ultramem-40":   true,
	"m1-ultramem-80":   true,
	"m1-ultramem-160":  true,
	"m1-megamem-96":    true,
	"n1-megamem-96":    true,
	"n1-ultramem-40":   true,
	"n1-ultramem-80":   true,
	"n1-ultramem-160":  true,
	"e2-micro":         true,
	"e2-small":         true,
	"e2-medium":        true,
	"f1-micro":         true,
	"g1-small":         true,
}

func isCustomType(machineType string) bool {
	return strings.HasPrefix(machineType, "e2-custom-") ||
		strings.HasPrefix(machineType, "n2-custom-") ||
		strings.HasPrefix(machineType, "n2d-custom-") ||
		strings.HasPrefix(machineType, "n1-custom-") ||
		strings.HasPrefix(machineType, "custom-")
}

func validateRegexp(re string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		if !regexp.MustCompile(re).MatchString(value) {
			errors = append(errors, fmt.Errorf(
				"%q (%q) doesn't match regexp %q", k, value, re))
		}

		return
	}
}

func isValidIAMMemberFormat(s string) bool {
	// See also https://cloud.google.com/iam/docs/overview
	return s == "allUsers" ||
		s == "allAuthenticatedUsers" ||
		strings.HasPrefix(s, "user:") ||
		strings.HasPrefix(s, "serviceAccount:") ||
		strings.HasPrefix(s, "group:") ||
		strings.HasPrefix(s, "domain:")
}
