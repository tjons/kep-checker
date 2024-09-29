package goldenkep

import (
	"testing"

	"github.com/tjons/kep-checker/pkg/sectionize"
)

const expectedGoldenSections = 57

func TestNumberOfSections(t *testing.T) {
	sections := GetSections()
	if len(sections) != expectedGoldenSections {
		t.Errorf("Expected %d sections from parsing Golden KEP... got %d", expectedGoldenSections, len(sections))
	}
}

func TestParseDirectives(t *testing.T) {
	testcases := []struct {
		bodyText        string
		required        bool
		requiredByStage Stage
	}{
		{
			bodyText:        "<!--DIRECTIVE: required: true-->",
			required:        true,
			requiredByStage: Alpha,
		}, {
			bodyText:        "<!--DIRECTIVE: required: true,stage: alpha-->",
			required:        true,
			requiredByStage: Alpha,
		},
		{
			bodyText:        "<!--DIRECTIVE: required: true,stage: beta-->",
			required:        true,
			requiredByStage: Beta,
		},
		{
			bodyText:        "<!--DIRECTIVE: required: true,stage: stable-->",
			required:        true,
			requiredByStage: Stable,
		},
		{
			bodyText:        "<!--DIRECTIVE: required: false-->",
			required:        false,
			requiredByStage: "",
		},
	}

	for _, tt := range testcases {
		usection := &sectionize.Section{
			Body: tt.bodyText,
		}

		gksection := &GoldenKepSection{
			Title: "something random",
		}

		parseDirectives(usection, gksection)

		if gksection.Required != tt.required || gksection.RequiredStage != tt.requiredByStage {
			t.Errorf("expected %s to be %t at stage %s, got %t at %s", usection.Body, tt.required, tt.requiredByStage, gksection.Required, gksection.RequiredStage)
		}
	}
}
