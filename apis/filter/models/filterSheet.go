package models

import (
	"phantom/dataLayer/dbModels"
	"phantom/dataLayer/uiModels/atoms"
	"sort"
	"strings"
)

type FilterSheetUiData struct {
	PropertyUiSections []FilterSheetPropertyUiSection `json:"property_ui_sections,omitempty"`
}

type FilterSheetPropertyUiSection struct {
	Name           atoms.TextData                   `json:"name,omitempty"`
	PropertyValues []FilterSheetPropertyValueUiData `json:"property_values,omitempty"`
}

type FilterSheetPropertyValueUiData struct {
	Id       int64          `json:"id,omitempty"`
	Name     atoms.TextData `json:"name,omitempty"`
	Selected bool           `json:"selected,omitempty"`
}

func MakeFilterSheetUiData(
	apiRequest *ApiRequest,
	propertyToPropertyValueMap map[dbModels.Property][]dbModels.PropertyValue,
) FilterSheetUiData {
	var propertyUiSections []FilterSheetPropertyUiSection
	for property, propertyValues := range propertyToPropertyValueMap {
		propertyUiSection := makePropertyUiSection(propertyValues, apiRequest, property)
		propertyUiSections = append(propertyUiSections, propertyUiSection)
	}
	sort.SliceStable(propertyUiSections, func(one, two int) bool {
		propertyUiSectionOne := propertyUiSections[one]
		propertyUiSectionTwo := propertyUiSections[two]
		return propertyUiSectionOne.Name.Text < propertyUiSectionTwo.Name.Text
	})
	return FilterSheetUiData{PropertyUiSections: propertyUiSections}
}

func makePropertyUiSection(
	propertyValues []dbModels.PropertyValue,
	apiRequest *ApiRequest,
	property dbModels.Property,
) FilterSheetPropertyUiSection {
	var propertyValueUiDataArr []FilterSheetPropertyValueUiData
	for _, propertyValue := range propertyValues {
		propertyValueUiData := makePropertyValueUiData(propertyValue, apiRequest)
		propertyValueUiDataArr = append(propertyValueUiDataArr, propertyValueUiData)
	}
	propertyUiSection := FilterSheetPropertyUiSection{
		Name:           atoms.TextData{Text: strings.Title(property.Name)},
		PropertyValues: propertyValueUiDataArr,
	}
	return propertyUiSection
}

func makePropertyValueUiData(
	propertyValue dbModels.PropertyValue,
	apiRequest *ApiRequest,
) FilterSheetPropertyValueUiData {
	propertyValueUiData := FilterSheetPropertyValueUiData{
		Id:       propertyValue.Id,
		Name:     atoms.TextData{Text: strings.Title(propertyValue.Name)},
		Selected: apiRequest.PropertyValueIdsMap[propertyValue.Id],
	}
	return propertyValueUiData
}
