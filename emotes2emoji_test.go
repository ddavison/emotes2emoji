package main

import (
	"testing"

	"github.com/ddavison/emotes2emoji/converter"
)

var conversionTests = []struct {
	in  string
	out string
}{
	{":)", ":smile:"},      // returns conversion
	{":O", ":open_mouth:"}, // returns conversion
	{"test", "test"},       // returns itself
}

func TestConversions(t *testing.T) {
	for _, tt := range conversionTests {
		if result := converter.GetEmojiFor(tt.in); result != tt.out {
			t.Errorf("%v did not convert to %v! (value was: %v)", tt.in, tt.out, result)
		}
	}
}
