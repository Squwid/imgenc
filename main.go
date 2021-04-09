package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.NewEntry(logrus.New())
	if len(os.Args) <= 1 {
		logger.Fatalf("Expected image argument")
	}

	imageFile := os.Args[1]
	logger = logger.WithField("Image", imageFile)

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
	fmt.Println(encoded)
}
