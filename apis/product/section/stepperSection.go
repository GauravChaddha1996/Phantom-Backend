package section

import (
	"phantom/apis/apiCommons"
	"phantom/apis/product/models"
	"phantom/dataLayer/uiModels/snippets"
)

func MakeStepperSection(apiDbResult *models.ApiDbResult) *snippets.SnippetSectionData {
	stepperSnippet := snippets.MakeStepperSnippet(apiDbResult.Product.Cost)
	return &snippets.SnippetSectionData{
		Type:       snippets.StepperSection,
		HeaderData: nil,
		Snippets:   apiCommons.ToBaseSnippets(stepperSnippet),
	}
}
