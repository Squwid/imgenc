package printer

import (
	"bytes"
	"fmt"
	"strings"
)

// PrintBytes prints a byte slice to stdout in hex format
func PrintBytes(bs []byte) {
	for _, b := range bs {
		fmt.Printf("0x%02x ", b)
	}
	fmt.Printf("\n")
}

// SprintBytes returns a byte slice as a string in hex format
func SprintBytes(bs []byte) string {
	var buf bytes.Buffer
	for _, b := range bs {
		if _, err := buf.WriteString(fmt.Sprintf("0x%02x ", b)); err != nil {
			// TODO: Really worth panicking here?
			panic(err)
		}
	}
	return strings.TrimRight(buf.String(), " ")
}
