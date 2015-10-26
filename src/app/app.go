package app

import (
	"token"
	"parser"
)

func Process(input string) string {
	tokens := make(chan token.Token)
	htmlOutput := make(chan string)
	
	go token.Tokenizer(input, tokens)
	go parser.Parser(tokens, htmlOutput)
	
	return format(htmlOutput)
}

func format(input chan string) string {
	content := ""
	
	for element := range input {
		content += element
	}
	
	return content
}

func GetTestHTML(input string) string {
	return "<!DOCTYPE html>\n<html>\n<body>\n" + Process(input) + "\n</body>\n</html>"
}
