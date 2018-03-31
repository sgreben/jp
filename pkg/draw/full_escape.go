package draw

import "strings"

var invertEscape = replacer(map[rune]string{
	'█': "\033[7m \033[27m",
})

var colorEscapeBW = replacer(map[rune]string{
	'█': "\033[48;5;231m \033[49m",
	'▓': "\033[48;5;252m \033[49m",
	'▒': "\033[48;5;248m \033[49m",
	'░': "\033[48;5;240m \033[49m",
	'·': "\033[48;5;236m \033[49m",
	' ': "\033[48;5;232m ",
})

var colorEscapeWB = replacer(map[rune]string{
	'█': "\033[48;5;232m \033[49m",
	'▓': "\033[48;5;236m \033[49m",
	'▒': "\033[48;5;240m \033[49m",
	'░': "\033[48;5;248m \033[49m",
	'·': "\033[48;5;252m \033[49m",
	' ': "\033[48;5;231m ",
})

func replacer(m map[rune]string) *strings.Replacer {
	r := make([]string, 0, len(m)*2)
	for old, new := range m {
		r = append(r, string(old))
		r = append(r, new)
	}
	return strings.NewReplacer(r...)
}

func FullEscape(full string) string {
	return invertEscape.Replace(full)
}

func FullEscapeBW(full string) string {
	return colorEscapeBW.Replace(full)
}

func FullEscapeWB(full string) string {
	return colorEscapeWB.Replace(full)
}
