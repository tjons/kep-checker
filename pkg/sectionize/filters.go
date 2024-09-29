package sectionize

import (
	"regexp"
	"strings"
)

const headingContentRegex = `\#{1,6}\s*(?P<content>.+)\s*$`

func trimHeading(line string) string {
	headingContent := regexp.MustCompile(headingContentRegex)
	matches := headingContent.FindStringSubmatch(line)
	contentIndex := headingContent.SubexpIndex("content")
	return matches[contentIndex]
}

const commentRegex = `(?P<comment><!--[^-->]+-->)`

func trimComments(body string) string {
	commentRexp := regexp.MustCompile(commentRegex)
	comments := commentRexp.FindAllString(body, -1)

	for i := range comments {
		body = strings.Replace(body, comments[i], "", -1)
	}

	return strings.TrimSpace(body)
}
