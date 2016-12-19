package main

import (
	"fmt"
)

// MyType : Just to test struct
type MyType struct {
	label string
}

func (myType MyType) setLabelWithoutPointer(newLabel string) {
	myType.label = newLabel
}

func (myType *MyType) setLabelWithPointer(newLabel string) {
	myType.label = newLabel
}

func main() {
	initValue := MyType{label: "On init"}
	fmt.Println(initValue.label)
	initValue.setLabelWithoutPointer("New label")
	fmt.Println("After setLabelWithoutPointer call ", initValue.label)
	initValue.setLabelWithPointer("Last label")
	fmt.Println("After setLabelWithPointer call ", initValue.label)
}
