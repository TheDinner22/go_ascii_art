package imageprocessor

import (
	"errors"
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
			average_brightness := (float64(r) + float64(g) + float64(b)) / float64(3)
            max_uint32 := ^uint32(0)
			normalized_average_brightness := average_brightness / float64(max_uint32)
			row[j] = normalized_average_brightness // ranges from 0-1 inclusive
		}
		brightnessMatrix[i] = row
	}

	return brightnessMatrix, nil
}

func MapBrightnessMatrixToChars(bm [][]float64) [][]rune {
	return mmap(bm, func(row []float64) []rune {
		return mmap(row, func(pixel float64) rune {
			index := int(pixel * 64) // there are 65 chars total
			return chars[index]
		})
	})
}

// https://stackoverflow.com/questions/71624828/is-there-a-way-to-map-an-array-of-objects-in-go
func mmap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
