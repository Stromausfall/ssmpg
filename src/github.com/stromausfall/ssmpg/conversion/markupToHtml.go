package conversion

import (
    "github.com/microcosm-cc/bluemonday"
    "github.com/russross/blackfriday"
)

func Convert(input string) string {
	inputBytes := []byte(input)
	unsafeBytes := blackfriday.MarkdownCommon(inputBytes)
	htmlBytes := bluemonday.UGCPolicy().SanitizeBytes(unsafeBytes)
	
	return string(htmlBytes)
}
