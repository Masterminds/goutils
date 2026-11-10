package goutils

import (
	"testing"
	"unicode/utf8"
)

func TestWrapCustomUnicode(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		newline  string
		wrapLong bool
		want     string
	}{
		{"é é é", 3, "", false, "é é\né"},
		{"你好 世界", 5, "", false, "你好 世界"},
		{"你好 世界 再见", 5, "", false, "你好 世界\n再见"},
		{"ééé é", 2, "", false, "ééé\né"},
		{"你好世界", 2, "", true, "你好\n世界"},
		{"😀😃😄", 1, "|", true, "😀|😃|😄"},
		{"ééé", 0, "", true, "é\né\né"},
		{" 你好 世界", 2, "<br>", true, "你好<br>世界"},
	}
	for _, test := range tests {
		got := WrapCustom(test.input, test.width, test.newline, test.wrapLong)
		if got != test.want {
			t.Errorf("WrapCustom(%q, %d, %q, %t) = %q, want %q", test.input, test.width, test.newline, test.wrapLong, got, test.want)
		}
		if !utf8.ValidString(got) {
			t.Errorf("WrapCustom(%q, %d, %q, %t) returned invalid UTF-8", test.input, test.width, test.newline, test.wrapLong)
		}
	}
}

func TestWrapUnicode(t *testing.T) {
	if got, want := Wrap("αβ γδ εζ", 5), "αβ γδ\nεζ"; got != want {
		t.Errorf("Wrap() = %q, want %q", got, want)
	}
}
