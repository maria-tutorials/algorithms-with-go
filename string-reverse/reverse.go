package string

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	w := []rune(word)

	l := len(w)

	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}

	return string(w)
}
