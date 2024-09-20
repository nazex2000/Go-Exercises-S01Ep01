package main;
import "fmt";

func numbers(){
	//Integers
	//Theres different lengths to declare int: int8, int16, int32, int64, unit8,...unit 64
	//We can perform different types of operations such arithmetic, Comparasion and Boolean
	var x int = 1; //Declaration of variable num with type int and value 10;
	fmt.Printf("%d\n", x)
	//Float
	//Theres different lengths to declare int: float32~6 digits precision, float64~15 digits precision
	var y float32 = 45.0
	fmt.Printf("%.2f", y)

}

func numbersConversion(){
	var a int = 42
    var b float64 = float64(a)
    fmt.Printf("Integer to Float: %d -> %.2f\n", a, b)

    // Float to integer conversion (truncation)
    var c float64 = 42.89
    var d int = int(c)
    fmt.Printf("Float to Integer: %.2f -> %d\n", c, d)

    // Integer type conversion
    var e int32 = 1000
    var f int64 = int64(e)
    fmt.Printf("int32 to int64: %d -> %d\n", e, f)

    // Unsigned integer conversion
    var g int = -10
    var h uint = uint(g) // Be careful! Negative to unsigned can cause issues.
    fmt.Printf("Signed to Unsigned: %d -> %d\n", g, h)

    // Float32 to Float64 conversion
    var i float32 = 3.14
    var j float64 = float64(i)
    fmt.Printf("float32 to float64: %.2f -> %.2f\n", i, j)
}

func main(){
	numbers()
}