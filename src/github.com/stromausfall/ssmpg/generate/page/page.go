package page

import (
	"fmt"
	"strings"
	"os"
    "bufio"
	"github.com/stromausfall/ssmpg/generate/page/content/category"
	"github.com/stromausfall/ssmpg/generate/page/content/markdown"
	"github.com/stromausfall/ssmpg/generate/sidebar"
	"github.com/stromausfall/ssmpg/generate/slug"
	"github.com/stromausfall/ssmpg/input"
)

func createPage(
	title string,
	outputFile string,
	configData input.ConfigData,
	elements []input.ContentData,
	isIndexPage bool,
	stringContent string,
	baseHtml string) {
	topBarContent := markdown.Convert(configData.Topbar)
	bottomBarContent := markdown.Convert(configData.Bottombar)
	sideBarContent := sidebar.GenerateSidebar(configData, elements, isIndexPage)
	
    data := string(baseHtml)
	
    data = strings.Replace(data, "XXX-TITLE-XXX", title, 1)
    data = strings.Replace(data, "XXX-TOPBAR-XXX", topBarContent, 1)
    data = strings.Replace(data, "XXX-BODY-XXX", stringContent, 1)
    data = strings.Replace(data, "XXX-INDEX-XXX", sideBarContent, 1)
    data = strings.Replace(data, "XXX-BOTTOMBAR-XXX", bottomBarContent, 1)
	
	file, fileCreateErr := os.Create(outputFile)
	defer file.Close()
	
	if fileCreateErr != nil {
		fmt.Println("%v", fileCreateErr)
		panic("unable to create file at " + outputFile)
	}
	
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	defer file.Sync()
	
	writer.WriteString(data)
	
	defer fmt.Println("Created file : " + outputFile)
}

func CreateContentPage(outputDirectory string, configData input.ConfigData, content input.ContentData, elements []input.ContentData, baseHtml string) {
	stringContent := markdown.CreateContent(configData, content)
	outputFile := outputDirectory + "/content/" + slug.GenerateSlug(content.Title) + ".html"

	createPage(content.Title, outputFile, configData, elements, false, stringContent, baseHtml)
}

func CreateCategoryPage(outputDirectory string, configData input.ConfigData, categoryToDisplay string, elements []input.ContentData, baseHtml string) {
	stringContent := category.CreateCategoryPage(configData, categoryToDisplay, elements)
	outputFile := outputDirectory + "/categories/" + slug.GenerateSlug(categoryToDisplay) + ".html"

	createPage(categoryToDisplay, outputFile, configData, elements, false, stringContent, baseHtml)
}

func CreateIndexPage(outputDirectory string, configData input.ConfigData, elements []input.ContentData, baseHtml string) {
	stringContent := category.CreateIndex(configData, elements)
	outputFile := outputDirectory + "/index.html"

	createPage(configData.IndexName, outputFile, configData, elements, true, stringContent, baseHtml)
}
