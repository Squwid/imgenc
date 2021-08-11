package main

import (
	"os"

	"github.com/Squwid/imgenc/pkg/cmd/b64"
	"github.com/Squwid/imgenc/pkg/models"
	"github.com/alecthomas/kingpin"
	"github.com/sirupsen/logrus"
)

var version = "0.0.0"

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	app := kingpin.New("imgenc", "The best command line tool to perform commands on image files")
	app.Version(version)

	var debugMode bool
	app.Flag("debug", "Run in debug mode").Short('v').BoolVar(&debugMode)

	flagsBasic := new(models.BasicFlags)
	app.Flag("input", "Input file to transform (usually an image file)").Required().Short('i').StringVar(&flagsBasic.InputFile)
	app.Flag("output", "Output file to save the image. If left blank, output will print to stdout").Short('o').StringVar(&flagsBasic.OutputFile)

	// "b64e" Command (Base 64 Encode)
	cmdB64e := app.Command("b64e", "Base64 encode an image")
	cmdB64d := app.Command("b64d", "Base64 decode an image")
	flagsB64 := new(models.B64CmdFlags)

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	switch command {
	case cmdB64e.FullCommand():
		b64.EncodeCmd(flagsBasic, flagsB64)
	case cmdB64d.FullCommand():
		b64.DecodeCmd(flagsBasic, flagsB64)
	}
}
