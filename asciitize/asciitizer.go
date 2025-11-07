package asciitize

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"golang.org/x/image/draw"
)

var (
	densityString = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
)

type (
	Decode     func(r io.Reader) (image.Image, error)
	Asciitizer struct {
		scale  float64
		writer io.Writer
	}
)

func NewAsciitizer(writer io.Writer, options ...func(*Asciitizer)) *Asciitizer {
	a := &Asciitizer{
		scale:  1.0,
		writer: writer,
	}
	for _, option := range options {
		option(a)
	}
	return a
}

func WithScale(scale float64) func(*Asciitizer) {
	return func(a *Asciitizer) {
		a.scale = scale
	}
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
	scaledContents := image.NewRGBA(image.Rect(0, 0, int(float64(contents.Bounds().Dx())*a.scale), int(float64(contents.Bounds().Dy())*a.scale)))
	draw.NearestNeighbor.Scale(scaledContents, scaledContents.Rect, contents, contents.Bounds(), draw.Over, nil)
	imgBounds := scaledContents.Bounds()
	asciiImage := ""
	for y := imgBounds.Min.Y; y < imgBounds.Max.Y; y++ {
		asciiRow := ""
		for x := imgBounds.Min.X; x < imgBounds.Max.X; x++ {
			r, g, b, _ := scaledContents.At(x, y).RGBA()
			brightness := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535
			index := int(brightness * float64(len(densityString)-1))
			asciiRow += string(densityString[index])
		}
		asciiImage += asciiRow + "\n"
	}
	_, err = a.writer.Write([]byte(asciiImage))
	fmt.Print(asciiImage)
	return err
}
