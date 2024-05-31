package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		if len(character) > 1 {
			repeated += " " + character
		} else {
			repeated += character
		}
	}
	return repeated
}