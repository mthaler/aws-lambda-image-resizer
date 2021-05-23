package helpers

import (
	"testing"
)

func TestIsImage(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{ "foo.jpg", true },
		{ "FOO.JPG", true },
		{ "foo.png", true },
		{ "FOO.PNG", true },
		{ "foo.gif", true },
		{ "FOO.GIF", true },
		{ "foo.txt", false },
		{ "FOO.TXT", false },
	}
	for _, tt := range tests {
		if got := IsImage(tt.name); got != tt.want {
			t.Errorf("IsImage() = %v, want %v", got, tt.want)
		}
	}
}
