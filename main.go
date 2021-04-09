package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.NewEntry(logrus.New())
	if len(os.Args) <= 1 {
		logger.Fatalf("Expected image argument")
	}

	imageFile := os.Args[1]
	logger = logger.WithField("Image", imageFile)

	var outputFile string
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
		logger = logger.WithField("Output", outputFile)
	}

	file, err := os.Open(imageFile)
	if err != nil {
		logger.WithError(err).Fatalf("Error opening image file")
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		logger.WithError(err).Fatalf("Error reading file")
	}

	encoded := base64.StdEncoding.EncodeToString(bs)
	if outputFile == "" {
		fmt.Println(encoded)
	} else {
		path, _ := parsePath(outputFile)

		// Ensure path exists, if not create
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0700); err != nil {
				logger.WithError(err).Fatalf("Error creating dir %s", path)
			}
		}

		// Save to output file
		outFile, err := os.OpenFile(outputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			logger.WithError(err).Fatalf("Error creating output file")
		}
		defer outFile.Close()

		n, err := outFile.WriteString(encoded)
		if err != nil {
			logger.WithError(err).Fatalf("Error writing to file")
		}
		logger.Infof("Wrote %v to file", bytes(n))
	}
}

func bytes(n int) string {
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

func parsePath(fullPath string) (path string, file string) {
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
