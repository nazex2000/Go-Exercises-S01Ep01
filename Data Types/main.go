package main;
import (
	"fmt"
	"strings"
);

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

func stringsType(){
	var frase string= "Its a Good Day In Milan, But i miss Tofo!"
	fmt.Print(frase)
}

func stringsOperations(){
	str := "  Go is a great programming language!  "
    substr := "great"

    // 1. Compare
    fmt.Println("1. Compare:")
    fmt.Println("Compare 'Go' and 'Go':", strings.Compare("Go", "Go")) // 0
    fmt.Println("Compare 'Go' and 'Golang':", strings.Compare("Go", "Golang")) // > 0

    // 2. Contains
    fmt.Println("\n2. Contains:")
    fmt.Println("Contains 'great':", strings.Contains(str, substr)) // true

    // 3. ContainsAny
    fmt.Println("\n3. ContainsAny:")
    fmt.Println("Contains any of 'aeiou':", strings.ContainsAny(str, "aeiou")) // true

    // 4. ContainsRune
    fmt.Println("\n4. ContainsRune:")
    fmt.Println("Contains rune 'g':", strings.ContainsRune(str, 'g')) // true

    // 5. Count
    fmt.Println("\n5. Count:")
    fmt.Println("Count of 'e':", strings.Count("cheese", "e")) // 3

    // 6. HasPrefix
    fmt.Println("\n6. HasPrefix:")
    fmt.Println("Starts with '  Go':", strings.HasPrefix(str, "  Go")) // true

    // 7. HasSuffix
    fmt.Println("\n7. HasSuffix:")
    fmt.Println("Ends with 'language!  ':", strings.HasSuffix(str, "language!  ")) // true

    // 8. Index
    fmt.Println("\n8. Index:")
    fmt.Println("Index of 'great':", strings.Index(str, substr)) // 10

    // 9. IndexAny
    fmt.Println("\n9. IndexAny:")
    fmt.Println("Index of any 'aeiou':", strings.IndexAny(str, "aeiou")) // 3

    // 10. IndexRune
    fmt.Println("\n10. IndexRune:")
    fmt.Println("Index of rune 'g':", strings.IndexRune(str, 'g')) // 10

    // 11. Join
    fmt.Println("\n11. Join:")
    words := []string{"Go", "is", "fun"}
    fmt.Println(strings.Join(words, " ")) // "Go is fun"

    // 12. Repeat
    fmt.Println("\n12. Repeat:")
    fmt.Println(strings.Repeat("na", 2)) // "nana"

    // 13. Replace
    fmt.Println("\n13. Replace:")
    replaced := strings.Replace(str, substr, "awesome", 1)
    fmt.Println("Replace 'great' with 'awesome':", replaced) // "  Go is a awesome programming language!  "

    // 14. Split
    fmt.Println("\n14. Split:")
    fmt.Println(strings.Split(str, " ")) // ["", "", "Go", "is", "a", "great", "programming", "language!", "", "", ""]

    // 15. SplitN
    fmt.Println("\n15. SplitN:")
    fmt.Println(strings.SplitN("a,b,c", ",", 2)) // ["a", "b,c"]

    // 16. SplitAfter
    fmt.Println("\n16. SplitAfter:")
    fmt.Println(strings.SplitAfter("a,b,c", ",")) // ["a,", "b,", "c"]

    // 17. Title
    fmt.Println("\n17. Title:")
    fmt.Println(strings.Title("go language")) // "Go Language"

    // 18. ToLower
    fmt.Println("\n18. ToLower:")
    fmt.Println(strings.ToLower("GOLANG")) // "golang"

    // 19. ToUpper
    fmt.Println("\n19. ToUpper:")
    fmt.Println(strings.ToUpper("golang")) // "GOLANG"

    // 20. Trim
    fmt.Println("\n20. Trim:")
    fmt.Println("Trimmed:", strings.TrimSpace(str)) // "Go is a great programming language!"

    // 21. TrimPrefix
    fmt.Println("\n21. TrimPrefix:")
    fmt.Println("Trimmed Prefix '  Go':", strings.TrimPrefix(str, "  Go")) // " is a great programming language!  "

    // 22. TrimSuffix
    fmt.Println("\n22. TrimSuffix:")
    fmt.Println("Trimmed Suffix 'language!  ':", strings.TrimSuffix(str, "language!  ")) // "  Go is a great programming "

    // 23. Fields
    fmt.Println("\n23. Fields:")
    fmt.Println(strings.Fields("Go is great")) // ["Go", "is", "great"]

    // 24. FieldsFunc
    fmt.Println("\n24. FieldsFunc:")
    f := func(c rune) bool {
        return c == 'a'
    }
    fmt.Println(strings.FieldsFunc("banana", f)) // ["b", "n", "n"]

}

func main(){
	stringsOperations()
}