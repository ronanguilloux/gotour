package exercices
    
import "fmt"

// see http://golang.org/doc/articles/gos_declaration_syntax.html
func add(x int, y int) int {
    return x + y
}

func E06() {
    fmt.Println(add(42, 13))
}