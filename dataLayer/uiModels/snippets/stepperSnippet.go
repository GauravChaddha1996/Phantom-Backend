package snippets

import (
	"fmt"
	"phantom/dataLayer/uiModels/atoms"
)

type StepperSnippetData struct {
	Type          string          `json:"type,omitempty"`
	Title         *atoms.TextData `json:"title,omitempty"`
	StepperConfig StepperConfig   `json:"stepper_config"`
}

type StepperConfig struct {
}

func MakeStepperSnippet(cost int64) StepperSnippetData {
	return StepperSnippetData{
		Type:          StepperSnippet,
		Title:         &atoms.TextData{Text: fmt.Sprintf("Add for ₹%d", cost)},
		StepperConfig: StepperConfig{},
	}
}
