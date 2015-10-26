package parser

import (
	"token"
)

type parseElement interface {
	setContent(content []parseElement)
	getContent() string
}


type enclosingParser struct {
	start string
	end string
	content []parseElement
}

func (b *enclosingParser) getContent() string {
	return b.start + merge(b.content) + b.end
}

func (b *enclosingParser) setContent(content []parseElement) {
	b.content = content
}

type rawTextParser struct {
	content string
}

func (b *rawTextParser) getContent() string {
	return b.content
}

func (b *rawTextParser) setContent(content []parseElement) {
// maybe throw an exception here !
}





func (b *rawTextParser) setStringContent(content string) {
	b.content = content
}


// returns the index of the first token of the given type
// starting from the startIndex inside the elements
// -1 is returned if none was found
func findFirstTokenIndex(elements []token.Token, tokenType int, startIndex int) int { 
	for i := startIndex; i < len(elements); i++ {
		element := elements[i]
		
		if element.TokenType == tokenType {
			return i
		} 
	}
	
	return -1
}

// parses for an enclosing element, for example with startTokenType * and endTokenType +
//  --> * fo foo fo +
// and then returns whether it was found and the cocntent of the enclosing element
// in the example provided before the content is ' fo foo fo '
func getContentForEnclosing(elements *[]token.Token,  startTokenType int, endTokenType int) (bool, []token.Token) {
	if len(*elements) == 0 {
		// nothing to parse
		return false, nil
	}
	
	enclosingStartTokenIndex := findFirstTokenIndex(*elements, startTokenType, 0)
	enclosingEndTokenIndex := findFirstTokenIndex(*elements, endTokenType, 1)
	
	if enclosingStartTokenIndex != 0 {
		// the first token was not the desired token
		return false, nil
	}
	
	if enclosingEndTokenIndex == -1 {
		// if there was no end token index !
		return false, nil
	}
	
	// get the content of the enclosing type
	content := (*elements)[1:enclosingEndTokenIndex]
	// start from (enclosingEndTokenIndex + 1) because we don't want to keep the enclosing token !
	*elements = (*elements)[enclosingEndTokenIndex + 1:]
	
	return true, content
}

// fills the parseElement implementation with the content if
// the enclosing type exists (then returns true otherwise false)
func parseForEnclosing(elements *[]token.Token, startTokenType int, endTokenType int, toCreate parseElement, target *[]parseElement) bool {
	exists, content := getContentForEnclosing(elements, startTokenType, endTokenType)
	
	if !exists {
		// not found
		return false
	}

	parseElementContent := make([]parseElement, 0)
	
	// parse the content and attach it to the content of the ParseElement !
	parser2(&content, &parseElementContent)
	
	toCreate.setContent(parseElementContent)
	
	// add the created parseElement to the target slice !
	*target = append(*target, toCreate)
	
	return true
}

func parseForRawText(elements *[]token.Token, target *[]parseElement) {
	if len(*elements) == 0 {
		return 
	}
	
	content := (*elements)[0].Value
	parseElement := rawTextParser { content }
	
	*target = append(*target, &parseElement)
	
	*elements = (*elements)[1:]
}

func parseForSingle(elements *[]token.Token, tokenType int, toCreate parseElement, target *[]parseElement) bool {
	content := (*elements)[0]
	
	if content.TokenType != tokenType {
		// requested token not found
		return false
	}
	
	*target = append(*target, toCreate)
	
	*elements = (*elements)[1:]
	
	return true
}

func parser2(elements *[]token.Token, target *[]parseElement) {
	for {
		if len(*elements) == 0 {
			break
		}
		
		switch {
			case parseForSingle(elements, token.Newline, &enclosingParser { "</br>\n", "", nil }, target):
			case parseForEnclosing(elements, token.Bold, token.Bold, &enclosingParser { "<b>", "</b>", nil }, target):
			case parseForEnclosing(elements, token.BoldCursive, token.BoldCursive, &enclosingParser { "<b><i>", "</i></b>", nil }, target):
			case parseForEnclosing(elements, token.Cursive, token.Cursive, &enclosingParser { "<i>", "</i>", nil }, target):
			case true:
				parseForRawText(elements, target)
		}
	}
}

func merge(elements []parseElement) string {
	result := "" 

	for _, element := range elements {
		result += element.getContent()
	}
	
	return result
}

func Parser(input chan token.Token, output chan string) {
	elements := make([]token.Token, 0)
	
	for element := range input {
		elements = append(elements, element)
	}
	
	rootParseElementContent := make([]parseElement, 0)
	parser2(&elements, &rootParseElementContent)
	
	output <- merge(rootParseElementContent)
	close(output)
}