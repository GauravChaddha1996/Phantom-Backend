package section

import (
	"phantom/apis/home/models"
	"phantom/dataLayer/uiModels/atoms"
	"phantom/dataLayer/uiModels/snippets"
	"strconv"
)

func getProductRailSnippet(productId int64, apiDbResult models.ApiDbResult) snippets.ProductRailSnippet {
	product := apiDbResult.ProductsMap[productId]
	category := apiDbResult.CategoriesMap[product.CategoryId]
	brand := apiDbResult.BrandsMap[product.BrandId]

	snippet := snippets.ProductRailSnippet{
		Id:        product.Id,
		Name:      &atoms.TextData{Text: product.Name},
		ShortDesc: &atoms.TextData{Text: product.ShortDescription},
		Brand:     &atoms.TextData{Text: brand.Name},
		Category:  &atoms.TextData{Text: category.Name},
		Cost:      &atoms.TextData{Text: strconv.FormatInt(product.Cost, 10)},
		Image:     &atoms.ImageData{Url: product.CardImage},
	}
	return snippet
}

func getProductFullSnippet(productId int64, apiDbResult models.ApiDbResult) snippets.ProductFullSnippet {
	product := apiDbResult.ProductsMap[productId]
	category := apiDbResult.CategoriesMap[product.CategoryId]
	brand := apiDbResult.BrandsMap[product.BrandId]

	snippet := snippets.ProductFullSnippet{
		Id:       product.Id,
		Name:     &atoms.TextData{Text: product.Name},
		LongDesc: &atoms.TextData{Text: product.LongDescription},
		Brand:    &atoms.TextData{Text: brand.Name},
		Category: &atoms.TextData{Text: category.Name},
		Cost:     &atoms.TextData{Text: strconv.FormatInt(product.Cost, 10)},
		Image:    &atoms.ImageData{Url: product.CardImage},
	}
	return snippet
}

func getProductDualSnippet(productId int64, apiDbResult models.ApiDbResult) snippets.ProductDualSnippet {
	product := apiDbResult.ProductsMap[productId]
	brand := apiDbResult.BrandsMap[product.BrandId]

	snippet := snippets.ProductDualSnippet{
		Id:        product.Id,
		Name:      &atoms.TextData{Text: product.Name},
		ShortDesc: &atoms.TextData{Text: product.ShortDescription},
		Brand:     &atoms.TextData{Text: brand.Name},
		Cost:      &atoms.TextData{Text: strconv.FormatInt(product.Cost, 10)},
		Image:     &atoms.ImageData{Url: product.CardImage},
	}
	return snippet
}

func getCategoryRailSnippet(apiDbResult models.ApiDbResult, categoryId int64) snippets.CategoryRailSnippet {
	category := apiDbResult.CategoriesMap[categoryId]
	snippet := snippets.CategoryRailSnippet{
		Id:   category.Id,
		Name: &atoms.TextData{Text: category.Name},
	}
	return snippet
}