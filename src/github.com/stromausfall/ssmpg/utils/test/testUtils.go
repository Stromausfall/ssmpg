package test

import (
	"os"
	"fmt"
    "reflect"
	"testing"
	"io/ioutil"
	"path/filepath"
)

func TestExpectException(expectedPanic string, t *testing.T) {
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
}

func CreateTestFileReturnPath(filename, content string) string {
	path := filepath.Join(GetTestPath(), filename)
	CreateTestFile(path, content)	
	
	return path
}

func CreateTestDirectory(directoryName string) {
	err := os.Mkdir(directoryName, 0777)
	
	if err != nil {
		fmt.Printf(fmt.Sprint("%v", err))
		panic("unable to create directory in path : " + directoryName)
	}
}

func CreateTestFile(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	
	if err != nil {
		fmt.Printf(fmt.Sprint("%v", err))
		panic("unable to create file in path : " + filename)
	}
}

func GetTestPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	
	if err != nil {
		fmt.Printf(fmt.Sprint("%v", err))
		panic("error while getting testPath !")
	}
	
	return dir
}

func TestValues(t *testing.T, expected, actual, message string) {
	if expected != actual {
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("actual   : '" + actual + "'")
		t.Error(message)
	}
}

func GetTestFileContent(path string) string {
	result, readFileError := ioutil.ReadFile(path)
	
	if readFileError != nil {
		panic("unable to load file : " + path)
	}
	
	return string(result)
}
