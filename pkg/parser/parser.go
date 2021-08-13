package parser

import (
	"strings"

	"github.com/mitchellh/go-homedir"
)

// Possible image extensions
const (
	ExtHEIC = "heic"
	ExtPNG  = "png"
	ExtJPG  = "jpg"
	ExtJPEG = "jpeg"
	ExtPSD  = "psd"
)

// ParsePath takes a relative path and creates an absolute path
func ParsePath(fullPath string) (path string, file string) {
	f, err := homedir.Expand(fullPath)
	if err == nil {
		fullPath = f
	}

	ss := strings.Split(fullPath, "/")
	if len(ss) == 1 {
		return "", fullPath
	}

	file = ss[len(ss)-1]
	ss = ss[:len(ss)-1]
	path = strings.Join(ss, "/")
	return
}

// ParseExtention returns the path extension of the image. Returns the following:
//   png, jpeg, jpg, heic, psd
func ParseExtention(file string) string {
	ss := strings.Split(file, ".")
	if len(ss) == 0 {
		return "'"
	}

	ext := ss[len(ss)-1]
	switch ext {
	case ExtHEIC, ExtJPEG, ExtJPG, ExtPNG, ExtPSD:
		return ext
	default:
		return ""
	}
}
