package actions

import (
	"regexp"
)

func CleanColorCodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[mK]`)
	return re.ReplaceAllString(input, "")
}
