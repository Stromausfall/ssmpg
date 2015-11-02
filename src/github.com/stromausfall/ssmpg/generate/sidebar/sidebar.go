package sidebar

import (
	"sort"
	"github.com/stromausfall/ssmpg/input"
	"github.com/stromausfall/ssmpg/generate/slug"
)

func generateLink(url, linkText string) string {
	return "<a href=\"" + url + "\">" + linkText + "</a>"
}

func CollectCategories(elements []input.ContentData) []string {
	content := make(map[string]bool)
	keys := sort.StringSlice{}
	
	// makes sure there are no duplicates
	for _, element := range elements {
		for _, category := range element.Categories {
			content[category] = true
		}
	}
	
	// create an array from it
	for key, _ := range content {
		keys = append(keys, key)
	}
	
	// sort the keys
	sort.Sort(keys)
	
	return keys
}

func GenerateSidebar(configData input.ConfigData, elements []input.ContentData, forIndexPage bool) string {
	categories := CollectCategories(elements)
	linkPrefix := "../"
	
	if forIndexPage {
		linkPrefix = ""
	}
	
	result := generateLink(linkPrefix + "index.html", configData.IndexName) + "</br>\n"
	
	for _, category := range categories {
		slug := slug.GenerateSlug(category)
		rawLink := linkPrefix + "categories/" + slug + ".html"
		result += "</br>\n" + generateLink(rawLink, category)
	}
	
	return result
}
