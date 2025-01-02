package main

import (
	"errors"
	"fmt"
)

func main() {
	q, err := divide(10, 50)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Erro: ", err)
	fmt.Println(q)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("DivisionByZeroError")
	}
	return a / b, nil
}
