package exercices
    
import "fmt"

// add should not be redeclared
func add2(x, y int) int {
    return x + y
}

func E7() {
    fmt.Println(add2(42, 13))
}