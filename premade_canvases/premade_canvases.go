package premadeCanvases

import (
    "os"
	"golang.org/x/term"
    "io"
)

// TODO crash if not a terminal as stdout
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
