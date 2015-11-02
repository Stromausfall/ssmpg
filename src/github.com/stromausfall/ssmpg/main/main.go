package main

import (
	"os"
	"fmt"
	"sync"
	"path/filepath"
	"github.com/stromausfall/ssmpg/input"
	"github.com/stromausfall/ssmpg/generate/sidebar"
	"github.com/stromausfall/ssmpg/generate/page"
)

func createCategories(
	waitGroup *sync.WaitGroup,
	categories []string,
	outputPath string,
	configData input.ConfigData,
	contentData []input.ContentData,
	baseHtmlContent string) {
	fmt.Printf("number of categories : %v\n", len(categories))
	for _, category := range categories {
		// get and store the current category (needed for the following lambda)
		currentCategory := category
		
		waitGroup.Add(1)
		
		go func() {
			defer waitGroup.Done()
			page.CreateCategoryPage(outputPath, configData, currentCategory, contentData, baseHtmlContent)
		}()
	}
}

func createContent(
	waitGroup *sync.WaitGroup,
	outputPath string,
	configData input.ConfigData,
	contentData []input.ContentData,
	baseHtmlContent string) {
	fmt.Printf("number of content : %v\n", len(contentData))
	for _, content := range contentData {
		// get and store the current content (needed for the following lambda)
		currentContent := content
		
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			page.CreateContentPage(outputPath, configData, currentContent, contentData, baseHtmlContent)
		}()
	}
}

func main() {
	// validate arguments
	configFile, baseHtml, contentFolder, outputPath := input.BasicValidationOfConsoleArguments(os.Args[1:])
	
	//configFile := os.Args[1]
	//baseHtml := os.Args[2]
	//contentFolder := os.Args[3]
	/*
	configFile := "C:/xxx/input/test.yaml"
	baseHtml := "C:/xxx/input/base.html"
	contentFolder := "C:/xxx/input/content"
	baseFolder := "C:/xxx/output/"
	*/
	
	os.Mkdir(filepath.Join(outputPath, "content"), 777)
	os.Mkdir(filepath.Join(outputPath, "categories"), 777)
	
	configData := input.CreateConfigData(configFile)
	baseHtmlContent := input.LoadBaseHtmlData(baseHtml)
	contentData := input.LoadContentData(contentFolder)
		
	// create index
	page.CreateIndexPage(outputPath, configData, contentData, baseHtmlContent)
	
	// create category pages
	categories := sidebar.CollectCategories(contentData)
	
	// we need to synchronize the go routines
	waitGroup := sync.WaitGroup{}
	
	createCategories(&waitGroup, categories, outputPath, configData, contentData, baseHtmlContent)	
	createContent(&waitGroup, outputPath, configData, contentData, baseHtmlContent)
	
	// wait until all go routines have finished
    waitGroup.Wait()
}
