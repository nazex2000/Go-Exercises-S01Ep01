package ex1;
import "fmt";

//Working with Pointers:
//Create a function that takes a pointer to an integer and increments the value of that integer by 10. Test the function with different variables.

func incrementBy10(ptr *int) *int{
	*ptr+=10 
	return ptr
}

func Exercise(){
	val:=20

	ptr:=incrementBy10(&val)
	fmt.Printf("The incremented value is %d\n", *ptr)

}