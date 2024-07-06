package main

import (
    imageReader "github.com/TheDinner22/go_ascii_art/image_reader"
    imageDisplayer "github.com/TheDinner22/go_ascii_art/image_displayer"
    imageProcessor "github.com/TheDinner22/go_ascii_art/image_processor"
    canvas "github.com/TheDinner22/go_ascii_art/canvas"
)

func main(){
    //img, err := imageReader.LoadFromFile("images/UF.png")
    //img, err := imageReader.LoadFromFile("images/mom.jpg")
    //img, err := imageReader.LoadFromFile("images/dog.jpg")
    //img, err := imageReader.LoadFromFile("images/mon.jpeg")
    //img, err := imageReader.LoadFromFile("images/goku.jpeg")
    img, err := imageReader.LoadFromFile("images/builder.jpeg")
    if err != nil {
        panic(err.Error())
    }

    term := canvas.Terminal{}
    bm := canvas.FitImageToCanvas(img, term)

    chars := imageProcessor.MapBrightnessMatrixToChars(bm)

    err = imageDisplayer.PrintImageToScreen(chars, term.Writer())
    if err != nil {
        panic(err.Error())
    }
}
