package canvas

import (
	"image"
	"io"

	cli "github.com/TheDinner22/go_ascii_art/cli"
    imageProcessor "github.com/TheDinner22/go_ascii_art/image_processor"
	"github.com/nfnt/resize"
)

// so you can make your own versions of this for different outputs but one for the terminal is made for you
type Canvas interface {
	WidthHeight() (uint, uint)
	Writer() io.Writer
}

// custom canvas (why would you do this)
// uses the args and a writer provided by someone (god ig)
type ArgCanvas struct {
    args cli.Args
    writer io.Writer
}

func (argCanvas ArgCanvas) WidthHeight() (uint, uint) {
    width := argCanvas.args.Width
    height := argCanvas.args.Height
	return width, height
}

func (argCanvas ArgCanvas) Writer() io.Writer {
	return argCanvas.writer
}


func CanvasFromArgs(args cli.Args, writer io.Writer) Canvas {
    return ArgCanvas{
        args: args,
        writer: writer,
    }
}

func FitImageToCanvas(img image.Image, canvas Canvas) [][]float64 {
	// if the image is smaller than the canvas we can just convert it to a slice of floats
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	canvasWidth, canvasHeight := canvas.WidthHeight()
	imageWidthSmallerThanCanvasWidth := imgWidth < int(canvasWidth)
	imageHeightSmallerThanCanvasHeight := imgHeight < int(canvasHeight)

	// if the image is too big, resize it
	if !(imageHeightSmallerThanCanvasHeight && imageWidthSmallerThanCanvasWidth) {
		img = resize.Resize(uint(canvasWidth), 0, img, resize.NearestNeighbor)
	}

	// convert image to a slice (each element is a char to be printed)
	// (right now tho, they are values from 0-1 representing lightness or darkness)
	// 0 is light and 1 is dark
	return imageProcessor.ImageToSlice(img)
}

