package utils

import (
	"fmt"
	"os"
	"bytes"
	"strings"
)

func revert(s string) string {
	var b bytes.Buffer
	l := len(s) - 1
	for ; l >= 0; l-- {
		b.WriteRune(rune(s[l]))
	}
	return b.String()
}

func strongComma(s string) string {
	if strings.Contains(s, ".") {
		parts := strings.Split(s, ".")
		left := comma(parts[0])
		right := revert(comma(revert(parts[1])))
		return left + "." + right
	} else {
		return comma(s)
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var b bytes.Buffer
	var c int
	n := len(s)
	c = n % 3
	if c == 0 {
		c = 3
	}
	b.Write([]byte(s[:c]))
	for ; c < n; c = c + 3 {
		b.Write([]byte(","))
		b.Write([]byte(s[c:c + 3]))
	}
	return b.String();
}

//!-