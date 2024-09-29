package sectionize

type HeadingPrefix string

const (
	HeadingPrefix1 HeadingPrefix = "#"
	HeadingPrefix2 HeadingPrefix = "##"
	HeadingPrefix3 HeadingPrefix = "###"
	HeadingPrefix4 HeadingPrefix = "####"
	HeadingPrefix5 HeadingPrefix = "#####"
	HeadingPrefix6 HeadingPrefix = "######"
)

const CommentStartsWith = "<!--"
const CommentEndsWith = "-->"
