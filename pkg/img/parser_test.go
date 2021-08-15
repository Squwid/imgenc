package img

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Squwid/imgenc/pkg/printer"
	"github.com/stretchr/testify/assert"
)

func TestInitialParser(t *testing.T) {
	tests := map[string]struct {
		filePath        string
		scans           int // number of lines to scan, -1 means all
		expectedOutputs []string
	}{
		"simple_test": {
			filePath: "../../tests/devito-smaller.jpg",
			scans:    3,
			expectedOutputs: []string{
				"0xff 0xd8",
				"0xff 0xdb 0x00 0x84 0x00 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x02 0x02 0x02 0x02 0x02 0x02 0x02 0x02 0x02 0x02 0x02 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x01 0x01 0x01 0x01 0x01 0x01 0x01 0x02 0x01 0x01 0x02 0x02 0x02 0x01 0x02 0x02 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03 0x03",
				"0xff 0xdd 0x00 0x04 0x00 0x7d",
			},
		},
	}

	for testName, test := range tests {
		file, err := os.Open(test.filePath)
		if err != nil {
			t.Logf("[%v] failed opening file %v\n", testName, err)
			t.Fail()
		}
		defer file.Close()

		s := bufio.NewScanner(file)

		fStats, err := file.Stat()
		if err != nil {
			t.Logf("error getting file stats: %v\n", err)
			t.Fail()
		}

		buffer := []byte{}
		s.Buffer(buffer, int(fStats.Size()))

		parser := NewParser()
		s.Split(parser.Split)

		var i = 0
		for s.Scan() {
			bs := s.Bytes()
			assert.Equal(t, test.expectedOutputs[i], printer.SprintBytes(bs))

			i++
			if i == test.scans {
				break
			}
		}
	}
}

func TestWorking(t *testing.T) {
	file, err := os.Open("../../tests/devito-smaller.jpg")
	assert.Equal(t, nil, err)
	defer file.Close()

	s := bufio.NewScanner(file)

	fStats, err := file.Stat()
	assert.Equal(t, nil, err)

	buffer := []byte{}
	s.Buffer(buffer, int(fStats.Size()))

	parser := NewParser()
	s.Split(parser.Split)

	var markers = map[byte]int{}
	for s.Scan() {
		bs := s.Bytes()

		if _, ok := markers[bs[1]]; ok {
			markers[bs[1]]++
		} else {
			markers[bs[1]] = 1
		}
	}

	for k, v := range markers {
		fmt.Printf("[0x%02x]: %v\n", k, v)
	}
}

func TestScannerParser(t *testing.T) {
	tests := map[string]struct {
		file string
	}{
		"devito": {
			file: "../../tests/devito-smaller.jpg",
		},
		"truck": {
			file: "../../tests/dumptruck.jpg",
		},
	}

	for _, test := range tests {
		// OG File Things
		ogFile, err := os.Open(test.file)
		assert.Equal(t, nil, err)
		defer ogFile.Close()
		ogBS, err := io.ReadAll(ogFile)
		assert.Equal(t, nil, err)

		// File from scanner
		file, err := os.Open(test.file)
		assert.Equal(t, nil, err)
		defer file.Close()
		s := bufio.NewScanner(file)
		fStats, err := file.Stat()
		assert.Equal(t, nil, err)

		buffer := []byte{}
		s.Buffer(buffer, int(fStats.Size()))

		parser := NewParser()
		s.Split(parser.Split)

		var out = []byte{}
		for s.Scan() {
			bs := s.Bytes()
			out = append(out, bs...)
		}

		// Make sure all the bytes are the same between file and parser
		assert.Equal(t, len(ogBS), len(out))

		if len(ogBS) == len(out) {
			for i := 0; i < len(ogBS); i++ {
				assert.Equal(t, ogBS[i], out[i])
			}
		}
	}
}
