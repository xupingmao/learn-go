
package main

import "fmt"

func main() {
	const LENGTH = 10

	const NAME string = "Name"

	// ./002-constant.go:9:9: cannot assign to LENGTH
	// LENGTH = 20

	fmt.Printf("LENGTH = %d\n", LENGTH)
	fmt.Printf("NAME   = %s\n", NAME)
}