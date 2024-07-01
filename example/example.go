package main

import (
    imageReader "github.com/TheDinner22/go_ascii_art/image_reader"
    imageDisplayer "github.com/TheDinner22/go_ascii_art/terminal_image_displayer"
    imageProcessor "github.com/TheDinner22/go_ascii_art/image_processor"
)

func processor(){
    imageProcessor.Test()
}

func displayer(){
    imageDisplayer.Test()
}

func reader(){
    //imageReader.LoadFromFile("images/thing.png")
    imageReader.Test()
}

func main(){
    displayer()
    reader()
    processor()
}
