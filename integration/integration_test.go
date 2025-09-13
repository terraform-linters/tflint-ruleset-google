package integration

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type result struct {
	Issues []issue `json:"issues"`
	Errors []any   `json:"errors"`
}

type issue struct {
	Rule    any `json:"rule"`
	Message any `json:"message"`
	Range   any `json:"range"`
	Callers any `json:"callers"`

	// The following fields are ignored because they are added in TFLint v0.59.1.
	// We can uncomment this once the minimum supported version is v0.59.1+.
	// Fixed   any `json:"fixed"`
	// Fixable any `json:"fixable"`
}

func TestIntegration(t *testing.T) {
	cases := []struct {
		Name    string
		Command *exec.Cmd
		Dir     string
	}{
		{
			Name:    "basic",
			Command: exec.Command("tflint", "--format", "json", "--force"),
			Dir:     "basic",
		},
	}

	dir, _ := os.Getwd()
	defer os.Chdir(dir)

	for _, tc := range cases {
		testDir := dir + "/" + tc.Dir
		os.Chdir(testDir)

		var stdout, stderr bytes.Buffer
		tc.Command.Stdout = &stdout
		tc.Command.Stderr = &stderr
		if err := tc.Command.Run(); err != nil {
			t.Fatalf("Failed `%s`: %s, stdout=%s stderr=%s", tc.Name, err, stdout.String(), stderr.String())
		}

		ret, err := ioutil.ReadFile("result.json")
		if err != nil {
			t.Fatalf("Failed `%s`: %s", tc.Name, err)
		}

		var expected result
		if err := json.Unmarshal(ret, &expected); err != nil {
			t.Fatalf("Failed `%s`: %s", tc.Name, err)
		}

		var got result
		if err := json.Unmarshal(stdout.Bytes(), &got); err != nil {
			t.Fatalf("Failed `%s`: %s", tc.Name, err)
		}

		if !cmp.Equal(got, expected) {
			t.Fatalf("Failed `%s`: diff=%s", tc.Name, cmp.Diff(expected, got))
		}
	}
}
