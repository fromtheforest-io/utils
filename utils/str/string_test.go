package str

import "testing"

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World!", "hello-world"},
		{"This is a Test", "this-is-a-test"},
		{"Go & Rust", "go-rust"},
		{"Already-slugified", "already-slugified"},
		{"  Lots   of    spaces  ", "lots-of-spaces"},
		{"Special__Chars##123", "special-chars-123"},
		{"--Leading and trailing--", "leading-and-trailing"},
		{"Emoji 👍🏽 text", "emoji-text"},
	}

	for _, test := range tests {
		got := Slugify(test.input)
		if got != test.expected {
			t.Errorf("Slugify(%q) = %q; expected %q", test.input, got, test.expected)
		}
	}
}
