package page

import (
	"testing"
	"path/filepath"
	"github.com/stromausfall/ssmpg/utils/test"
	"github.com/stromausfall/ssmpg/input"
	"github.com/stromausfall/ssmpg/generate/slug"
)

func createConfigFile(baseDir string) string {
	configDataPath := filepath.Join(baseDir, "configFile.yaml")
	test.CreateTestFile(configDataPath, test.ConfigDataContent)
	
	return configDataPath
}

func createContent(baseDir string) string {
	pathToFolder := filepath.Join(baseDir, "input")

	test.CreateTestDirectory(pathToFolder)
	test.CreateTestFile(filepath.Join(pathToFolder, "foo.md"), test.ContentData)
	
	return pathToFolder
}

func setupTestEnvironment(id string) (string, input.ConfigData, []input.ContentData) {
	baseDir := filepath.Join(test.GetTestPath(), "pageTest" + id)
	outputDirectory := filepath.Join(baseDir, "output")
	test.CreateTestDirectory(baseDir)
	test.CreateTestDirectory(outputDirectory)
	
	configDataPath := createConfigFile(baseDir)
	contentPath := createContent(baseDir)
	
	configData := input.CreateConfigData(configDataPath)
	elements := input.LoadContentData(contentPath)
	
	return outputDirectory, configData, elements
}

func TestCreateIndex(t *testing.T) {
	outputDirectory, configData, elements := setupTestEnvironment("1")
	
	CreateIndexPage(outputDirectory, configData, elements, test.BaseHtmlContent)
	expected := "Index ||| <p>foo</p>\n ||| <h2>Index</h2>\n<a href=\"content/old-pages.html\">(2015-03-24) Old Pages</a></br>\n ||| <a href=\"index.html\">Index</a></br>\n</br>\n<a href=\"categories/x1.html\">X1</a></br>\n<a href=\"categories/x2.html\">X2</a> ||| <p>© foo</p>\n"
	
	filePath := filepath.Join(outputDirectory, "index.html")
	result := test.GetTestFileContent(filePath)
	
	test.TestValues(t, expected, result, "incorrect page generated")
}

func TestCreateCategory(t *testing.T) {
	outputDirectory, configData, elements := setupTestEnvironment("2")
	category := test.ContentDataCategories[0]
	categoryPath := filepath.Join(outputDirectory, "categories")
	test.CreateTestDirectory(categoryPath)
	
	CreateCategoryPage(outputDirectory, configData, category, elements, test.BaseHtmlContent)
	expected := "X1 ||| <p>foo</p>\n ||| <h2>X1</h2>\n<a href=\"../content/old-pages.html\">(2015-03-24) Old Pages</a></br>\n ||| <a href=\"../index.html\">Index</a></br>\n</br>\n<a href=\"../categories/x1.html\">X1</a></br>\n<a href=\"../categories/x2.html\">X2</a> ||| <p>© foo</p>\n"
	
	filePath := filepath.Join(categoryPath, slug.GenerateSlug(category) + ".html")
	result := test.GetTestFileContent(filePath)
	
	test.TestValues(t, expected, result, "incorrect page generated")
}

func TestCreateContent(t *testing.T) {
	outputDirectory, configData, elements := setupTestEnvironment("3")
	contentPath := filepath.Join(outputDirectory, "content")
	test.CreateTestDirectory(contentPath)
	content := elements[0]
	
	CreateContentPage(outputDirectory, configData, content, elements, test.BaseHtmlContent)
	expected := "Old Pages ||| <p>foo</p>\n ||| <p>2015-03-24 22:28</p></br>\n<h2>Old Pages</h2>\n<p>Old unsupported pages :</p>\n ||| <a href=\"../index.html\">Index</a></br>\n</br>\n<a href=\"../categories/x1.html\">X1</a></br>\n<a href=\"../categories/x2.html\">X2</a> ||| <p>© foo</p>\n"
	
	filePath := filepath.Join(contentPath, slug.GenerateSlug(test.ContentDataTitle) + ".html")
	result := test.GetTestFileContent(filePath)
	
	test.TestValues(t, expected, result, "incorrect page generated")
}
