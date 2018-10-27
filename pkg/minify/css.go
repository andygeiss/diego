package minify

import (
	"regexp"
)

// CSS ...
func CSS(in string) string {
	out := in
	out = regexp.MustCompile(`\n`).ReplaceAllString(out, "")         // Newline
	out = regexp.MustCompile(`\r`).ReplaceAllString(out, "")         // Newline
	out = regexp.MustCompile(`\t`).ReplaceAllString(out, " ")        // Tab
	out = regexp.MustCompile(`\s{2,}`).ReplaceAllString(out, "")     // Whitespaces
	out = regexp.MustCompile(`\s*\{\s*`).ReplaceAllString(out, "{")  // Selector
	out = regexp.MustCompile(`\s*}`).ReplaceAllString(out, "}")      // End
	out = regexp.MustCompile(`\s+\>\s+`).ReplaceAllString(out, ">")  // Selector Child
	out = regexp.MustCompile(`\,\s+`).ReplaceAllString(out, ",")     // Comma
	out = regexp.MustCompile(`\/\*.*\*\/`).ReplaceAllString(out, "") // Comment
	out = regexp.MustCompile(`\s*\;\s*`).ReplaceAllString(out, ";")  // Attribute Separator
	out = regexp.MustCompile(`\s*\:\s*`).ReplaceAllString(out, ":")  // Attribute Key:Value
	out = regexp.MustCompile(`\s*\+\s*`).ReplaceAllString(out, "+")  // Selector Combination
	return out
}
