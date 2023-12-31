package helloworld

const englishHelloPrefix = "Hello, "

//	func Hello(name string) string {
//		if name == "" {
//			name = "World"
//		}
//		return englishHelloPrefix + name
//	}
// func Hello(name string, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	return englishHelloPrefix + name
// }

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
