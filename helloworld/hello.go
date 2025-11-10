package helloworld

const (
	french  = "french"
	spanish = "spanish"
	english = "english"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	switch language {
	case "spanish":
		return spanishPrefix + name
	case "english":
		return englishPrefix + name
	case "french":
		return frenchPrefix + name
	}

	return englishPrefix + name
}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	default:
		return englishPrefix
	}
}
