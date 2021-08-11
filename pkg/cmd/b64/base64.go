package b64

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/Squwid/imgenc/pkg/models"
	"github.com/sirupsen/logrus"
)

func EncodeCmd(basicFlags *models.BasicFlags, b64Flags *models.B64CmdFlags) {
	logger := logrus.NewEntry(logrus.New()).WithFields(logrus.Fields{
		"Image": basicFlags.InputFile,
	})

	file, err := os.Open(basicFlags.InputFile)
	if err != nil {
		logger.WithError(err).Fatalf("Error opening image file")
	}
	defer file.Close()

	bs, err := io.ReadAll(file)
	if err != nil {
		logger.WithError(err).Fatalf("Error reading file")
	}

	encoded := base64.StdEncoding.EncodeToString(bs)
	if basicFlags.OutputFile == "" {
		// No output file, print to stdout
		fmt.Println(encoded)
	} else {
		path, _ := models.ParsePath(basicFlags.OutputFile)

		// Ensure path exists, if not create
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0700); err != nil {
				logger.WithError(err).Fatalf("Error creating dir %s", path)
			}
		}

		// Save to output file
		outFile, err := os.OpenFile(basicFlags.OutputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			logger.WithError(err).Fatalf("Error creating output file")
		}
		defer outFile.Close()

		n, err := outFile.WriteString(encoded)
		if err != nil {
			logger.WithError(err).Fatalf("Error writing to file")
		}
		logger.Infof("Wrote %v to file", models.DisplayByte(n))
	}

}
