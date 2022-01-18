package section

import (
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeStepperSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	stepperSnippet := snippets.MakeStepperSnippet(apiDbResult.Product.Cost)
	return &snippets.SnippetSectionData{
		HeaderData: nil,
		Snippets:   []snippets.StepperSnippetData{stepperSnippet},
	}
}
