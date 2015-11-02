package sidebar

import (
	"fmt"
	"github.com/stromausfall/ssmpg/input"
	"github.com/stromausfall/ssmpg/utils/test"
	"testing"
	"strings"
)

var testContentData1 = input.ContentData { Categories : []string{ "fsdf", "sdf" }}
var testContentData2 = input.ContentData { Categories : []string{ "r1", "fsdf", "r2" }}
var testContentData3 = input.ContentData { Categories : []string{ "tfsdf", "sdf" }}
var testConfigData = input.ConfigData { IndexName : "Index" }

func TestGenerateSideBar(t *testing.T) {
	data := []input.ContentData { testContentData1 }
	sideBarContent := GenerateSidebar(testConfigData, data, false)
	expected := "<a href=\"../index.html\">Index</a></br>\n</br>\n<a href=\"../categories/fsdf.html\">fsdf</a></br>\n<a href=\"../categories/sdf.html\">sdf</a>"
	
	test.TestValues(t, expected, sideBarContent, "incorrect generated content")
}

func TestGenerateSideBarTestLinksWithDuplicateCategories(t *testing.T) {
	data := []input.ContentData { testContentData1, testContentData2 }
	sideBarContent := GenerateSidebar(testConfigData, data, false)
	linkCount := strings.Count(sideBarContent, "<a ")
	
	if linkCount != 5 {
		fmt.Println("expected : 5 links (4 categories + index)")
		fmt.Println("actual   : " + fmt.Sprintf("%v", linkCount) + " links")
		t.Error("incorrect generated content !")
	}
}

func TestGenerateSideBarTestLinksAreSorted(t *testing.T) {
	data := []input.ContentData { testContentData3 }
	sideBarContent := GenerateSidebar(testConfigData, data, false)
	expected := "<a href=\"../index.html\">Index</a></br>\n</br>\n<a href=\"../categories/sdf.html\">sdf</a></br>\n<a href=\"../categories/tfsdf.html\">tfsdf</a>"
	
	test.TestValues(t, expected, sideBarContent, "incorrect generated content")
}

func TestGenerateSideBarIndexNameIsUsed(t *testing.T) {
	testConfigData2 := input.ConfigData { IndexName : "Index33" }
	data := []input.ContentData { testContentData3 }
	sideBarContent := GenerateSidebar(testConfigData2, data, false)
	expected := "<a href=\"../index.html\">Index33</a></br>\n</br>\n<a href=\"../categories/sdf.html\">sdf</a></br>\n<a href=\"../categories/tfsdf.html\">tfsdf</a>"
	
	test.TestValues(t, expected, sideBarContent, "incorrect generated content")
}

func TestGenerateSideBarForIndexPage(t *testing.T) {
	testConfigData2 := input.ConfigData { IndexName : "Index33" }
	data := []input.ContentData { testContentData3 }
	sideBarContent := GenerateSidebar(testConfigData2, data, true)
	expected := "<a href=\"index.html\">Index33</a></br>\n</br>\n<a href=\"categories/sdf.html\">sdf</a></br>\n<a href=\"categories/tfsdf.html\">tfsdf</a>"
	
	test.TestValues(t, expected, sideBarContent, "incorrect generated content")
}
