package input

import (
	"testing"
	"github.com/stromausfall/ssmpg/utils"
)

func testAndExpectPanic(configFile, baseHtml, contentFolder, expectedPanic string, t *testing.T) {
	args := []string{ configFile, baseHtml, contentFolder }
	
	testAndExpectPanicWithArgs(args, expectedPanic, t)
}    
    
func testAndExpectPanicWithArgs(args []string, expectedPanic string, t *testing.T) {
	utils.ExpectException(
		func() {
			BasicValidationOfConsoleArguments(args)
		},
		expectedPanic,
		t)
}

func TestArgumentNull(t *testing.T) {
	testAndExpectPanicWithArgs(nil, "argument error - no args", t)
}

func TestIncorrectArgumentCount(t *testing.T) {
	testAndExpectPanicWithArgs([]string{ }, "argument error - incorrect argument count", t)
	testAndExpectPanicWithArgs([]string{ "" }, "argument error - incorrect argument count", t)
	testAndExpectPanicWithArgs([]string{ "", "" }, "argument error - incorrect argument count", t)
	testAndExpectPanicWithArgs([]string{ "", "", "", "" }, "argument error - incorrect argument count", t)
}

func TestCheckBasicPathForConfigFile(t *testing.T) {
	testAndExpectPanic(string(0), "asdf", "asdf", "argument error - problem with the path of the configFile", t)
}

func TestCheckBasicPathForBaseHtml(t *testing.T) {
	testAndExpectPanic("asdf", string(0), "asdf", "argument error - problem with the path of the baseHtml", t)
}

func TestCheckBasicPathContentFolder(t *testing.T) {
	testAndExpectPanic("asdf", "asdf", string(0), "argument error - problem with the path of the contentFolder", t)
}

func TestCheckFileExistsForConfigFile(t *testing.T) {
	testAndExpectPanic("asdf", "asdf", "asdf", "argument error - file does not exist", t)
}

func TestCheckFileHasCorrectEndingForConfigFile(t *testing.T) {
	filename := utils.CreateTestFile("fileWithWrongEnding.txt", "")
	testAndExpectPanic(filename, "asdf", "asdf", "argument error - file has incorrect file ending", t)
}

func TestCheckFileExistsForBaseHtml(t *testing.T) {
	filename := utils.CreateTestFile("validEnding.json", "")
	testAndExpectPanic(filename, "asdf", "asdf", "argument error - file does not exist", t)
}

func TestCheckFileHasCorrectEndingForBaseHtml(t *testing.T) {
	filename := utils.CreateTestFile("validEnding.json", "")
	filename2 := utils.CreateTestFile("fileWithWrongEnding.txt", "")
	testAndExpectPanic(filename, filename2, "asdf", "argument error - file has incorrect file ending", t)
}

func TestCheckDirectoryExistsForContentFolder(t *testing.T) {
	filename := utils.CreateTestFile("validEnding.json", "")
	filename2 := utils.CreateTestFile("validEnding.html", "")
	testAndExpectPanic(filename, filename2, "asdf", "argument error - directory does not exist", t)
}

func TestCheckFilePathForContentFolderIsADirectory(t *testing.T) {
	filename := utils.CreateTestFile("validEnding.json", "")
	filename2 := utils.CreateTestFile("validEnding.html", "")
	testAndExpectPanic(filename, filename2, filename2, "argument error - is not a directory", t)
}
