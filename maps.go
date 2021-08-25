package main

import "fmt"

func maps() {
	myMap := map[string]int{"foo":1}

	myMap["bar"] = 2

	fmt.Println(myMap)

	delete(myMap, "bar")

	fmt.Println(myMap)
}