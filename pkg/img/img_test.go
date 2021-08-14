package img

import (
	"io/ioutil"
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

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		t.FailNow()
	}

	printer.PrintBytes(bs)
}

func TestMisc(t *testing.T) {

}
