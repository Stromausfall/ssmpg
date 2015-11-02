package input

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

const correctDateFormat = "2006-01-02 15:4"

// the content loaded from a file
type ContentData struct {
	Title      string
	Date       time.Time
	Categories []string
	Content    string
}

// the struct to which the yaml config data is converted
type contentConfigData struct {
	Title      string
	Date       string
	Categories []string
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

func getContentFrom(content string) string {
	content = strings.Trim(content, " \r\n\t")

	if content == "" {
		fmt.Println("only description no content in the content file !")
		panic("Content missing in content file")
	}

	return content
}

func loadContentRawData(path string) (yamlData, rawContent string) {
	data, _ := ioutil.ReadFile(path)
	dataString := string(data)
	splitIndex := strings.Index(dataString, "---")
	
	if splitIndex == -1 {
		fmt.Println("Problem with file : " + path)
		panic("Content file must contain '---' to separate config data from content")
	}
	
	createdRawContent := dataString[splitIndex+len("---"):]
	createdYamlData := dataString[:splitIndex]

	return createdYamlData, createdRawContent
}

func createStructFromYamlString(yamlData string) contentConfigData {
	structFromFile := contentConfigData{}
	yaml.Unmarshal([]byte(yamlData), &structFromFile)

	return structFromFile
}

func loadDateFromStructFromFile(structFromFile contentConfigData, path string) time.Time {
	if structFromFile.Date == "" {
		fmt.Println("problem with file : " + path)
		panic("Date missing in content file")
	}

	convertedDate, dateConversionError := time.Parse(correctDateFormat, structFromFile.Date)

	if dateConversionError != nil {
		fmt.Println("problem with file : " + path)
		fmt.Println("Incorrect date : " + structFromFile.Date + " should be in format : " + correctDateFormat)
		panic("Incorrect format of date in content file")
	}

	return convertedDate
}

func loadDateFromStructFromTitle(structFromFile contentConfigData, path string) string {
	if structFromFile.Title == "" {
		fmt.Println("problem with file : " + path)
		panic("Title missing in content file")
	}

	return structFromFile.Title
}

func createContentData(path string) ContentData {
	yamlData, rawContent := loadContentRawData(path)
	structFromFile := createStructFromYamlString(yamlData)

	result := ContentData{}
	result.Content = getContentFrom(rawContent)
	result.Date = loadDateFromStructFromFile(structFromFile, path)
	result.Title = loadDateFromStructFromTitle(structFromFile, path)
	result.Categories = structFromFile.Categories

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
