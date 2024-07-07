package cli

import (
	"flag"
	"log"
)

type percent uint8

func toPercent(num *int) percent{
    if *num > 100 {
        log.Fatalf("Width and height cannot be greater than 100 or less than 0.")
    }
    return percent(*num)
}

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
    width percent
    height percent
    output outputMode
    pixelWidth *uint // using ptr's here so these can be nil (using the ptr like an option)
    pixelHeight *uint
}

func GetArgs() Args{
    arg_width := flag.Int("width", 100, "Value from 0-100 inclusive. Width of the image. A width of 0 means preserve aspect ratio with respect to the specified height.")
    arg_height := flag.Int("height", 100, "Value from 0-100 inclusive. Height of the image. A height of 0 means preserve aspect ratio with respect to the specified width.")
    arg_output := flag.String("output", "terminal", "Where to output/write the ascii image (terminal|file).")

    arg_pixelHeight := flag.Uint("PHeight", 0, "Height of the output in pixel's")
    arg_pixelWidth := flag.Uint("PWidth", 0, "")

    // scaling func
    // char set
    // img path
    // verbose mode

    width := toPercent(arg_width)
    height := toPercent(arg_height)
    output := toOutputMode(arg_output)

    return Args{
    	width:  width,
    	height: height,
    	output: output,
    }
}
