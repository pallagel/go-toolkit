package toolkit

import "testing"

// It checks if the method correctly returns a string of the requested length.
func TestTools_RandomString(t *testing.T) {
	// Create an instance of the Tools struct to test its methods.
	var testTools Tools

	// Generate a random string of length 10 using the RandomString method.
	randomString := testTools.RandomString(10)

	// Verify that the length of the generated string is exactly 10.
	// If it's not, report an error.
	if len(randomString) != 10 {
		t.Error("Expected random string of length 10, but got a different length")
	}
}
