package main

import "fmt"

const (
	first = iota
	second 
)

const (
	third = iota + 3
	fourth
)

func types() {
	var myInt int = 10 
	fmt.Println("myInt",myInt)

	var myFloat float32 = 3.14 
	fmt.Println("myFloat", myFloat)

	myString := "João"
	fmt.Println("myString", myString)

	var myBoolean bool = true 
	fmt.Println("myBoolean", myBoolean)

	var myPointer *string = new(string)
	*myPointer = "Pointer"
	fmt.Println("myPointer", *myPointer)

	myStringReference := &myString
	fmt.Println("myPointerReference", *myStringReference)

	myString = "Novo João"

	fmt.Println("myPointerReference", *myStringReference)

	const myConst = 3.14 
	fmt.Println(myConst)

	fmt.Println("first, second, third, fourth", first, second, third, fourth)
}