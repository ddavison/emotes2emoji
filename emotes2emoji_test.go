package main

import (
	"testing"

	"github.com/ddavison/emotes2emoji/converter"
	"github.com/ddavison/emotes2emoji/hooker"
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

func TestDetectsGitRepo(t *testing.T) {
	if !hooker.IsGitRepository() {
		t.Errorf("Isn't detecting when there is a repo!")
	}
}
