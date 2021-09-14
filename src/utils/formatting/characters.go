package formatting

import (
	"log"
	"strings"
)

func ReplaceSpecialCharacters(slug *string) {
	searchAndRemplace := func(old, new string) {
		*slug = strings.ReplaceAll(*slug, old, new)
	}

	*slug = strings.ToLower(*slug)

	specialCharacters := []string{"/", ".", "#", "[", "]", "&", "/", "*", "+", "$", "'", "(", ")", "=", "|", "°", ",", "%"}
	for _, character := range specialCharacters {
		searchAndRemplace(character, "")
	}

	letters := map[string]string{
		"n": "ñ",
		"y": "&",
		"e": "é",
		"o": "ó",
		"i": "í",
		"u": "ú",
		"a": "á",
	}

	for new, old := range letters {
		searchAndRemplace(old, new)
	}

	searchAndRemplace(" ", "-")
	searchAndRemplace("ü", "u")

	*slug = strings.TrimRight(*slug, "-")

	log.Println(*slug)
}
