package converter

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/ddavison/emotes2emoji/utils"
	"gopkg.in/yaml.v2"
)

type Emoji struct {
	Emotes map[string][]string
}

var emotes Emoji

func init() {
	// If the emotes2emoji.yaml file does not exist in the home directory, put it there
	usr, _ := user.Current()

	filePath := usr.HomeDir + "/.emotes2emoji/emotes.yaml"
	if !utils.FileExists(filePath) {
		wd, _ := os.Getwd()
		os.Mkdir(usr.HomeDir+"/.emotes2emoji", 0777)
		os.Symlink(wd+"/emotes.yaml", filePath)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("Couldn't find/read " + filePath)
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
	log.Print(emotes)
}
