package models

import (
	"fmt"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// DisplayByte takes a number and displays it in a user friendly way
func DisplayByte(n int) string {
	const unit = 1000
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for n := n / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(n)/float64(div), "kMGTPE"[exp])
}

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
