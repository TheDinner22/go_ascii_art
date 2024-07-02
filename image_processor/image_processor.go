package imageprocessor

import (
	"errors"
	"fmt"
	"image"
)

var chars = [...]rune{'`', '^', '\'', ',', ':', ';', 'I', 'l', '!', 'i', '~', '+', '_', '-', '?', ']', '[', '}', '{', '1', ')', '(', '|', '\\', '/', 't', 'f', 'j', 'r', 'x', 'n', 'u', 'v', 'c', 'z', 'X', 'Y', 'U', 'J', 'C', 'L', 'Q', '0', 'O', 'Z', 'm', 'w', 'q', 'p', 'd', 'b', 'k', 'h', 'a', 'o', '*', '#', 'M', 'W', '&', '8', '%', 'B', '@', '$'}

func GetBrightnessMatrix(img image.Image) ([][]float64, error) {
	if img == nil {
		return nil, errors.New("image was nil?!")
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	brightnessMatrix := make([][]float64, height)

	for i := 0; i < height; i++ {
		row := make([]float64, width)
		for j := 0; j < width; j++ {
            r, g, b, _ := img.At(i, j).RGBA()
            average_brightness := (float64(r)+float64(g)+float64(b))/ float64(3)
            normalized_average_brightness := average_brightness / float64(255)
            row[j] = normalized_average_brightness // ranges from 0-1 inclusive
		}
		brightnessMatrix[i] = row
	}

	return brightnessMatrix, nil
}

func Test() {
	fmt.Println("processor")
}
