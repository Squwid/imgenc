package parser

import (
	"fmt"
	"testing"
)

func TestParseExtensions(t *testing.T) {
	tests := map[string]struct {
		file     string
		expected string
	}{
		"heic": {
			file:     "something.asdfasdf.heic",
			expected: "heic",
		},
		"heic_wrong": {
			file:     "abcasdfasdf.heic.a",
			expected: "",
		},
		"png": {
			file:     "thisisaphoto.png",
			expected: "png",
		},
		"png_wrong": {
			file:     ".png.a123123:2343424,234234.adsf.png.a",
			expected: "",
		},
		"jpeg": {
			file:     ".asdf.df.adf.asdfasd;fasdfasd.fa.dsf.jpeg",
			expected: "jpeg",
		},
		"jpeg_wrong": {
			file:     ".asdf.df.adf.asdfasd;fasdfasd.fa.dsf",
			expected: "",
		},
		"jpg": {
			file:     "asdkjfalksdfjalksdfja;lksdfjalksdjf..daf.asdf.as.df.a......jpg",
			expected: "jpg",
		},
		"jpg_wrong": {
			file:     "aksdjfalsdkfjalsdfa.dfasdfa.jpaa.df.asdf.jpg.",
			expected: "",
		},
	}

	for k, v := range tests {
		ext := ParseExtention(v.file)
		if ext != v.expected {
			t.Logf("Test %s failed, expected %s got %s\n", k, v.expected, ext)
			t.Fail()
		}
	}
}

func TestParsePath(t *testing.T) {
	tests := map[string]struct {
		path         string
		expectedPath string
		expectedFile string
	}{
		"simple": {
			path:         "testing/one/two/pic.png",
			expectedPath: "testing/one/two",
			expectedFile: "pic.png",
		},
		"no_path": {
			path:         "pic.png",
			expectedPath: "",
			expectedFile: "pic.png",
		},
		"single_path": {
			path:         "images/file.jpeg",
			expectedPath: "images",
			expectedFile: "file.jpeg",
		},
		"dir": {
			path:         "images/",
			expectedPath: "images",
			expectedFile: "",
		},
	}

	for k, v := range tests {
		path, file := ParsePath(v.path)
		if path != v.expectedPath {
			t.Logf("Expected path %s but got path %s (%s)\n", v.expectedPath, path, k)
			t.Fail()
		}

		if file != v.expectedFile {
			t.Logf("Expected file %s but got file %s (%s)\n", v.expectedFile, file, k)
			t.Fail()
		}
	}
}

func TestMisc(t *testing.T) {
	var bs = []byte{'B', 'e', 'n'}
	fmt.Println(bs)
}
