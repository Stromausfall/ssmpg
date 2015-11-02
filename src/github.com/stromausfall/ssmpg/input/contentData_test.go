package input

import (
	"fmt"
	"github.com/stromausfall/ssmpg/utils/test"
	"path/filepath"
	"strings"
	"testing"
)

func testLoadContentDataExpectException(expectedPanic, argument string, t *testing.T) {
	defer test.TestExpectException(expectedPanic, t)

	LoadContentData(argument)
}

func TestLoadContentDataWrongPath(t *testing.T) {
	path := test.CreateTestFileReturnPath("foo.html", test.ContentData)

	testLoadContentDataExpectException("error path has to point to a directory", path, t)
}

func TestLoadContentCorrectContentCount(t *testing.T) {
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), test.ContentData)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo2.md"), test.ContentData)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo3.md"), test.ContentData)

	files := LoadContentData(pathToFolder)

	if len(files) != 3 {
		t.Error("expected 3 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentCountWithSubfolders(t *testing.T) {
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo2")
	pathToFolder2 := filepath.Join(pathToFolder, "foo2")
	pathToFolder3 := filepath.Join(pathToFolder2, "foo2")
	pathToFolder4 := filepath.Join(pathToFolder3, "foo2")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestDirectory(pathToFolder2)
	test.CreateTestDirectory(pathToFolder3)
	test.CreateTestDirectory(pathToFolder4)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), test.ContentData)
	test.CreateTestFile(filepath.Join(pathToFolder2, "foo2.md"), test.ContentData)
	test.CreateTestFile(filepath.Join(pathToFolder4, "foo3.md"), test.ContentData)

	files := LoadContentData(pathToFolder)

	if len(files) != 3 {
		t.Error("expected 3 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentCountAccordingToFileEnding(t *testing.T) {
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo3")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), test.ContentData)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo2.markdown"), test.ContentData)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo3.m2"), test.ContentData)

	files := LoadContentData(pathToFolder)

	if len(files) != 2 {
		t.Error("expected 2 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentForData(t *testing.T) {
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo4")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), test.ContentData)

	files := LoadContentData(pathToFolder)

	if len(files) != 1 {
		t.Error("expected 1 files but were : " + fmt.Sprintf("%v", len(files)))
	}

	file := files[0]

	test.TestValues(t, test.ContentDataTitle, file.Title, "title not correctly loaded")
	
	if file.Date.Unix() != 1427236080 {
		t.Error("date not correctly loaded")
	}
	if len(file.Categories) != len(test.ContentDataCategories) {
		t.Error("category not correctly loaded")
	}
	if file.Categories[0] != test.ContentDataCategories[0] || file.Categories[1] != test.ContentDataCategories[1] {
		t.Error("category content not correctly loaded")
	}
	
	test.TestValues(t, test.ContentDataContent, file.Content, "content not correctly loaded")
}

func TestLoadContentMissingTitleException(t *testing.T) {
	alteredContent := strings.Replace(test.ContentData, test.ContentDataTitle, "", -1)
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo5")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Title missing in content file", pathToFolder, t)
}

func TestLoadContentMissingDateException(t *testing.T) {
	alteredContent := strings.Replace(test.ContentData, test.ContentDataDate, "", -1)
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo6")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Date missing in content file", pathToFolder, t)
}

func TestLoadContentIncorrectDateException(t *testing.T) {
	alteredContent := strings.Replace(test.ContentData, test.ContentDataDate, "a"+test.ContentDataDate+"1", -1)
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo7")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Incorrect format of date in content file", pathToFolder, t)
}

func TestLoadContentMissingContentException(t *testing.T) {
	alteredContent := strings.Replace(test.ContentData, test.ContentDataContent, "", -1)
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo8")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Content missing in content file", pathToFolder, t)
}

func TestLoadContentMissingSeparator(t *testing.T) {
	alteredContent := strings.Replace(test.ContentData, "---", "", -1)
	dirPath := test.GetTestPath()
	pathToFolder := filepath.Join(dirPath, "foo9")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), alteredContent)

	testLoadContentDataExpectException("Content file must contain '---' to separate config data from content", pathToFolder, t)
}
