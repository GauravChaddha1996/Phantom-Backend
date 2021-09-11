package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const categoryRailSectionHeader = "Categories"

func CategoryRailSection(
	categoryDao *cacheDaos.AllCategoryIdsRedisDao,
	apiDbResult models.ApiDbResult,
) snippets.SnippetSectionData {
	var categoryRailSnippets []snippets.CategoryRailSnippet
	allCategoryIds, err := categoryDao.ReadAllCategoryIds()
	if err == nil {
		for _, categoryId := range *allCategoryIds {
			snippet := getCategoryRailSnippet(apiDbResult, categoryId)
			categoryRailSnippets = append(categoryRailSnippets, snippet)
		}
	}
	return snippets.SnippetSectionData{
		Type: snippets.CategoryRailSection,
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{Text: categoryRailSectionHeader},
		},
		Snippets: apiCommons.ToBaseSnippets(categoryRailSnippets),
	}
}
