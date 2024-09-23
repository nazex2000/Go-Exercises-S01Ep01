// Find the factorial of a given number
package main

import (
	"fmt"
)

func PrintInt(num int) {
	conversion := "0123456789ABCDEF"
	base := 16
	digit := num % base
	num = num / base
	if num != 0 {
		PrintInt(num)
	}
	fmt.Print(string(conversion[digit]))
}

func main() {
	var num int
	fmt.Print("====================Conversion Decimal to Hexadecimal======================\n")
	fmt.Print("Insert a Decimal Number :: ")
	fmt.Scanf("%d", &num)
	fmt.Scanf("Conversion to hexadecimal is :: ")
	PrintInt(num)
}
