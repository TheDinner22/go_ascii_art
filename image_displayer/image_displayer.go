package imagedisplayer

import (
	"errors"
	"fmt"
	"io"
)

func PrintImageToScreen(asciiImage [][]rune, writer io.Writer) error {
	if asciiImage == nil || writer == nil {
		return errors.New("args were nil")
	}

	for _, row := range asciiImage {
            bs := []byte(string(row))
            _, err := writer.Write(bs)
            if err != nil {
                return err
            }
            writer.Write([]byte(string("\n")))
	}

	return nil
}

func Test() {
	fmt.Println("displayer")
}
