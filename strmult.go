// strmult provides methods to manipulate multilined strings
package strmult

import "strings"

func IsMultiLined(s string) bool {
	return strings.Contains(s, "\n")
}

func SplitStringIntoLines(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, "\n")
}

func TotalLines(s string) int {
	return len(SplitStringIntoLines(s))
}
