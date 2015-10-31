package input

import (
	"github.com/stromausfall/ssmpg/utils"
	"testing"
)

func testLoadBaseHtmlDataExpectException(expectedPanic, argument string, t *testing.T) {
	defer utils.TestExpectException(expectedPanic, t)

	LoadBaseHtmlData(argument)
}

func TestCreateBaseHtmlDataWrongFile(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.html", "")

	testLoadBaseHtmlDataExpectException("error while opening baseHtmlData", path+"234", t)
}

func TestCreateBaseHtmlData(t *testing.T) {
	validContent := "1234 XXX-TOPBAR-XXX XXX-BODY-XXX XXX-INDEX-XXX XXX-BOTTOMBAR-XXX 1234"
	path := utils.CreateTestFileReturnPath("foo.html", validContent)
	loadedData := LoadBaseHtmlData(path)

	if validContent != loadedData {
		t.Error("content of baseHtml was not correctly loaded !")
	}
}

func TestCreateBaseHtmlDataMissingPlaceHolder1(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.html", validConfigFileContent+"XXX-BODY-XXX XXX-INDEX-XXX XXX-BOTTOMBAR-XXX")

	testLoadBaseHtmlDataExpectException("the baseHtml didn't contain all required placeholders - missing XXX-TOPBAR-XXX", path, t)
}

func TestCreateBaseHtmlDataMissingPlaceHolder2(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.html", validConfigFileContent+"XXX-TOPBAR-XXX XXX-INDEX-XXX XXX-BOTTOMBAR-XXX")

	testLoadBaseHtmlDataExpectException("the baseHtml didn't contain all required placeholders - missing XXX-BODY-XXX", path, t)
}

func TestCreateBaseHtmlDataMissingPlaceHolder3(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.html", validConfigFileContent+"XXX-TOPBAR-XXX XXX-BODY-XXX XXX-BOTTOMBAR-XXX")

	testLoadBaseHtmlDataExpectException("the baseHtml didn't contain all required placeholders - missing XXX-INDEX-XXX", path, t)
}

func TestCreateBaseHtmlDataMissingPlaceHolder4(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.html", validConfigFileContent+"XXX-TOPBAR-XXX XXX-BODY-XXX XXX-INDEX-XXX")

	testLoadBaseHtmlDataExpectException("the baseHtml didn't contain all required placeholders - missing XXX-BOTTOMBAR-XXX", path, t)
}
