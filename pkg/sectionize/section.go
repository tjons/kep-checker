package sectionize

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
)

type Sectionizer struct {
	sections       []*Section
	currentSection *Section
	content        []byte
	stripComments  bool
	trimHeadings   bool
}

type Section struct {
	Title string
	Body  string
	Lines []string
}

func NewSectionizer(content []byte) *Sectionizer {
	return &Sectionizer{
		sections:       make([]*Section, 0),
		currentSection: nil,
		content:        content,
	}
}

func (s *Sectionizer) WithStripComments(stripComments bool) *Sectionizer {
	s.stripComments = stripComments
	return s
}

func (s *Sectionizer) WithTrimHeadings(trimHeadings bool) *Sectionizer {
	s.trimHeadings = trimHeadings
	return s
}

func (s *Sectionizer) Sectionize() []*Section {
	reader := bytes.NewReader(s.content)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if IsHeading(line) {
			// close the current section and initiate a new one
			s.completeSection()
			_ = s.initSection(line)
		} else if s.currentSection != nil {
			s.writeToCurrentSection(line)
		}
	}
	s.completeSection()

	return s.sections
}

func (s *Sectionizer) completeSection() {
	if s.currentSection != nil {
		s.currentSection.Body = strings.Join(s.currentSection.Lines, "\n")
		if s.stripComments {
			s.currentSection.Body = trimComments(s.currentSection.Body)
		}

		s.sections = append(s.sections, s.currentSection)
		s.currentSection = nil
	}
}

func (s *Sectionizer) initSection(heading string) error {
	if s.currentSection != nil {
		return errors.New("You must close the current section before opening a new one")
	}

	if s.trimHeadings {
		heading = trimHeading(heading)
	}

	s.currentSection = &Section{
		Title: heading,
	}

	return nil
}

func (s *Sectionizer) writeToCurrentSection(text string) {
	s.currentSection.Lines = append(s.currentSection.Lines, text)
}

func IsHeading(line string) bool {
	if strings.HasPrefix(strings.TrimSpace(line), string(HeadingPrefix1)) {
		return true
	}

	return false
}
