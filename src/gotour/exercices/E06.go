package exercices

import "fmt"

// see http://golang.org/doc/articles/gos_declaration_syntax.html
// appending 06 to avoid redeclaration error
func add06(x int, y int) int {
    return x + y
}

func E06() {
    fmt.Println(add06(42, 13))
}
