/* Utilizando alguns tipos de dados */

package main

import "fmt"

func main() {
	fmt.Printf("Type: %T - Value: %v\n", "tux", "tux")
	fmt.Printf("Type: %T - Value: %v\n", "10", "10")
	fmt.Printf("Type: %T - Value: %v\n", 100000000000, 100000000000)
	fmt.Printf("Type: %T - Value: %v\n", 1.23, 1.23)
	fmt.Printf("Type: %T - Value: %v\n", true, true)
}
