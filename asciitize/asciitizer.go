package asciitize

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

type (
	Decode     func(r io.Reader) (image.Image, error)
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
	img, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer img.Close()
	contents, _, err := image.Decode(img)
	if err != nil {
		return err
	}
	imgBounds := contents.Bounds()
	for y := imgBounds.Min.Y; y < imgBounds.Max.Y; y++ {
		for x := imgBounds.Min.X; x < imgBounds.Max.X; x++ {
			//r, g, b, _ := contents.At(x, y).RGBA()
		}
	}
	return nil
}
