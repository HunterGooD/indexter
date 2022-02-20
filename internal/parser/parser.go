package parser

import (
	"io"
)

type Parser struct {
	files []File
}

type File struct {
	open  *io.Reader
	close func()
}

func (p *Parser) LoadDir() {

}
