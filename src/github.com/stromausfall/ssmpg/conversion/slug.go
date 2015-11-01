package conversion

import (
	"strings"
	"unicode"
)

func punctuationToWhitespace(element rune) rune {
	if unicode.IsLetter(element) || unicode.IsDigit(element) {
		return unicode.ToLower(element)
	} else {
		return ' '
	}
}

func changePunctuationToWhitespace(toConvert string) string {
	return strings.Map(punctuationToWhitespace, toConvert)
}

func cleanWhitespaces(toConvert string) string {
	data := toConvert
	oldData := ""

	// repeat until there's no more change
	// -> we want to collapse all whitespaces
	for data != oldData {
		oldData = data
		data = strings.Replace(data, "  ", " ", -1)
	}

	// remove trailing and leading whitespaces
	data = strings.Trim(data, " ")

	return data

}

func createSlugWhiteSpaceToHyphen(toConvert string) string {
	// convert whitespaces
	return strings.Replace(toConvert, " ", "-", -1)
}

func GenerateSlug(toConvert string) string {
	if toConvert == "" {
		panic("no argument passed to convert")
	}

	toConvertWithWhitespaces := changePunctuationToWhitespace(toConvert)
	cleanedToConvertWithWhitespaces := cleanWhitespaces(toConvertWithWhitespaces)
	result := createSlugWhiteSpaceToHyphen(cleanedToConvertWithWhitespaces)

	return result
}
