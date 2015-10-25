package token

import (
	"strings"
)

const (
        RawText = iota
        Bold = iota
        Cursive = iota
        Newline = iota
        
        rawTextString = ""
        boldString = "**"
        cursiveString = "*"
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
		case matchToken(&input, boldString, Bold, unmergedOutput):
		case matchToken(&input, cursiveString, Cursive, unmergedOutput):
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
