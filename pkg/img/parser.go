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

func (p *Parser) Values() (int, []byte) {
	return p.advance, p.buf.Bytes()
}

func (p *Parser) Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if p.advance != 0 {
		data = data[p.advance:]
	}
	// printer.PrintBytes(data)
	for _, b := range data {
		if b == byte(0xff) {
			if p.advance != 0 {
				advance, token = p.Values()
				p.Reset()
				return
			}
		}

		p.advance++
		p.buf.WriteByte(b)
	}

	if atEOF && p.advance != 0 {
		advance, token = p.Values()
		p.Reset()
		return
	}

	return 0, nil, nil
}
