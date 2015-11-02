package category

import (
	"fmt"
	"github.com/stromausfall/ssmpg/generate/slug"
	"github.com/stromausfall/ssmpg/input"
	"strings"
	"testing"
	"time"
)

var testConfigData = input.ConfigData{TitleType: "h3", IndexName: "Index"}

func getTime(toParse string) time.Time {
	time, err := time.Parse(time.RFC822, toParse)

	if err != nil {
		panic("proplem when trying to parse : " + toParse)
	}

	return time
}

var testContentData1 = input.ContentData{
	Title:      "Title#1",
	Date:       getTime("02 Jan 15 10:00 UTC"),
	Categories: []string{"fsdf", "sdf"},
	Content:    "foo"}
var testContentData2 = input.ContentData{
	Title:      "Title#2",
	Date:       getTime("01 Jan 15 10:00 UTC"),
	Categories: []string{"sdf", "sdf"},
	Content:    "foo"}
var testContentData3 = input.ContentData{
	Title:      "Title#3",
	Date:       getTime("01 Jan 14 10:00 UTC"),
	Categories: []string{"assdf"},
	Content:    "foo"}

func TestCreateIndexPage(t *testing.T) {
	data := CreateIndex(testConfigData, []input.ContentData{testContentData3, testContentData2})
	expected :=
		"<h3>Index</h3>\n" +
			"<a href=\"content/title-3.html\">(2014-01-01) Title#3</a></br>\n" +
			"<a href=\"content/title-2.html\">(2015-01-01) Title#2</a></br>\n"

	if data != expected {
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("data     : '" + data + "'")
		t.Error("incorrect index page")
	}
}

func TestCreateIndexPageOrdered(t *testing.T) {
	data := CreateIndex(testConfigData, []input.ContentData{testContentData1, testContentData2, testContentData3})

	index1 := strings.Index(data, slug.GenerateSlug(testContentData1.Title))
	index2 := strings.Index(data, slug.GenerateSlug(testContentData2.Title))
	index3 := strings.Index(data, slug.GenerateSlug(testContentData3.Title))

	correctOrder := (index3 < index2) && (index2 < index1)
	validIndices := (index1 > 0) && (index2 > 0) && (index3 > 0)

	if !validIndices {
		t.Error("invalid  indices (the titles are not present) !")
	}

	if !correctOrder {
		t.Error("incorrect order !")
	}
}

func TestCreateCategoryPage(t *testing.T) {
	data :=
		CreateCategoryPage(
			testConfigData,
			"sdf",
			[]input.ContentData{testContentData1, testContentData2, testContentData3})
	expected :=
		"<h3>sdf</h3>\n" +
			"<a href=\"../content/title-2.html\">(2015-01-01) Title#2</a></br>\n" +
			"<a href=\"../content/title-1.html\">(2015-01-02) Title#1</a></br>\n"

	if data != expected {
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("data     : '" + data + "'")
		t.Error("incorrect category page")
	}
}

func TestCreateIndexChangeIndexAndTitle(t *testing.T) {
	testConfigData2 := input.ConfigData{TitleType: "h2", IndexName: "Index2"}
	data := CreateIndex(testConfigData2, []input.ContentData{testContentData3, testContentData2})
	expected :=
		"<h2>Index2</h2>\n" +
			"<a href=\"content/title-3.html\">(2014-01-01) Title#3</a></br>\n" +
			"<a href=\"content/title-2.html\">(2015-01-01) Title#2</a></br>\n"

	if data != expected {
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("data     : '" + data + "'")
		t.Error("incorrect index page")
	}
}

func TestCreateCategoryPageChangeCategoryAndTitle(t *testing.T) {
	testConfigData2 := input.ConfigData{TitleType: "h2"}
	data :=
		CreateCategoryPage(
			testConfigData2,
			"fsdf",
			[]input.ContentData{testContentData1, testContentData2, testContentData3})
	expected :=
		"<h2>fsdf</h2>\n" +
			"<a href=\"../content/title-1.html\">(2015-01-02) Title#1</a></br>\n"

	if data != expected {
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("data     : '" + data + "'")
		t.Error("incorrect category page")
	}
}
