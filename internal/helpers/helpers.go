package helpers

import "github.com/dimuska139/urlshortener/internal/constants"

func GenerateShortcode(id int) string {
	base := len(constants.ShortcodeAlphabet)

	var encoded string
	for id > 0 {
		encoded += string(constants.ShortcodeAlphabet[id%base])
		id = id / base
	}

	var reversed string
	for _, v := range encoded {
		reversed = string(v) + reversed
	}

	return reversed
}
