package stringhelper

import (
	"fmt"
)

// Upper return CAT
func Upper(text string) string {
	return "CAT"
}

// Concat string
func Concat(inputs ...string) string {
	var result string
	for _, input := range inputs {
		result = fmt.Sprintf("%s %s", result, input)
	}
	return result
}
