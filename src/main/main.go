package main

import (
	"fmt"
	"app"
)

func main() {
	dummyInput := ""
	dummyInput += "Text attributes *italic*, **bold**, `monospace`, ~~strikethrough~~\n"
	result := app.GetTestHTML(dummyInput)
	
	fmt.Println(result)
}
