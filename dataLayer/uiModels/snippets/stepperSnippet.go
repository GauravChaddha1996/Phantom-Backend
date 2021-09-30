package snippets

import (
	"fmt"
	"phantom/dataLayer/uiModels/atoms"
)

type StepperSnippet struct {
	Title         *atoms.TextData `json:"title,omitempty"`
	StepperConfig StepperConfig   `json:"stepper_config"`
}

type StepperConfig struct {
}

func MakeStepperSnippet(cost int64) StepperSnippet {
	return StepperSnippet{
		Title:         &atoms.TextData{Text: fmt.Sprintf("₹%d", cost)},
		StepperConfig: StepperConfig{},
	}
}
