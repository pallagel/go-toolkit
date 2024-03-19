package toolkit

import (
	"crypto/rand"
	"math/big"
)

// randomStringSource defines the set of characters that can be used to generate a random string.
const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

//Tools in the type used to instantiate the module. Any variable of this
//type will have access to all the methods with the receiver *Tool

type Tools struct{}

// RandomString generates a random string of length n using characters from randomStringSource.
// It uses crypto/rand for secure random number generation.
func (t *Tools) RandomString(length int) string {
	sourceRunes := []rune(randomStringSource) // Convert the string source to a slice of runes for indexing.
	randomString := make([]rune, length)      // Create a slice of runes to hold the generated characters.

	for i := range randomString {
		// Generate a random index for selecting a character from the sourceRunes slice.
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(sourceRunes))))
		// Assign the selected character to the current position in the randomString slice.
		randomString[i] = sourceRunes[index.Int64()]
	}

	return string(randomString) // Convert the slice of runes back to a string and return.
}
