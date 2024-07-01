package imagedisplayer

import "fmt"

func PrintImageToScreen(asciiImage string) error{
    _, err := fmt.Printf(asciiImage)
    return err
}

func Test(){
    fmt.Println("displayer")
}
