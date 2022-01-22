package section

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"phantom/apis/apiCommons"
	"phantom/apis/home/models"
	"phantom/dataLayer/cacheDaos"
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
)

const categoryToProductRailSectionCount = 4
const seeAll = "See all"

func CategoryToProductRailSections(
	ctx *gin.Context,
	categoryToProductCacheDao *cacheDaos.CategoryIdToProductIdRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) []*snippets.SnippetSectionData {
	categoryOne, categoryTwo := getTwoRandomCategories(apiDbResult.CategoriesMap)
	categoryOneRailSection := categoryToProductRailSection(ctx, categoryOne, categoryToProductCacheDao, productIdMap, apiDbResult)
	categoryTwoRailSection := categoryToProductRailSection(ctx, categoryTwo, categoryToProductCacheDao, productIdMap, apiDbResult)
	return []*snippets.SnippetSectionData{categoryOneRailSection, categoryTwoRailSection}
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
	ctx *gin.Context,
	category *dbModels.Category,
	categoryToProductCacheDao *cacheDaos.CategoryIdToProductIdRedisDao,
	productIdMap *apiCommons.ProductIdMap,
	apiDbResult models.ApiDbResult,
) *snippets.SnippetSectionData {
	var productsOfCategorySnippets []snippets.ProductRailSnippetData

	productsOfCategoryId, err := categoryToProductCacheDao.ReadNProductsOfCategoryId(&category.Id, categoryToProductRailSectionCount)
	if err != nil {
		logData := apiCommons.NewApiErrorLogData(ctx, "Error reading first N Products of category from cache", err)
		logData.Data["category_id"] = cast.ToString(category.Id)
		apiCommons.LogApiError(logData)
		return nil
	}
	productIdMap.PutAllInt64s(productsOfCategoryId)
	for _, productId := range *productsOfCategoryId {
		product := apiDbResult.ProductsMap[productId]
		category := apiDbResult.CategoriesMap[product.CategoryId]
		brand := apiDbResult.BrandsMap[product.BrandId]
			snippet := snippets.MakeProductRailSnippet(*product, *category, *brand, apiDbResult.NewProductIdsMap.Contains(productId))
		productsOfCategorySnippets = append(productsOfCategorySnippets, snippet)
	}
	return &snippets.SnippetSectionData{
		HeaderData: &snippets.SnippetSectionHeaderData{
			Title: &atoms.TextData{
				Text: category.Name,
			},
			RightButton: &atoms.ButtonData{
				Type: atoms.ButtonTypeText,
				Text: atoms.TextData{Text: seeAll},
				Click: atoms.CategoryClickData{
					Type:       atoms.ClickTypeOpenCategory,
					CategoryId: category.Id,
				},
			},
		},
		Snippets: &productsOfCategorySnippets,
	}
}
