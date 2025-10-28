package asciitize

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type (
	Asciitizer struct {
		writer io.Writer
	}
)

func NewAsciitizer(writer io.Writer, options ...func(*Asciitizer)) *Asciitizer {
	a := &Asciitizer{
		writer: writer,
	}
	for _, option := range options {
		option(a)
	}
	return a
}

func (a *Asciitizer) Asciitize(filepath string) error {
	return nil
}
