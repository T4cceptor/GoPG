package greetings

import "math/rand"

func RandomEntry[S ~[]E, E any](s S) E {
	// Returns a random entry of the given slice
	return s[rand.Intn(len(s))]
}

func randomEntryTest[S ~[]E, E any](s S) E {
	return RandomEntry(s)
}
