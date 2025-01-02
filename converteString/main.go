package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s := "3.1415926535a"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Nil: ", err)
	}

	fmt.Println(f, reflect.TypeOf(f))
}
