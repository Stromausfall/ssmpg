package markdown

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/stromausfall/ssmpg/input"
)

func CreateContent(configData input.ConfigData, contentData input.ContentData) string {
	result :=
		fmt.Sprintf(
			"<p>%d-%02d-%02d %02d:%02d</p></br>\n",
			contentData.Date.Year(),
			contentData.Date.Month(),
			contentData.Date.Day(),
			contentData.Date.Hour(),
			contentData.Date.Minute())
	result += "<" + configData.TitleType + ">" + contentData.Title + "</" + configData.TitleType + ">\n"
	result += Convert(contentData.Content)

	return result
}

func Convert(input string) string {
	inputBytes := []byte(input)
	unsafeBytes := blackfriday.MarkdownCommon(inputBytes)
	htmlBytes := bluemonday.UGCPolicy().SanitizeBytes(unsafeBytes)

	return string(htmlBytes)
}
