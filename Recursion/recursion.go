//Find the factorial of a given number
package main;

import "fmt";

func Factorial(num int) int{
	if (num <= 1){
		return 1;
	}
	return num * Factorial(num-1);
}

func main(){
	var num, fac int;
	fmt.Print("====================Factorial Calculator======================\n");
	fmt.Print("Insert a Number :: ");
	fmt.Scanf("%d", &num);
	fac = Factorial(num)
	fmt.Printf("The Factorial of %d is %d",num, fac);
}

