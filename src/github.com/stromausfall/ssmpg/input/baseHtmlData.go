package input

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func checkForPlaceHolder(data, placeholder string) {
	if !strings.Contains(data, placeholder) {
		fmt.Println("missing placeholder in baseHtml : " + placeholder)
		panic("the baseHtml didn't contain all required placeholders - missing " + placeholder)
	}
}

func LoadBaseHtmlData(filepath string) string {
	data, readFileError := ioutil.ReadFile(filepath)

	if readFileError != nil {
		fmt.Printf(fmt.Sprintf("%v", readFileError))
		panic("error while opening baseHtmlData")
	}

	result := string(data)

	checkForPlaceHolder(result, "XXX-TOPBAR-XXX")
	checkForPlaceHolder(result, "XXX-BODY-XXX")
	checkForPlaceHolder(result, "XXX-INDEX-XXX")
	checkForPlaceHolder(result, "XXX-BOTTOMBAR-XXX")

	return result
}
