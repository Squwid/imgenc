package img

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/Squwid/imgenc/pkg/printer"
)

func TestImageParsing(t *testing.T) {
	const f = "../../tests/devito-smaller.jpg"
	file, err := os.Open(f)
	if err != nil {
		t.Logf("error opening file\n")
		t.FailNow()
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	fStats, err := file.Stat()
	if err != nil {
		t.Logf("error getting file stats : %v\n", err)
		t.FailNow()
	}

	buffer := []byte{}
	s.Buffer(buffer, int(fStats.Size())) // buffer has to be as large as the image worst case scenario

	// atEOF means no mare data to give, advance is number of bytes to advance data, scanning stops if an error occurs.
	// return (0 nil nil) if token end is not found

	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	t.Logf("error reading file : %v\n", err)
	// 	t.Fail()
	// }

	// fmt.Println("Expected:")
	// printer.PrintBytes(fileBytes[0:500])
	// fmt.Println("")

	var tkn bytes.Buffer
	var c = 0
	var start = true
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// TODO: first check if file is in jpeg format

		for _, b := range data {
			if b == MarkerIdentifier && start {
				start = false
			} else if b == MarkerIdentifier && !start {
				advance = c
				c = 0
				token = tkn.Bytes()
				tkn.Reset()
				start = true
				return
			}

			c++
			tkn.WriteByte(b)
		}

		return
	}

	s.Split(split)

	var i = 0

	for s.Scan() {
		i++
		bbs := s.Bytes()
		fmt.Printf("\nLength: %v\n", len(bbs))
		printer.PrintBytes(bbs)
		if i > 2 {
			break
		}
	}

	// bss := printer.SprintBytes(bs[4:136])
	// fmt.Println(len(bs[4:136]))
	// fmt.Println(bss)
}

func TestMisc(t *testing.T) {

}
