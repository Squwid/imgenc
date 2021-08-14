package img

import (
	"bytes"
)

// Parser is the JPEG token parser
type Parser struct {
	buf     bytes.Buffer
	advance int
}

// NewParser returns default values for the parser
func NewParser() *Parser {
	return &Parser{}
}

// Reset resets the parsers fields
func (p *Parser) Reset() {
	p.buf.Reset()
	p.advance = 0
}

func (p *Parser) Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i, b := range data {
		if b == MarkerIdentifier {
			if i != 0 {
				token = p.buf.Bytes()
				advance = len(token)
				p.Reset()
				// fmt.Println("Advance", advance)
				// fmt.Println("atEOF", atEOF)
				// fmt.Printf("0x%02x - 0x%02x\n", b, MarkerIdentifier)
				// printer.PrintBytes(token)
				// fmt.Println()
				return
			}
		}

		p.advance++
		p.buf.WriteByte(b)
	}

	return 0, nil, nil
}
