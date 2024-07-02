package main

import (
    imageReader "github.com/TheDinner22/go_ascii_art/image_reader"
    imageDisplayer "github.com/TheDinner22/go_ascii_art/image_displayer"
    imageProcessor "github.com/TheDinner22/go_ascii_art/image_processor"
)

func main(){
    img, err := imageReader.LoadFromFile("images/UF.png")
    if err != nil {
        panic(err.Error())
    }

    bm, err := imageProcessor.GetBrightnessMatrix(img)

    if err != nil {
        panic(err.Error())
    }

    runes := imageProcessor.MapBrightnessMatrixToChars(bm)

    err = imageDisplayer.PrintImageToScreen(runes)

    if err != nil {
        panic(err.Error())
    }
}
