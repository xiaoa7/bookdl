package util

import (
	"regexp"
	"strings"
)

//使用正则进行清理
var (
	clean_html_tag, single_html_tag, br_html_tag *regexp.Regexp
	CRLF                                         = "\r\n"
)

//
func init() {
	clean_html_tag, _ = regexp.Compile("<[^>]+>.*</[^>]+>")
	br_html_tag, _ = regexp.Compile("<br[^>]*>")
	single_html_tag, _ = regexp.Compile("<[^>]+>")
}

//
func Clean(str string) string {
	str = clean_html_tag.ReplaceAllString(str, "")
	str = br_html_tag.ReplaceAllString(str, CRLF)
	str = single_html_tag.ReplaceAllString(str, "")
	str = strings.Replace(str, "&nbsp;", " ", -1)
	return str
}
