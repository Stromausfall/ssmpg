package category

import (
	"fmt"
	"github.com/stromausfall/ssmpg/generate/slug"
	"github.com/stromausfall/ssmpg/input"
	"sort"
)

type sortableByDate []input.ContentData

func (s sortableByDate) Len() int {
	return len(s)
}
func (s sortableByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortableByDate) Less(i, j int) bool {
	return s[i].Date.Before(s[j].Date)
}

func generateLink(url, linkText string) string {
	return "<a href=\"" + url + "\">" + linkText + "</a>"
}

func filterByCategory(category string, elements []input.ContentData) []input.ContentData {
	result := make([]input.ContentData, 0)

	for _, element := range elements {
		for _, elementCategory := range element.Categories {
			if category == elementCategory {
				result = append(result, element)
				break
			}
		}
	}

	return result
}

func CreateCategoryPage(configData input.ConfigData, category string, elements []input.ContentData) string {
	filteredElements := filterByCategory(category, elements)
	content := "<" + configData.TitleType + ">" + category + "</" + configData.TitleType + ">\n"
	content += createLinks(filteredElements, false)

	return content
}

func createLinks(elements []input.ContentData, fromIndex bool) string {
	sortedElements := elements[:]
	sort.Sort(sortableByDate(sortedElements))
	content := ""

	for _, element := range sortedElements {
		slug := slug.GenerateSlug(element.Title)
		dateString :=
			fmt.Sprintf(
				"(%d-%02d-%02d) ",
				element.Date.Year(),
				element.Date.Month(),
				element.Date.Day())
		rawLink := "content/" + slug + ".html"

		if !fromIndex {
			rawLink = "../" + rawLink
		}

		content += generateLink(rawLink, dateString+element.Title) + "</br>\n"
	}

	return content
}

func CreateIndex(configData input.ConfigData, elements []input.ContentData) string {
	content := "<" + configData.TitleType + ">" + configData.IndexName + "</" + configData.TitleType + ">\n"
	content += createLinks(elements, true)

	return content
}
