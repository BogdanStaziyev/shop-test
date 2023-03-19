package task2

import (
	"strings"
)

// concat function is proposed in the test task
func concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

// optimizeStringsBuilder are creating an optimized version using the strings package and
// strings.Builder type for efficient string creation using the Write methods, which minimizes memory copying
func optimizeStringsBuilder(str []string) string {
	var builder strings.Builder
	for _, v := range str {
		builder.WriteString(v)
	}
	return builder.String()
}

// optimizeStringsJoin strings.Join from the standard library, which is the most concise answer.
func optimizeStringsJoin(str []string) string {
	return strings.Join(str, "")
}
