package main

import (
    imageReader "github.com/TheDinner22/image_reader"
    imageDisplayer "github.com/TheDinner22/terminal_image_displayer"
)

func displayer(){
    imageDisplayer.Test()
}

func reader(){
    imageReader.LoadFromFile("images/thing.png")
}

func main(){
    displayer()
    reader()
}
