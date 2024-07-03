package imagedisplayer

import "fmt"
import "errors"

func PrintImageToScreen(asciiImage [][]rune) error {
	if asciiImage == nil {
		return errors.New("slice of images was nil!?")
	}

	for _, row := range asciiImage {
		for _, char := range row {
			_, err := fmt.Print(string(char))
			if err != nil {
				return err
			}
		}
		println("")
	}

	return nil

}

func Test() {
	fmt.Println("displayer")
}
