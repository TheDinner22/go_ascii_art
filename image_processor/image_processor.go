package imageprocessor

import (
	"errors"
	"fmt"
	"image"
)

var chars = [...]rune{'`', '^', '\'', ',', ':', ';', 'I', 'l', '!', 'i', '~', '+', '_', '-', '?', ']', '[', '}', '{', '1', ')', '(', '|', '\\', '/', 't', 'f', 'j', 'r', 'x', 'n', 'u', 'v', 'c', 'z', 'X', 'Y', 'U', 'J', 'C', 'L', 'Q', '0', 'O', 'Z', 'm', 'w', 'q', 'p', 'd', 'b', 'k', 'h', 'a', 'o', '*', '#', 'M', 'W', '&', '8', '%', 'B', '@', '$'}

func GetBrightnessMatrix(img image.Image) ([][]int, error) {
	if img == nil {
		return nil, errors.New("image was nil?!")
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	brightnessMatrix := make([][]int, height)

	for i := 0; i < height; i++ {
		row := make([]int, width)
		for j := 0; j < width; j++ {
            r, g, b, _ := img.At(i, j).RGBA()
            row[j] = (int(r)+int(g)+int(b))/3
		}
		brightnessMatrix[i] = row
	}

	return brightnessMatrix, nil
}

func Test() {
	fmt.Println("processor")
}
