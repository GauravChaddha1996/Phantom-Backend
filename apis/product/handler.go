package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phantom/apis/product/models"
	"phantom/apis/product/section"
	"phantom/dataLayer/uiModels/snippets"
)

const apiDbReadErr = "Err code: 1"

func ApiHandler(ctx *gin.Context) {
	// Read api request model
	apiRequest, apiRequestReadErr := models.ReadApiRequestModel(ctx)
	if apiRequestReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiRequestReadErr.Error())
		return
	}
	// Get db results
	apiDbResult, dbReadErr := readFromDb(ctx, apiRequest)
	if dbReadErr != nil {
		ctx.JSON(http.StatusInternalServerError, apiDbReadErr)
		return
	}

	// Make all sections
	imagesSection := section.MakeImagesSection(apiDbResult)
	headerSection := section.MakeHeaderSection(apiDbResult)
	longDescSection := section.MakeLongDescSection(apiDbResult)
	propertyMappingSection := section.MakePropertyMappingSection(apiDbResult)

	// Arrange sections
	var sections []*snippets.SnippetSectionData
	sections = append(sections, imagesSection)
	sections = append(sections, headerSection)
	sections = append(sections, longDescSection)
	sections = append(sections, propertyMappingSection)

	// Make api response
	apiResponse := models.ProductApiResponse{
		Status:   "success",
		Message:  "",
		Snippets: sections,
		StepperSnippet: snippets.MakeStepperSnippet(
			*apiDbResult.Product,
			*apiDbResult.Brand,
			*apiDbResult.Category,
		),
	}

	ctx.JSON(http.StatusOK, apiResponse)
}
