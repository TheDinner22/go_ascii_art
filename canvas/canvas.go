package canvas

import (
	"image"
	"io"
	"os"

	"golang.org/x/term"
    "github.com/nfnt/resize"
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
    return imageToSlice(img)
}

func imageToSlice(img image.Image) [][]float64 {
	// max value that r g and b can be from the pixel color values
	const MAX_COLOR_VALUE float64 = 0xffff

	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	result := make([][]float64, imgHeight)
	for i := 0; i < imgHeight; i++ {
		row := make([]float64, imgWidth)
		for j := 0; j < imgWidth; j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			average_color := (float64(r) + float64(b) + float64(g)) / float64(3)
			normalized_average_color := average_color / MAX_COLOR_VALUE
			row[j] = normalized_average_color
		}
		result[i] = row
	}
	return result
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
