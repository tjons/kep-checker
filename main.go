package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tjons/kep-checker/pkg/goldenkep"
	"github.com/tjons/kep-checker/pkg/sectionize"
)

var (
	KepUrl, DefaultKepTemplateURL string
)

// should actually pull the kep directory from the users fork or from the PR -
// this is much more powerful and helpful
func init() {
	flag.StringVar(&KepUrl, "kep-url", "", "required: URL to KEP to validate")
	flag.StringVar(&DefaultKepTemplateURL, "kep-template-url", "https://raw.githubusercontent.com/kubernetes/enhancements/master/keps/NNNN-kep-template/README.md", "URL to KEP template")
	flag.Parse()

	if KepUrl == "" {
		fmt.Println("kep-url is required")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Getting latest KEP template...")

	resp, err := http.Get(DefaultKepTemplateURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error getting KEP template:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	template, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading KEP template:", err)
		os.Exit(1)
	}

	fmt.Println(string(template))

	fmt.Println("Getting KEP to validate...")
	resp, err = http.Get(KepUrl)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error getting KEP:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	kep, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading KEP:", err)
		os.Exit(1)
	}

	s := sectionize.NewSectionizer(kep).
		WithStripComments(true).
		WithTrimHeadings(true)

	sections := s.Sectionize()

	for i := range sections {
		fmt.Printf("Section: %s\nWords: %d\n", sections[i].Title, len(strings.Split(sections[i].Body, " ")))
	}

	// we need to take the template and extract sections from it
	// for our purposes, a section is a heading followed by a list of lines
	templateSections := goldenkep.GetSections()
	for _, templateSection := range templateSections {
		exists := false
		for _, section := range sections {
			if templateSection.Title == section.Title {
				exists = true
			}
		}

		if !exists {
			fmt.Printf("KEP is not using the latest template: section %s is missing", templateSection.Title)
		}
	}
}
