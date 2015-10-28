package arguments

import (
	"os"
	"fmt"
    "strings"
    "path/filepath"
)



func checkArguments(args []string) []string {
	if args == nil {
		panic("argument error - no args")
	}
	if len(args) != 3 {
		fmt.Println("program needs exactly three arguments (" + fmt.Sprint(len(args)) + " were provided) : configFile baseHtml contentFolder")
		fmt.Println(" configFile - path to the *.json config file")
		fmt.Println(" baseHtml - path to the *.html file which is used as base for the homepage")
		fmt.Println(" contentFolder - path to the folder that contains the content files *.md") 
		
		panic("argument error - incorrect argument count")
	}
	
	return args
}

func getAbsPath(path, name string) string {
	absolutePath, err := filepath.Abs(path)

	if err != nil {
		fmt.Println("Unable to get absolute path of the " + name + " + argument !");
		panic("argument error - problem with the path of the " + name);
	}
	
	return absolutePath
}

func checkFile(path, fileEnding string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file " + path + " does not exist !")
		} else {
			fmt.Println("problem with file " + path + " !")
		}
		
		panic("argument error - file does not exist")
    }
	
	if !strings.HasSuffix(path, fileEnding) {
		fmt.Println("file " + path + " should have file ending : " + fileEnding)
		panic("argument error - file has incorrect file ending")
	}
}

func checkDir(path string) {
    fileInfo, err := os.Stat(path)
	
    if err != nil {
    	fmt.Println("problem with directory path : " + path)
    	panic("argument error - directory does not exist")
    }
    
    if !fileInfo.IsDir() {
    	fmt.Println("path : " + path + " is not a directory !")
    	panic("argument error - is not a directory")
    }
}

func BasicValidationOfConsoleArguments(args []string) (configFile, baseHtml, contentFolder string) {
	args = checkArguments(args)
	
	configFile = args[0]
	baseHtml = args[1]
	contentFolder = args[2]
	
	configFile = getAbsPath(configFile, "configFile")
	baseHtml = getAbsPath(args[1], "baseHtml")
	contentFolder = getAbsPath(args[2], "contentFolder")
	checkFile(configFile, ".json")
	checkFile(baseHtml, ".html")
	checkDir(contentFolder)
	
	return
}
