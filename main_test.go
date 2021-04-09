package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePath(t *testing.T) {
	tests := map[string]struct {
		fullPath string
		path     string
		file     string
	}{
		"same dir": {
			fullPath: "some-pic.jpg",
			path:     "",
			file:     "some-pic.jpg",
		},
		"1 deep dir": {
			fullPath: "somedir/some-pic.jpg",
			path:     "somedir",
			file:     "some-pic.jpg",
		},
		"3 deep path": {
			fullPath: "abc/123/xyz/picture.png",
			path:     "abc/123/xyz",
			file:     "picture.png",
		},
	}

	for _, test := range tests {
		path, file := parsePath(test.fullPath)
		assert.Equal(t, path, test.path)
		assert.Equal(t, file, test.file)
	}
}

func TestBytes(t *testing.T) {
	tests := map[string]struct {
		bytes     int
		formatted string
	}{
		"big size": {
			bytes:     353370112,
			formatted: "353.4 MB",
		},
		"small size": {
			bytes:     124,
			formatted: "124 B",
		},
		"kilo": {
			bytes:     9284,
			formatted: "9.3 kB",
		},
		"iphone pic": {
			bytes:     4823449,
			formatted: "4.8 MB",
		},
	}

	for _, test := range tests {
		f := bytes(test.bytes)
		assert.Equal(t, test.formatted, f)
	}
}
