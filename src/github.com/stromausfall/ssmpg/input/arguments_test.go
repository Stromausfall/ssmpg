package input

import (
	"github.com/stromausfall/ssmpg/utils/test"
	"testing"
)

func testAndExpectPanic(configFile, baseHtml, contentFolder, outputFolder, expectedPanic string, t *testing.T) {
	args := []string{configFile, baseHtml, contentFolder, outputFolder}

	testAndExpectPanicWithArgs(args, expectedPanic, t)
}

func testAndExpectPanicWithArgs(args []string, expectedPanic string, t *testing.T) {
	defer test.TestExpectException(expectedPanic, t)

	BasicValidationOfConsoleArguments(args)
}

func TestArgumentNull(t *testing.T) {
	testAndExpectPanicWithArgs(nil, "argument error - no args", t)
}

func TestIncorrectArgumentCount(t *testing.T) {
	testAndExpectPanicWithArgs([]string{}, "argument error - incorrect argument count", t)
	testAndExpectPanicWithArgs([]string{""}, "argument error - incorrect argument count", t)
	testAndExpectPanicWithArgs([]string{"", ""}, "argument error - incorrect argument count", t)
	testAndExpectPanicWithArgs([]string{"", "", ""}, "argument error - incorrect argument count", t)
	testAndExpectPanicWithArgs([]string{"", "", "", "", ""}, "argument error - incorrect argument count", t)
}

func TestCheckBasicPathForConfigFile(t *testing.T) {
	testAndExpectPanic(string(0), "asdf", "asdf", "asdf", "argument error - problem with the path of the configFile", t)
}

func TestCheckBasicPathForBaseHtml(t *testing.T) {
	testAndExpectPanic("asdf", string(0), "asdf", "asdf", "argument error - problem with the path of the baseHtml", t)
}

func TestCheckBasicPathContentFolder(t *testing.T) {
	testAndExpectPanic("asdf", "asdf", string(0), "asdf", "argument error - problem with the path of the contentFolder", t)
}

func TestCheckFileExistsForConfigFile(t *testing.T) {
	testAndExpectPanic("asdf", "asdf", "asdf", "asdf", "argument error - file does not exist", t)
}

func TestCheckFileHasCorrectEndingForConfigFile(t *testing.T) {
	filename := test.CreateTestFileReturnPath("fileWithWrongEnding.txt", "")
	testAndExpectPanic(filename, "asdf", "asdf", "asdf", "argument error - file has incorrect file ending", t)
}

func TestCheckFileExistsForBaseHtml(t *testing.T) {
	filename := test.CreateTestFileReturnPath("validEnding.yaml", "")
	testAndExpectPanic(filename, "asdf", "asdf", "asdf", "argument error - file does not exist", t)
}

func TestCheckFileHasCorrectEndingForBaseHtml(t *testing.T) {
	filename := test.CreateTestFileReturnPath("validEnding.yaml", "")
	filename2 := test.CreateTestFileReturnPath("fileWithWrongEnding.txt", "")
	testAndExpectPanic(filename, filename2, "asdf", "asdf", "argument error - file has incorrect file ending", t)
}

func TestCheckDirectoryExistsForContentFolder(t *testing.T) {
	filename := test.CreateTestFileReturnPath("validEnding.yaml", "")
	filename2 := test.CreateTestFileReturnPath("validEnding.html", "")
	testAndExpectPanic(filename, filename2, "asdf", "asdf", "argument error - directory does not exist", t)
}

func TestCheckFilePathForContentFolderIsADirectory(t *testing.T) {
	filename := test.CreateTestFileReturnPath("validEnding.yaml", "")
	filename2 := test.CreateTestFileReturnPath("validEnding.html", "")
	testAndExpectPanic(filename, filename2, filename2, "asdf", "argument error - is not a directory", t)
}

func TestCheckFilePathForOutputFolderIsADirectory(t *testing.T) {
	filename := test.CreateTestFileReturnPath("validEnding.yaml", "")
	filename2 := test.CreateTestFileReturnPath("validEnding.html", "")
	workingDir := test.GetTestPath()
	testAndExpectPanic(filename, filename2, workingDir, filename, "argument error - is not a directory", t)
}

func TestCheckFilePathForOutputFolderExists(t *testing.T) {
	filename := test.CreateTestFileReturnPath("validEnding.yaml", "")
	filename2 := test.CreateTestFileReturnPath("validEnding.html", "")
	workingDir := test.GetTestPath()
	testAndExpectPanic(filename, filename2, workingDir, filename + "sdf", "argument error - directory does not exist", t)
}
