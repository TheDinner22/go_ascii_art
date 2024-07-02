package canvas

import (
	"image"
	"io"
	"os"

	"golang.org/x/term"
)

// so you can make your own versions of this for different outputs but one for the terminal is made for you
type Canvas interface{
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
    if imageHeightSmallerThanCanvasHeight && imageWidthSmallerThanCanvasWidth {
        // TODO image to slice function plus return
    }

    // shrink the image somehow
}

// todo crash if not a terminal as stdout
type Terminal struct { }

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

