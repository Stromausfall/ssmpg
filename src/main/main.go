package main

import (
	"fmt"
	"token"
)

func print(input chan token.Token) {
	for element := range input {
		fmt.Println(fmt.Sprint(element.TokenType) + " - " + element.Value)
	}
}

func main() {
	dummyInput := "bla bla bla \nbla **bla** bl*a bla ** bla**\noi!" 
	
	tokens := make(chan token.Token)
	
	go token.Tokenizer(dummyInput, tokens)
	
	print(tokens)
}
