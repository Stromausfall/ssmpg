package token

import (
	"strings"
)

const (
		RawHtml = iota
        RawText = iota
        Bold = iota
        Cursive = iota
        BoldCursive = iota
        Newline = iota
        
        cursiveString_1 = "*"
        cursiveString_2 = "_"
        
        boldString_1 = "**"
        boldString_2 = "__"
        
        boldCursiveString_1 = "***"
        boldCursiveString_2 = "___"
        
        
        
        newlineString = "\n"
)

type Token struct {
	TokenType int
	Value string
}

func sendToken(input string, tokenId int, output chan Token) {
	// we successfully built this token !
	output <- Token { tokenId, input }
}

func matchToken(rawInput *string, tokenString string, tokenId int, output chan Token) bool {
	input := *rawInput
	
	if strings.HasPrefix(input, tokenString) {
		// remove the token from the rawInput !
		*rawInput = input[len(tokenString):]
		
		// send the token
		sendToken(tokenString, tokenId, output)
		
		return true
	}
	
	return false
}

func mergeRawTextTokens(input chan Token, output chan Token) {
	mergedRawText := ""
	
	for element := range input {
		if element.TokenType == RawText {
			mergedRawText += element.Value
		} else {
			if len(mergedRawText) != 0 {
				// we have some merged raw text !
				sendToken(mergedRawText, RawText, output)
				
				// reset
				mergedRawText = ""
			}
			
			output <- element
		}
	}
	
	// maybe there is still some mergedRawText to emit !
	if len(mergedRawText) != 0 {
		// we have some merged raw text !
		sendToken(mergedRawText, RawText, output)
	}
	
	close(output)
}

func matchBlockHtml(rawInput *string, output chan Token) bool {
	// first check whether it looks like a html block statement
	if !strings.HasPrefix(*rawInput, "\n\n<") {
		return false
	}
	
	// find the end of the html block statement
	endIndexOfBlockStatement := strings.Index(*rawInput, ">\n\n")
	
	if endIndexOfBlockStatement == -1 {
		// not found - no block statement
		return false
	}
	
	// if none matched - let's emit the first char
	toSend := (*rawInput)[:endIndexOfBlockStatement+3]
	*rawInput = (*rawInput)[endIndexOfBlockStatement+3:]
	
	sendToken(toSend, RawText, output)

	return true
}

func Tokenizer(input string, output chan Token) {
	unmergedOutput := make(chan Token)
	
	// start the goroutine that cleans the 
	// unmergedOutput and forwards it to the output !
	go mergeRawTextTokens(unmergedOutput, output)
	
	for {
		if len(input) == 0 {
			break
		}
		
		switch {
			case matchBlockHtml(&input, unmergedOutput):
			
			case matchToken(&input, boldString_1, Bold, unmergedOutput):
			case matchToken(&input, boldString_2, Bold, unmergedOutput):
			
			case matchToken(&input, cursiveString_1, Cursive, unmergedOutput):
			case matchToken(&input, cursiveString_2, Cursive, unmergedOutput):
			
			case matchToken(&input, boldCursiveString_1, BoldCursive, unmergedOutput):
			case matchToken(&input, boldCursiveString_2, BoldCursive, unmergedOutput):
			
			case matchToken(&input, newlineString, Newline, unmergedOutput):
			case true:
				// if none matched - let's emit the first char
				toSend := input[:1]
				input = input[1:]
				
				sendToken(toSend, RawText, unmergedOutput)
		}
	}
	
	close(unmergedOutput)
}
