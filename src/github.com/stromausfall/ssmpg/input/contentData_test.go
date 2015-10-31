package input

import (
	"fmt"
	"github.com/stromausfall/ssmpg/utils"
	"path/filepath"
	"strings"
	"testing"
)

const correctContent = "title: Old Pages\ndate: 2015-03-24 22:28\ncategories:\n- X1\n- X2\n---\nOld unsupported pages :"

func testLoadContentDataExpectException(expectedPanic, argument string, t *testing.T) {
	defer utils.TestExpectException(expectedPanic, t)

	LoadContentData(argument)
}

func TestLoadContentDataWrongPath(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.html", correctContent)

	testLoadContentDataExpectException("error path has to point to a directory", path, t)
}

func TestLoadContentCorrectContentCount(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctContent)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctContent)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo2.md"), correctContent)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo3.md"), correctContent)

	files := LoadContentData(pathToFolder)

	if len(files) != 3 {
		t.Error("expected 3 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentCountWithSubfolders(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctContent)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo2")
	pathToFolder2 := filepath.Join(pathToFolder, "foo2")
	pathToFolder3 := filepath.Join(pathToFolder2, "foo2")
	pathToFolder4 := filepath.Join(pathToFolder3, "foo2")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestDirectory(pathToFolder2)
	utils.CreateTestDirectory(pathToFolder3)
	utils.CreateTestDirectory(pathToFolder4)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctContent)
	utils.CreateTestFile(filepath.Join(pathToFolder2, "foo2.md"), correctContent)
	utils.CreateTestFile(filepath.Join(pathToFolder4, "foo3.md"), correctContent)

	files := LoadContentData(pathToFolder)

	if len(files) != 3 {
		t.Error("expected 3 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentCountAccordingToFileEnding(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctContent)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo3")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctContent)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo2.markdown"), correctContent)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo3.m2"), correctContent)

	files := LoadContentData(pathToFolder)

	if len(files) != 2 {
		t.Error("expected 2 files but were : " + fmt.Sprintf("%v", len(files)))
	}
}

func TestLoadContentCorrectContentForData(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.md", correctContent)
	dirPath := strings.TrimSuffix(path, "foo.md")
	pathToFolder := filepath.Join(dirPath, "foo4")

	utils.CreateTestDirectory(pathToFolder)
	utils.CreateTestFile(filepath.Join(pathToFolder, "foo1.md"), correctContent)

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
	if len(file.Categories) != 2 {
		t.Error("category not correctly loaded")
	}
	if file.Categories[0] != "X1" || file.Categories[1] != "X2" {
		t.Error("category content not correctly loaded")
	}
	if file.Content != "Old unsupported pages :" {
		t.Error("content not correctly loaded : '" + file.Content + "'")
	}
}
