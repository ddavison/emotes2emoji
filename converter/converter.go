package converter

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Emoji struct {
	Emotes map[string][]string
}

var emotes Emoji

func init() {
	data, err := ioutil.ReadFile("emotes.yaml")
	if err != nil {
		panic("Couldn't find/read emotes.yaml")
	}

	emotes = Emoji{}

	e := yaml.Unmarshal(data, &emotes)
	if e != nil {
		log.Fatalf("Malformed yaml and/or invalid unmarshal object match! %v", e)
	}
}

// Convert returns the converted version of a specific string
// E.g.: Convert(":)")  will return ":smile:"
func Convert(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		words[i] = GetEmojiFor(word)
	}
	return strings.Join(words, " ")
}

// GetEmojiFor returns the emoji conversion of a specific emote.
// GetEmojiFor will return the same thing it was passed if it does not find an
// emote for the word
func GetEmojiFor(emote string) string {
	toReturn := emote

	for conversion, emoji := range emotes.Emotes {
		for _, theEmote := range emoji {
			if theEmote == emote {
				toReturn = conversion
			}
		}
	}

	return toReturn
}

// ListEmotes will simply barf out the emotes that are available
func ListEmotes() {
	fmt.Println(emotes)
}
