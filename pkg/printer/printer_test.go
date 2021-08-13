package printer

import "testing"

func TestPrintOutput(t *testing.T) {
	tests := map[string]struct {
		input []byte
	}{
		"only_test": {
			input: []byte{'b', 'e', 0x4f, 1},
		},
	}

	for k, v := range tests {
		_ = k
		PrintBytes(v.input)
	}
}

func TestSprintBytes(t *testing.T) {
	tests := map[string]struct {
		input  []byte
		output string
	}{
		"simple_hex": {
			input:  []byte{0xFF, 0x91, 0x33},
			output: "0xff 0x91 0x33",
		},
	}

	for k, v := range tests {
		out := SprintBytes(v.input)
		if out != v.output {
			t.Logf("[%v] expected output %s but got %s\n", k, v.output, out)
			t.Fail()
		}
	}
}
