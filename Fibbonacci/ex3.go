package ex3

import "fmt"

//print the first 50 fibbonacci numbers

func printFibo(x int, y int, counter int) {
	if counter == 0 {
		return
	}
	counter -= 1
	fmt.Printf("%d\n", y)
	aux := x + y
	x = y
	printFibo(x, aux, counter)
}

func Exercise() {
	printFibo(0, 1, 10)
}
