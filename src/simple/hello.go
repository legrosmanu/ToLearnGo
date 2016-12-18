package main

import (
	"fmt"
	"strconv"
)

func multipleVarReturn(first int, second float64) (label string, result string) {
	label = "The paramaters are : "
	result = strconv.Itoa(first) + " " + strconv.FormatFloat(second, 'E', -1, 64)
	return
}

func main() {
	fmt.Println("Hello, GO")
	firstReturn, secondeReturn := multipleVarReturn(10, 3.14)
	fmt.Println(firstReturn + secondeReturn)
}
