package arguments

import (
	"fmt"
    "reflect"
	"testing"
	"os"
	"io/ioutil"
	"path/filepath"
)

func testAndExpectPanic(configFile, baseHtml, contentFolder, expectedPanic string, t *testing.T) {
	args := []string{ configFile, baseHtml, contentFolder }
	
	testAndExpectPanicWithArgs(args, expectedPanic, t)
}    
    
func testAndExpectPanicWithArgs(args []string, expectedPanic string, t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
        	if r == expectedPanic {
        		// as expected !
        		return
        	}
			
			switch r.(type) {
				case string:
					t.Error("incorrect panic message : " + fmt.Sprintf("%v", r))
                default:
					t.Error("incorrect panic type : " + fmt.Sprintf("%v", reflect.TypeOf(r)))
            }
			
			return
        }
        
        // either no panic or not the correct one !
        t.Error("expected panic did not happen")
    }()
    
    BasicValidationOfConsoleArguments(args)
}

func createFile(filename string) {
	err := ioutil.WriteFile(filename, []byte{}, 0644)
	
	if err != nil {
		fmt.Printf(fmt.Sprint("%v", err))
		panic("unable to create file in path : " + filename)
	}
}

func getTestPath(filename string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	
	if err != nil {
		fmt.Printf(fmt.Sprint("%v", err))
		panic("error while getting testPath !")
	}
	
	return filepath.Join(dir, filename)
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
	filename := getTestPath("fileWithWrongEnding.txt")
	createFile(filename)
	testAndExpectPanic(filename, "asdf", "asdf", "argument error - file has incorrect file ending", t)
}

func TestCheckFileExistsForBaseHtml(t *testing.T) {
	filename := getTestPath("validEnding.json")
	createFile(filename)
	testAndExpectPanic(filename, "asdf", "asdf", "argument error - file does not exist", t)
}

func TestCheckFileHasCorrectEndingForBaseHtml(t *testing.T) {
	filename := getTestPath("validEnding.json")
	filename2 := getTestPath("fileWithWrongEnding.txt")
	createFile(filename)
	createFile(filename2)
	testAndExpectPanic(filename, filename2, "asdf", "argument error - file has incorrect file ending", t)
}

func TestCheckDirectoryExistsForContentFolder(t *testing.T) {
	filename := getTestPath("validEnding.json")
	filename2 := getTestPath("validEnding.html")
	createFile(filename)
	createFile(filename2)
	testAndExpectPanic(filename, filename2, "asdf", "argument error - directory does not exist", t)
}

func TestCheckFilePathForContentFolderIsADirectory(t *testing.T) {
	filename := getTestPath("validEnding.json")
	filename2 := getTestPath("validEnding.html")
	createFile(filename)
	createFile(filename2)
	testAndExpectPanic(filename, filename2, filename2, "argument error - is not a directory", t)
}
