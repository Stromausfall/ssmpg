package input

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"strings"
	"time"
)

type ContentData struct {
	Title      string
	Date       time.Time
	Categories []string
	Content    string
}

type ContentConfigData struct {
	Title      string
	Date       string
	Categories []string `yaml:",flow"`
}

func hasValidFileEnding(path string) bool {
	if strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".markdown") {
		return true
	} else {
		return false
	}
}

func traverseForMatchingFilePaths(startDir string, filepaths *[]string) {
	files, err := ioutil.ReadDir(startDir)

	if err != nil {
		fmt.Println("content could not be read from path :  + string")
		panic("error path has to point to a directory")
	}

	for _, file := range files {
		path := filepath.Join(startDir, file.Name())

		if file.IsDir() {
			traverseForMatchingFilePaths(path, filepaths)
		} else {
			if hasValidFileEnding(path) {
				*filepaths = append(*filepaths, path)
			}
		}
	}
}

func createContentData(path string) ContentData {
	data, _ := ioutil.ReadFile(path)
	dataString := string(data)
	splitIndex := strings.Index(dataString, "---")
	yamlData := dataString[:splitIndex]
	content := dataString[splitIndex+len("---"):]
	
	result := ContentData{}
	result.Content = strings.Trim(content, " \r\n");
	
	xxx := ContentConfigData{}
	yaml.Unmarshal([]byte(yamlData), &xxx)
	
	result.Title = xxx.Title
	result.Categories = xxx.Categories
	result.Date, _ = time.Parse("2006-01-02 15:4", xxx.Date)

	return result
}

func LoadContentData(filepath string) []ContentData {
	result := make([]ContentData, 0)
	files := make([]string, 0)
	traverseForMatchingFilePaths(filepath, &files)

	for _, path := range files {
		result = append(result, createContentData(path))
	}

	return result
}
