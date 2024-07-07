package imageprocessor

import "image"

var chars = [...]rune{'`', '^', '\'', ',', ':', ';', 'I', 'l', '!', 'i', '~', '+', '_', '-', '?', ']', '[', '}', '{', '1', ')', '(', '|', '\\', '/', 't', 'f', 'j', 'r', 'x', 'n', 'u', 'v', 'c', 'z', 'X', 'Y', 'U', 'J', 'C', 'L', 'Q', '0', 'O', 'Z', 'm', 'w', 'q', 'p', 'd', 'b', 'k', 'h', 'a', 'o', '*', '#', 'M', 'W', '&', '8', '%', 'B', '@', '$'}

func MapBrightnessMatrixToChars(bm [][]float64) [][]rune {
	return mmap(bm, func(row []float64) []rune {
		return mmap(row, func(pixel float64) rune {
			index := int(pixel * 64) // there are 65 chars total
			return chars[index]
		})
	})
}

func ImageToSlice(img image.Image) [][]float64 {
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


// https://stackoverflow.com/questions/71624828/is-there-a-way-to-map-an-array-of-objects-in-go
func mmap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

