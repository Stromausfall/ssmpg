package markdown

import (
	"fmt"
	"github.com/stromausfall/ssmpg/input"
	"github.com/stromausfall/ssmpg/utils/test"
	"path/filepath"
	"strings"
	"testing"
)

var testConfigData = input.ConfigData{TitleType: "h3"}

func TestCreateContent(t *testing.T) {
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo66")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), test.ContentData)

	files := input.LoadContentData(pathToFolder)
	result := CreateContent(testConfigData, files[0])
	expected := "<p>2015-03-24 22:28</p></br>\n<h3>Old Pages</h3>\n<p>Old unsupported pages :</p>\n"

	if result != expected {
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("result   : '" + result + "'")
		t.Error("incorrect created content")
	}
}

func TestCreateContentChangeVariousProperties(t *testing.T) {
	alteredCorrectData := strings.Replace(test.ContentData, "2015-03-24 22:28", "2015-03-24 22:29", -1)
	alteredCorrectData = strings.Replace(alteredCorrectData, "Old Pages", "New Pages", -1)
	alteredCorrectData = strings.Replace(alteredCorrectData, "Old unsupported pages :", "New supported pages :", -1)
	alteredTestConfigData := input.ConfigData{TitleType: "h2"}

	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo67")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredCorrectData)

	files := input.LoadContentData(pathToFolder)
	result := CreateContent(alteredTestConfigData, files[0])
	expected := "<p>2015-03-24 22:29</p></br>\n<h2>New Pages</h2>\n<p>New supported pages :</p>\n"

	if result != expected {
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("result   : '" + result + "'")
		t.Error("incorrect created content")
	}
}
