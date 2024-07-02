package imageprocessor

var chars = [...]rune{'`', '^', '\'', ',', ':', ';', 'I', 'l', '!', 'i', '~', '+', '_', '-', '?', ']', '[', '}', '{', '1', ')', '(', '|', '\\', '/', 't', 'f', 'j', 'r', 'x', 'n', 'u', 'v', 'c', 'z', 'X', 'Y', 'U', 'J', 'C', 'L', 'Q', '0', 'O', 'Z', 'm', 'w', 'q', 'p', 'd', 'b', 'k', 'h', 'a', 'o', '*', '#', 'M', 'W', '&', '8', '%', 'B', '@', '$'}

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

