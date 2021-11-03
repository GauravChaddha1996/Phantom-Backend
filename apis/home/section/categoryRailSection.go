package section

import (
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const categoryRailSectionHeader = "Categories"

func CategoryRailSection(
	ctx *gin.Context,
	categoryDao *cacheDaos.AllCategoryIdsRedisDao,
	apiDbResult models.ApiDbResult,
) *snippets.SnippetSectionData {
	var categoryRailSnippets []snippets.CategoryRailSnippetData

	allCategoryIds, err := categoryDao.ReadAllCategoryIds()
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading all category ids from cache", err)
		apiCommons.LogApiError(logData)
		return nil
	}

	for _, categoryId := range *allCategoryIds {
		snippet := snippets.MakeCategoryRailSnippet(*apiDbResult.CategoriesMap[categoryId])
		categoryRailSnippets = append(categoryRailSnippets, snippet)
	}
	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{Text: categoryRailSectionHeader},
		},
		Snippets: &categoryRailSnippets,
	}
}
