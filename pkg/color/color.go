package color

import "fmt"

var CRed string = "\033[31m"
var CGreen string = "\033[32m"
var reset string = "\033[0m"

func Print(text string, color string) {
	fmt.Printf("%s%s%s", color, text, reset)
}
