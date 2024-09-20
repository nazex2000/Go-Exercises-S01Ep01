package concepts;
import "fmt";

func General(){
	var x int = 1;
	var y int;
	var p *int;
	p = &x;
	y = *p;

	fmt.Printf("The value of p is %d\n The value of y is %d",*p, y)
}