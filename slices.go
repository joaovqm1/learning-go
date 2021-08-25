package main

import "fmt"

func slices() {
	slice := []int{1,2,3} 

	fmt.Println(slice)

	slice = append(slice, 4, 5)

	fmt.Println(slice)

	copy := slice[1:]

	fmt.Println(copy)

}