package main

import (
	"fmt"

	"github.com/lpallage/toolkit"
)

func main() {
	// Create an instance of Tools struct from the toolkit package.
	tools := toolkit.Tools{}

	// Generate a random string of length 10.
	randomString := tools.RandomString(10)

	// Print the generated random string to the console.
	fmt.Println("Random String:", randomString)
}
