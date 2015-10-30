package utils

import (
	"os"
	"fmt"
    "reflect"
	"testing"
	"io/ioutil"
	"path/filepath"
)
    
func ExpectException(testFunction func(), expectedPanic string, t *testing.T) {
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
	
	testFunction()
}

func CreateTestFile(filename, content string) string {
	path := GetTestPath(filename)
	CreateFile(path, content)	
	
	return path
}

func CreateFile(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	
	if err != nil {
		fmt.Printf(fmt.Sprint("%v", err))
		panic("unable to create file in path : " + filename)
	}
}

func GetTestPath(filename string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	
	if err != nil {
		fmt.Printf(fmt.Sprint("%v", err))
		panic("error while getting testPath !")
	}
	
	return filepath.Join(dir, filename)
}
