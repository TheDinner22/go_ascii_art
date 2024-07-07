package canvas

import (
	"image"
	"io"
	"os"

	imageProcessor "github.com/TheDinner22/go_ascii_art/image_processor"

	"github.com/nfnt/resize"
	"golang.org/x/term"
)

// so you can make your own versions of this for different outputs but one for the terminal is made for you
type Canvas interface {
	WidthHeight() (int, int)
	Writer() io.Writer
}

func FitImageToCanvas(img image.Image, canvas Canvas) [][]float64 {
	// if the image is smaller than the canvas we can just convert it to a slice of floats
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	canvasWidth, canvasHeight := canvas.WidthHeight()
	imageWidthSmallerThanCanvasWidth := imgWidth < canvasWidth
	imageHeightSmallerThanCanvasHeight := imgHeight < canvasHeight

	// if the image is too big, resize it
	if !(imageHeightSmallerThanCanvasHeight && imageWidthSmallerThanCanvasWidth) {
		img = resize.Resize(uint(canvasWidth), 0, img, resize.NearestNeighbor)
	}

	// convert image to a slice (each element is a char to be printed)
	// (right now tho, they are values from 0-1 representing lightness or darkness)
	// 0 is light and 1 is dark
	return imageProcessor.ImageToSlice(img)
}

// todo crash if not a terminal as stdout
type Terminal struct{}

func (terminal Terminal) WidthHeight() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err.Error())
	}
	return width, height
}

func (terminal Terminal) Writer() io.Writer {
	return os.Stdout
}
