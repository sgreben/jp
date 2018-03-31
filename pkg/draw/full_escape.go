package draw

import "strings"

const invertedSpace = "\033[7m \033[27m"

func FullEscape(full string) string {
	return strings.Replace(full, string(fullBlock), invertedSpace, -1)
}
