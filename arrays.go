package main

import "fmt"

func arrays() {
	var arr1 [2]int 
	arr1[0] = 1
	arr1[1] = 2

	fmt.Println(arr1)

	arr2 := [3]int{1,2,3}
	fmt.Println(arr2)

}