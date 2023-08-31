package stringutils_test

import (
	"fmt"
	"github.com/seymasa/stringutils/stringutils"
	"testing"
)

func TestReverse(t *testing.T) {
	var testCases = []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"this is vigo", "ogiv si siht"},
		{"uğur", "ruğu"},
		{"kırmızı şapka ve ÖĞRENCİ", "İCNERĞÖ ev akpaş ızımrık"},
		{"Präzisionsmeßgerät", "täregßemsnoisizärP"},
	}

	for _, tc := range testCases {
		got := stringutils.Reverse(tc.input)
		if got != tc.want {
			t.Errorf("Reverse(%q) = %q; want %q", tc.input, got, tc.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = stringutils.Reverse("merhaba dünya!")
	}
	// use s to avoid compiler optimization
	_ = s
}

func ExampleReverse() {
	fmt.Println(stringutils.Reverse("vigo"))
	// Output: ogiv
}
