package hello

const englishPrefix = "hello "
const spanishPrefix = "hola "
const frenchPrefix = "bonjue "

func getString(name string, lang string) string {
	var localName = name
	if name == "" {
		localName = "world"
	}

	return getPrefix(lang) + localName
}

func getPrefix(lang string) string {
	switch lang {
	case "english":
		return englishPrefix
	case "french":
		return frenchPrefix
	case "spanish":
		return spanishPrefix
	default:
		return englishPrefix
	}
}