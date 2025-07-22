package functions

import (
	"strings"
)

func SplitInput(input string) []string {
	wordsSlice := strings.Split(input, "\\n")
	for i, val := range wordsSlice {
		if val != "" {
			break
		}
		if i == len(wordsSlice)-1 && val == "" {
			wordsSlice = wordsSlice[1:]
		}
	}
	return wordsSlice
}