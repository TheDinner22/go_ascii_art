package cli

import (
	"flag"
	"log"
	"math"
	"path"
)

type outputMode string

func toOutputMode(s *string) outputMode {
	switch *s {
	case "terminal":
		return outputMode(*s)
	case "file":
		return outputMode(*s)
	}

	log.Fatalf("%s is not a valid output mode.", *s)
	return outputMode("unreachable") // some reason go requires this :(
}

// where to output
// % width
// % height
type Args struct {
	Width  uint
	Height uint
	Output outputMode
	Src    string
}

func GetArgs() Args {
	// TODO what to do with these for optional stuff
	arg_width := flag.Uint(
		"width",
		math.MaxUint,
		"Width of the out put image (in pixels). A width of 0 means preserve aspect ratio with respect to the specified height. Both width and height cannot be 0.",
	)

	arg_height := flag.Uint(
		"width",
		math.MaxUint,
		"Height of the out put image (in pixels). A height of 0 means preserve aspect ratio with respect to the specified width. Both width and height cannot be 0.",
	)

	arg_output := flag.String("o", "terminal", "Where to output/write the ascii image (terminal|file).")

	arg_src_path := flag.String("img", "images/UF.png", "Source image to be converted to ascii. Path relative to go.mod for this module.")

	// scaling func
	// char set
	// img path
	// verbose mode

	flag.Parse()

	output := toOutputMode(arg_output)

	return Args{
		Width:  *arg_width,
		Height: *arg_height,
		Output: output,
		Src:    path.Clean(*arg_src_path),
	}
}

