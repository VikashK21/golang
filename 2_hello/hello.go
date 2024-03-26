package main

import "fmt"

const (
	spanish  = "Spanish"
	french   = "French"
	bhojpuri = "Bhojpuri"
	kannada  = "Kannada"
	japanese = "Japanese"

	englishHelloPrefix  = "Hello world, I am %v here!"
	spanishHelloPrefix  = "Hola mundo, ¡soy %v aquí!"
	frenchHelloPrefix   = "Bonjour tout le monde, je suis %v ici!"
	bhojpuriHelloPrefix = "नमस्ते दुनिया, हम %v यहाँ बा!"
	kannadaHelloPrefix  = "ಹಲೋ ವರ್ಲ್ಡ್, ನಾನು %v ಇಲ್ಲಿ!"
	japaneseHelloPrefix = "こんにちは世界、私はここで%vです!"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := fmt.Sprintf(englishHelloPrefix, name)
	switch language {
	case spanish:
		prefix = fmt.Sprintf(spanishHelloPrefix, name)
	case french:
		prefix = fmt.Sprintf(frenchHelloPrefix, name)
	case bhojpuri:
		prefix = fmt.Sprintf(bhojpuriHelloPrefix, name)
	case kannada:
		prefix = fmt.Sprintf(kannadaHelloPrefix, name)
	case japanese:
		prefix = fmt.Sprintf(japaneseHelloPrefix, name)
	}

	return prefix
}

func main() {
	fmt.Println(Hello("World", ""))
}
