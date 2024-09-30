package main

import "fmt"

func main() {
	//Declaration an Array
	//var arr [5]int
	//second way to delcare
	arrsec := [5]int{1, 2, 3, 4, 5}

	//implicit decalration, the [...] allows Go to infer the size
	//arr := [...]int{1,2,3}

	//Accessing Array Elements
	//fmt.Println(arrsec[0])

	//Iterating Trough Arrays
	for i, v := range arrsec {
		fmt.Printf("ind %d, val %d", i, v)
	}
}
