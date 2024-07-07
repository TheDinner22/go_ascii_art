package main

import (
	"log"

	canvas "github.com/TheDinner22/go_ascii_art/canvas"
	cli "github.com/TheDinner22/go_ascii_art/cli"
	imageDisplayer "github.com/TheDinner22/go_ascii_art/image_displayer"
	imageProcessor "github.com/TheDinner22/go_ascii_art/image_processor"
	imageReader "github.com/TheDinner22/go_ascii_art/image_reader"
)

func main(){
    args := cli.GetArgs()

    //img, err := imageReader.LoadFromFile("images/UF.png")
    //img, err := imageReader.LoadFromFile("images/mom.jpg")
    //img, err := imageReader.LoadFromFile("images/dog.jpg")
    //img, err := imageReader.LoadFromFile("images/mon.jpeg")
    //img, err := imageReader.LoadFromFile("images/goku.jpeg")
    //img, err := imageReader.LoadFromFile("images/builder.jpeg")
    img, err := imageReader.LoadFromFile(args.Src)
    if err != nil {
        log.Fatal(err.Error())
    }

    term := canvas.CanvasFromArgs(args, nil)

    bm := canvas.FitImageToCanvas(img, term)

    chars := imageProcessor.MapBrightnessMatrixToChars(bm)

    err = imageDisplayer.PrintImageToScreen(chars, term.Writer())
    if err != nil {
        panic(err.Error())
    }
}
