package main

/*********************************************************************************
* Just a hello world, with a try of the function which have 2 variables on return.
**********************************************************************************/

import (
	"fmt"
	"strconv"
)

func multipleVarReturn(first int, second float64) (label string, result string) {
	label = "The paramaters are : "
	result = strconv.Itoa(first) + " and " + strconv.FormatFloat(second, 'E', -1, 64)
	return
}

func main() {
	fmt.Println("Hello, GO")
	firstReturn, secondeReturn := multipleVarReturn(10, 3.14)
	fmt.Println(firstReturn, secondeReturn)
}
