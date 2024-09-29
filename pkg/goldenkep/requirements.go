package goldenkep

import (
	"regexp"

	"github.com/tjons/kep-checker/pkg/sectionize"
)

type Stage string

const (
	Alpha       Stage = "alpha"
	Beta        Stage = "beta"
	Stable      Stage = "stable"
	Deprecation Stage = "deprecation"
)

type GoldenKepSection struct {
	Title         string
	Required      bool
	RequiredStage Stage
}

func GetSections() []*GoldenKepSection {
	sectionizer := sectionize.NewSectionizer([]byte(Contents)).
		WithStripComments(false).
		WithTrimHeadings(true)

	sections := sectionizer.Sectionize()

	gks := make([]*GoldenKepSection, 0, len(sections))
	for _, section := range sections {
		gk := &GoldenKepSection{
			Title: section.Title,
		}
		parseDirectives(section, gk)
		gks = append(gks, gk)
	}

	return gks
}

const directivesRegex = `DIRECTIVE: required:\s*(?<required>true|false)?(,stage:\s)?(?<stage>[a-zA-Z]*)?`

func parseDirectives(usection *sectionize.Section, gksection *GoldenKepSection) {
	re := regexp.MustCompile(directivesRegex)
	matches := re.FindStringSubmatch(usection.Body)
	requiredIndex := re.SubexpIndex("required")

	if matches[requiredIndex] == "true" {
		gksection.Required = true

		stageIndex := re.SubexpIndex("stage")
		if stageIndex != -1 && matches[stageIndex] != "" {
			gksection.RequiredStage = Stage(matches[stageIndex])
		} else {
			gksection.RequiredStage = Alpha
		}
	}
}
