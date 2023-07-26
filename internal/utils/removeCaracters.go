package utils

import "regexp"

func RemoveCaracteresEspeciais(text string) string {
	mapaCaracteres := map[string]*regexp.Regexp{
		"a": regexp.MustCompile("[áàãâä]"),
		"e": regexp.MustCompile("[éèêë]"),
		"i": regexp.MustCompile("[íìîï]"),
		"o": regexp.MustCompile("[óòõôö]"),
		"u": regexp.MustCompile("[úùûü]"),
		"c": regexp.MustCompile("[ç]"),
		"":  regexp.MustCompile("[?]"),
	}

	for chave, regex := range mapaCaracteres {
		text = regex.ReplaceAllString(text, chave)
	}

	return text
}