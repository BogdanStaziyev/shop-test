package task2

import (
	"fmt"
	"testing"
)

func testPreparation() []string {
	testStrings := make([]string, 30)
	for i := 0; i < len(testStrings); i++ {
		testStrings[i] = fmt.Sprintf("string%d", i)
	}
	return testStrings
}

func BenchmarkDefaultFunc(b *testing.B) {
	testStrings := testPreparation()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concat(testStrings)
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	testStrings := testPreparation()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimizeStringsBuilder(testStrings)
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	testStrings := testPreparation()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		optimizeStringsJoin(testStrings)
	}
}
