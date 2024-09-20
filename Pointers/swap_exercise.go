package main;
import "fmt";

func swapping(x *int, y *int){
	aux:=*x;
	*x = *y;
	*y = aux;
}

func main(){
	var val1 int = 4;
	var val2 int = 5;
	fmt.Println("---Before Swapping---")
	fmt.Printf("Value 1 is %d and Value 2 is %d",val1, val2);
	swapping(&val1, &val2);
	fmt.Println("\n---After Swapping---")
	fmt.Printf("Value 1 is %d and Value 2 is %d",val1, val2);

}