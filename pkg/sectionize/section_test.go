package sectionize

import "testing"

func TestSectionize(t *testing.T) {
	rawMD := []byte(`
		# Section 1
		Section 1 body
		Line 2
		Line 3
		## Section 2
		Section 2 body
		Line 2
		Line 3
		### Section 3
		Section 3 body
		`)
	sectionizer := NewSectionizer(rawMD)

	sections := sectionizer.Sectionize()
	if len(sections) != 3 {
		t.Errorf("Expected 3 sections, got %d", len(sections))
	}
}

func TestSectionizeWithStripComments(t *testing.T) {
	testcases := []struct {
		input    []byte
		expected *Section
	}{
		{
			input: []byte(`
# Heading
<!-- this is a comment -->
this is valid body text`),
			expected: &Section{
				Title: "Heading",
				Body:  "this is valid body text",
			},
		},
		{
			input: []byte(`
# Heading 1
<!--
this is
a very poorly
formed comment-->
and here is
my body text`),
			expected: &Section{
				Title: "Heading 1",
				Body:  "and here is\nmy body text",
			},
		},
	}

	for i, tt := range testcases {
		s := NewSectionizer(tt.input).
			WithStripComments(true).
			WithTrimHeadings(true)

		sections := s.Sectionize()
		if count := len(sections); count != 1 {
			t.Errorf("Testcase %d: Expected exactly 1 section, got %d", i, count)
		}

		if sections[0].Title != tt.expected.Title {
			t.Errorf("Testcase %d: Expected title %s to equal \n\n %s", i, sections[0].Title, tt.expected.Title)
		}

		if sections[0].Body != tt.expected.Body {
			t.Errorf("Testcase %d: Expected \n\n%s\n\n to equal \n\n%s\n\n", i, sections[0].Body, tt.expected.Body)
		}
	}
}

func TestTrimHeading(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{
			input:    "# Section 1",
			expected: "Section 1",
		},
		{
			input:    "## Section 2",
			expected: "Section 2",
		},
		{
			input:    "### Section 3",
			expected: "Section 3",
		},
		{
			input:    "#### Section 4",
			expected: "Section 4",
		},
		{
			input:    "##### Section 5",
			expected: "Section 5",
		},
		{
			input:    "###### Section 6",
			expected: "Section 6",
		},
	}

	for _, tt := range testcases {
		result := trimHeading(tt.input)
		if result != tt.expected {
			t.Errorf("expected trimHeading with input %s to result in %s, got %s",
				tt.input, tt.expected, result)
		}
	}
}
