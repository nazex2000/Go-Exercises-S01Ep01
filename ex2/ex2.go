package ex2

import "fmt"

//Factorial of n

func factorialCalc(val int) int {
	var num = val - 1
	if num == 0 {
		return val
	}
	return val * factorialCalc(num)
}

func Exercise() {
	num := 5
	facNum := factorialCalc(num)
	fmt.Printf("The factorial of %d is %d", num, facNum)
}
