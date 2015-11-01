package input

import (
	"fmt"
	"github.com/stromausfall/ssmpg/utils"
	"path/filepath"
	"strings"
	"testing"
)

const correctTitle = "Old Pages"
const correctDate = "2015-03-24 22:28"
const correctContent = "Old unsupported pages :"

var correctCategories = []string{"X1", "X2"}
var correctData = "title: " + correctTitle + "\ndate: " + correctDate + "\ncategories:\n- " + correctCategories[0] + "\n- " + correctCategories[1] + "\n---\n" + correctContent

func testLoadContentDataExpectException(expectedPanic, argument string, t *testing.T) {
	defer utils.TestExpectException(expectedPanic, t)

	LoadContentData(argument)
}

func TestLoadContentDataWrongPath(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.html", correctData)

	testLoadContentDataExpectException("error path has to point to a directory", path, t)
}

func TestLoadContentCorrectContentCount(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctData)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctData)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo2.md"), correctData)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo3.md"), correctData)

	files := LoadContentData(pathToFolder)

	if len(files) != 3 {
		t.Error("expected 3 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentCountWithSubfolders(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctData)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo2")
	pathToFolder2 := filepath.Join(pathToFolder, "foo2")
	pathToFolder3 := filepath.Join(pathToFolder2, "foo2")
	pathToFolder4 := filepath.Join(pathToFolder3, "foo2")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestDirectory(pathToFolder2)
	utils.CreateTestDirectory(pathToFolder3)
	utils.CreateTestDirectory(pathToFolder4)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctData)
	utils.CreateTestFile(filepath.Join(pathToFolder2, "foo2.md"), correctData)
	utils.CreateTestFile(filepath.Join(pathToFolder4, "foo3.md"), correctData)

	files := LoadContentData(pathToFolder)

	if len(files) != 3 {
		t.Error("expected 3 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentCountAccordingToFileEnding(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctData)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo3")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctData)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo2.markdown"), correctData)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo3.m2"), correctData)

	files := LoadContentData(pathToFolder)

	if len(files) != 2 {
		t.Error("expected 2 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentForData(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctData)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo4")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctData)

	files := LoadContentData(pathToFolder)

	if len(files) != 1 {
		t.Error("expected 1 files but were : " + fmt.Sprintf("%v", len(files)))
	}

	file := files[0]

	if file.Title != "Old Pages" {
		t.Error("title not correctly loaded : '" + file.Title + "'")
	}
	if file.Date.Unix() != 1427236080 {
		t.Error("date not correctly loaded")
	}
	if len(file.Categories) != len(correctCategories) {
		t.Error("category not correctly loaded")
	}
	if file.Categories[0] != correctCategories[0] || file.Categories[1] != correctCategories[1] {
		t.Error("category content not correctly loaded")
	}
	if file.Content != correctContent {
		t.Error("content not correctly loaded : '" + file.Content + "'")
	}
}

func TestLoadContentMissingTitleException(t *testing.T) {
	alteredContent := strings.Replace(correctData, correctTitle, "", -1)
	path := utils.CreateTestFileReturnPath("foo.md", "")
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo5")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Title missing in content file", pathToFolder, t)
}

func TestLoadContentMissingDateException(t *testing.T) {
	alteredContent := strings.Replace(correctData, correctDate, "", -1)
	path := utils.CreateTestFileReturnPath("foo.md", "")
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo6")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Date missing in content file", pathToFolder, t)
}

func TestLoadContentIncorrectDateException(t *testing.T) {
	alteredContent := strings.Replace(correctData, correctDate, "a"+correctDate+"1", -1)
	path := utils.CreateTestFileReturnPath("foo.md", "")
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo7")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Incorrect format of date in content file", pathToFolder, t)
}

func TestLoadContentMissingContentException(t *testing.T) {
	alteredContent := strings.Replace(correctData, correctContent, "", -1)
	path := utils.CreateTestFileReturnPath("foo.md", "")
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo8")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Content missing in content file", pathToFolder, t)
}
