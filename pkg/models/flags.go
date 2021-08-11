package models

// BasicFlags contains objects that are needed for every command
type BasicFlags struct {
	InputFile  string
	OutputFile string // OutputFile is not required. If left blank, print to stdout
}

// B64CmdFlags are flags specific to
type B64CmdFlags struct {
}
