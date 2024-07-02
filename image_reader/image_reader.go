package imagereader

import (
	"image"
	_ "image"
	_ "image/png"
	_ "image/jpeg"
	"os"
)

func LoadFromFile(filename string) (image.Image, error){
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, _, err := image.Decode(f)
	return image, err
}
