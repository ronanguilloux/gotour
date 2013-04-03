package exercices

import "fmt"

// add should not be redeclared
func add07(x, y int) int {
    return x + y
}

func E07() {
    fmt.Println(add07(42, 13))
}
