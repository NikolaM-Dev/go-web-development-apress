package strcon

import (
	"bytes"
	"unicode"
)

// Swap characters case from uppper to lower or lower to upper.
func SwapCase(str string) string {
	buf := &bytes.Buffer{}

	for _, r := range str {
		if unicode.IsUpper(r) {
			buf.WriteRune(unicode.ToLower(r))

			continue
		}

		buf.WriteRune(unicode.ToUpper(r))
	}

	return buf.String()
}
