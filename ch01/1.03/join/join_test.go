package join_test

import (
	"strings"
	"testing"

	"github.com/mi-bear/gopl.io/ch1/1.3/join"
)

func TestJoin(t *testing.T) {
	s := []string{"1", "2", "3"}
	result := join.Join(s)
	if result != "1 2 3" {
		t.Errorf("actual: '%s', expected: '1 2 3'", result)
	}
}

func BenchmarkJoin(b *testing.B) {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < b.N; i++ {
		join.Join(s)
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < b.N; i++ {
		strings.Join(s, " ")
	}
}
