package main

import (
	"fmt"
	"time"
)

func main() {
	numero := lerInt()
	fmt.Printf("\n%d! = %d\n\n", numero, fatorial(numero))

	// for range
	frutas := []string{"maça", "banana", "laranja"}
	for i, fruta := range frutas {
		fmt.Printf("%d: %s\n", i+1, fruta)
	}

	var n int = 1
	fmt.Println()
	for n <= 10 {
		fmt.Printf("%d ", n)
		time.Sleep(time.Second)
		n++
	}
	fmt.Printf("\n\n")
}

func lerInt() int {
	fmt.Println()
	var num int
	for {
		fmt.Print("Digite um nº inteiro [>=2]: ")
		fmt.Scan(&num)
		if num >= 2 {
			break
		}
	}
	return num
}

func fatorial(n int) int {
	fat := n
	for i := 1; i < n; i++ {
		fat *= i
	}
	return fat
}
