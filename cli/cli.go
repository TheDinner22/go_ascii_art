package cli

import (
	"errors"
	"flag"
)

type percent uint8

func toPercent(num uint8) (percent, error){
    if num > 100 {
        return 0, errors.New("percents cannot be greater than 100")
    }
    return percent(num), nil
}


// where to output
// % width
// % height
type args struct {

}

func init(){
    nFlag := flag.Int("width", 100, "Value from 0-100 inclusive. Width of the image")
}
