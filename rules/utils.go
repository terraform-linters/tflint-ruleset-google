package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var validMachineTypes = map[string]bool{
	// Regular machine types: https://cloud.google.com/compute/docs/machine-types
	// C2
	"c2-standard-4":    true,
	"c2-standard-8":    true,
	"c2-standard-16":   true,
	"c2-standard-30":   true,
	"c2-standard-60":   true,
	"c2d-standard-2":   true,
	"c2d-standard-4":   true,
	"c2d-standard-8":   true,
	"c2d-standard-16":  true,
	"c2d-standard-32":  true,
	"c2d-standard-56":  true,
	"c2d-standard-112": true,
	"c2d-highmem-2":    true,
	"c2d-highmem-4":    true,
	"c2d-highmem-8":    true,
	"c2d-highmem-16":   true,
	"c2d-highmem-32":   true,
	"c2d-highmem-56":   true,
	"c2d-highmem-112":  true,
	"c2d-highcpu-2":    true,
	"c2d-highcpu-4":    true,
	"c2d-highcpu-8":    true,
	"c2d-highcpu-16":   true,
	"c2d-highcpu-32":   true,
	"c2d-highcpu-56":   true,
	"c2d-highcpu-112":  true,
	// C3
	"c3-standard-4":         true,
	"c3-standard-4-lssd":    true,
	"c3-standard-8":         true,
	"c3-standard-8-lssd":    true,
	"c3-standard-22":        true,
	"c3-standard-22-lssd":   true,
	"c3-standard-44":        true,
	"c3-standard-44-lssd":   true,
	"c3-standard-88":        true,
	"c3-standard-88-lssd":   true,
	"c3-standard-176":       true,
	"c3-standard-176-lssd":  true,
	"c3-highmem-4":          true,
	"c3-highmem-8":          true,
	"c3-highmem-22":         true,
	"c3-highmem-44":         true,
	"c3-highmem-88":         true,
	"c3-highmem-176":        true,
	"c3-highcpu-4":          true,
	"c3-highcpu-8":          true,
	"c3-highcpu-22":         true,
	"c3-highcpu-44":         true,
	"c3-highcpu-88":         true,
	"c3-highcpu-176":        true,
	"c3d-standard-4":        true,
	"c3d-standard-8":        true,
	"c3d-standard-8-lssd":   true,
	"c3d-standard-16":       true,
	"c3d-standard-16-lssd":  true,
	"c3d-standard-30":       true,
	"c3d-standard-30-lssd":  true,
	"c3d-standard-60":       true,
	"c3d-standard-60-lssd":  true,
	"c3d-standard-90":       true,
	"c3d-standard-90-lssd":  true,
	"c3d-standard-180":      true,
	"c3d-standard-180-lssd": true,
	"c3d-standard-360":      true,
	"c3d-standard-360-lssd": true,
	"c3d-highmem-4":         true,
	"c3d-highmem-8":         true,
	"c3d-highmem-8-lssd":    true,
	"c3d-highmem-16":        true,
	"c3d-highmem-16-lssd":   true,
	"c3d-highmem-30":        true,
	"c3d-highmem-30-lssd":   true,
	"c3d-highmem-60":        true,
	"c3d-highmem-60-lssd":   true,
	"c3d-highmem-90":        true,
	"c3d-highmem-90-lssd":   true,
	"c3d-highmem-180":       true,
	"c3d-highmem-180-lssd":  true,
	"c3d-highmem-360":       true,
	"c3d-highmem-360-lssd":  true,
	"c3d-highcpu-4":         true,
	"c3d-highcpu-8":         true,
	"c3d-highcpu-16":        true,
	"c3d-highcpu-30":        true,
	"c3d-highcpu-60":        true,
	"c3d-highcpu-90":        true,
	"c3d-highcpu-180":       true,
	"c3d-highcpu-360":       true,
	// E2
	"e2-micro":       true,
	"e2-small":       true,
	"e2-medium":      true,
	"e2-standard-2":  true,
	"e2-standard-4":  true,
	"e2-standard-8":  true,
	"e2-standard-16": true,
	"e2-standard-32": true,
	"e2-highmem-2":   true,
	"e2-highmem-4":   true,
	"e2-highmem-8":   true,
	"e2-highmem-16":  true,
	"e2-highcpu-2":   true,
	"e2-highcpu-4":   true,
	"e2-highcpu-8":   true,
	"e2-highcpu-16":  true,
	"e2-highcpu-32":  true,
	// F1
	"f1-micro": true,
	// G1
	"g1-small": true,
	// H3
	"h3-standard-88": true,
	// M1
	"m1-megamem-96":   true,
	"m1-ultramem-40":  true,
	"m1-ultramem-80":  true,
	"m1-ultramem-160": true,
	// M2
	"m2-megamem-416":  true,
	"m2-hypermem-416": true,
	"m2-ultramem-208": true,
	"m2-ultramem-416": true,
	// M3
	"m3-megamem-64":   true,
	"m3-megamem-128":  true,
	"m3-ultramem-32":  true,
	"m3-ultramem-64":  true,
	"m3-ultramem-128": true,
	// N1
	"n1-standard-1":   true,
	"n1-standard-2":   true,
	"n1-standard-4":   true,
	"n1-standard-8":   true,
	"n1-standard-16":  true,
	"n1-standard-32":  true,
	"n1-standard-64":  true,
	"n1-standard-96":  true,
	"n1-highmem-2":    true,
	"n1-highmem-4":    true,
	"n1-highmem-8":    true,
	"n1-highmem-16":   true,
	"n1-highmem-32":   true,
	"n1-highmem-64":   true,
	"n1-highmem-96":   true,
	"n1-megamem-96":   true,
	"n1-ultramem-40":  true,
	"n1-ultramem-80":  true,
	"n1-ultramem-160": true,
	"n1-highcpu-2":    true,
	"n1-highcpu-4":    true,
	"n1-highcpu-8":    true,
	"n1-highcpu-16":   true,
	"n1-highcpu-32":   true,
	"n1-highcpu-64":   true,
	"n1-highcpu-96":   true,
	// N2
	"n2-standard-2":    true,
	"n2-standard-4":    true,
	"n2-standard-8":    true,
	"n2-standard-16":   true,
	"n2-standard-32":   true,
	"n2-standard-48":   true,
	"n2-standard-64":   true,
	"n2-standard-80":   true,
	"n2-standard-96":   true,
	"n2-standard-128":  true,
	"n2-highmem-2":     true,
	"n2-highmem-4":     true,
	"n2-highmem-8":     true,
	"n2-highmem-16":    true,
	"n2-highmem-32":    true,
	"n2-highmem-48":    true,
	"n2-highmem-64":    true,
	"n2-highmem-80":    true,
	"n2-highmem-96":    true,
	"n2-highmem-128":   true,
	"n2-highcpu-2":     true,
	"n2-highcpu-4":     true,
	"n2-highcpu-8":     true,
	"n2-highcpu-16":    true,
	"n2-highcpu-32":    true,
	"n2-highcpu-48":    true,
	"n2-highcpu-64":    true,
	"n2-highcpu-80":    true,
	"n2-highcpu-96":    true,
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

	//n4
	"n4-standard-2":  true,
	"n4-standard-4":  true,
	"n4-standard-8":  true,
	"n4-standard-16": true,
	"n4-standard-32": true,
	"n4-standard-48": true,
	"n4-standard-64": true,
	"n4-standard-80": true,
	"n4-highcpu-2":   true,
	"n4-highcpu-4":   true,
	"n4-highcpu-8":   true,
	"n4-highcpu-16":  true,
	"n4-highcpu-32":  true,
	"n4-highcpu-48":  true,
	"n4-highcpu-64":  true,
	"n4-highcpu-80":  true,
	"n4-highmem-2":   true,
	"n4-highmem-4":   true,
	"n4-highmem-8":   true,
	"n4-highmem-16":  true,
	"n4-highmem-32":  true,
	"n4-highmem-48":  true,
	"n4-highmem-64":  true,
	"n4-highmem-80":  true,

	// T2
	"t2a-standard-1":  true,
	"t2a-standard-2":  true,
	"t2a-standard-4":  true,
	"t2a-standard-8":  true,
	"t2a-standard-16": true,
	"t2a-standard-32": true,
	"t2a-standard-48": true,
	"t2d-standard-1":  true,
	"t2d-standard-2":  true,
	"t2d-standard-4":  true,
	"t2d-standard-8":  true,
	"t2d-standard-16": true,
	"t2d-standard-32": true,
	"t2d-standard-48": true,
	"t2d-standard-60": true,

	// A100 machine types: https://cloud.google.com/compute/docs/gpus
	"a2-highgpu-1g":  true,
	"a2-highgpu-2g":  true,
	"a2-highgpu-4g":  true,
	"a2-highgpu-8g":  true,
	"a2-megagpu-16g": true,
	"a2-ultragpu-1g": true,
	"a2-ultragpu-2g": true,
	"a2-ultragpu-4g": true,
	"a2-ultragpu-8g": true,

	// H100 machine types: https://cloud.google.com/compute/docs/gpus#h100-gpus
	"a3-highgpu-8g": true,

	// L4 machine types: https://cloud.google.com/compute/docs/gpus#l4-gpus
	"g2-standard-4":  true,
	"g2-standard-8":  true,
	"g2-standard-12": true,
	"g2-standard-16": true,
	"g2-standard-24": true,
	"g2-standard-32": true,
	"g2-standard-48": true,
	"g2-standard-96": true,

	// CT TPU machine types https://cloud.google.com/tpu/docs/tpus-in-gke
	"ct4p-hightpu-4t":  true,
	"ct5l-hightpu-1t":  true,
	"ct5l-hightpu-4t":  true,
	"ct5l-hightpu-8t":  true,
	"ct5lp-hightpu-1t": true,
	"ct5lp-hightpu-4t": true,
	"ct5lp-hightpu-8t": true,

	// Z3
	"z3-highmem-88":  true,
	"z3-highmem-176": true,
}

func isCustomType(machineType string) bool {
	return strings.HasPrefix(machineType, "e2-custom-") ||
		strings.HasPrefix(machineType, "n2-custom-") ||
		strings.HasPrefix(machineType, "n2d-custom-") ||
		strings.HasPrefix(machineType, "n1-custom-") ||
		strings.HasPrefix(machineType, "n4-custom-") ||
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
