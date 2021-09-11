package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const categoryToProductRailSectionCount = 1

func CategoryToProductRailSections(
	categoryToProductCacheDao *cacheDaos.CategoryIdToProductIdRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) []snippets.SnippetSectionData {
	categoryOne, categoryTwo := getTwoRandomCategories(apiDbResult.CategoriesMap)
	categoryOneRailSection := categoryToProductRailSection(categoryOne, categoryToProductCacheDao, productIdMap, apiDbResult)
	categoryTwoRailSection := categoryToProductRailSection(categoryTwo, categoryToProductCacheDao, productIdMap, apiDbResult)
	return []snippets.SnippetSectionData{categoryOneRailSection, categoryTwoRailSection}
}

func getTwoRandomCategories(categoryIdMap map[int64]*dbModels.Category) (*dbModels.Category, *dbModels.Category) {
	var categoryOne *dbModels.Category
	var categoryTwo *dbModels.Category

	for _, category := range categoryIdMap {
		if categoryOne == nil {
			categoryOne = category
			continue
		}
		if categoryTwo == nil {
			categoryTwo = category
			break
		}
	}
	return categoryOne, categoryTwo
}

func categoryToProductRailSection(
	category *dbModels.Category,
	categoryToProductCacheDao *cacheDaos.CategoryIdToProductIdRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) snippets.SnippetSectionData {
	var productsOfCategorySnippets []snippets.ProductRailSnippet

	productsOfCategoryId, err := categoryToProductCacheDao.ReadNProductsOfCategoryId(&category.Id, categoryToProductRailSectionCount)
	if err == nil {
		productIdMap.PutAllInt64s(productsOfCategoryId)
		for _, productId := range *productsOfCategoryId {
			snippet := getProductRailSnippet(productId, apiDbResult)
			productsOfCategorySnippets = append(productsOfCategorySnippets, snippet)
		}
	}
	return snippets.SnippetSectionData{
		Type: snippets.ProductRailSection,
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{Text: category.Name},
		},
		Snippets: apiCommons.ToBaseSnippets(productsOfCategorySnippets),
	}
}
