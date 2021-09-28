package models

import (
	"phantom/dataLayer"
	"phantom/dataLayer/uiModels/atoms"
)

type SortSheetUiData struct {
	Methods []SortMethodUiData `json:"methods,omitempty"`
}

type SortMethodUiData struct {
	Id       int64          `json:"id"`
	Name     atoms.TextData `json:"name,omitempty"`
	Selected bool           `json:"selected,omitempty"`
}

func MakeSortSheetUiData(selectedSortMethodId int64) SortSheetUiData {
	var sortSheetMethods []SortMethodUiData
	for _, sortMethod := range dataLayer.AllSortMethods {
		sortMethodUiData := SortMethodUiData{
			Id:       sortMethod.Id,
			Name:     atoms.TextData{Text: sortMethod.Title},
			Selected: sortMethod.Id == selectedSortMethodId,
		}
		sortSheetMethods = append(sortSheetMethods, sortMethodUiData)
	}
	return SortSheetUiData{Methods: sortSheetMethods}
}
