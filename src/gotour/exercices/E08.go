package exercices

import "fmt"

func swap(x,y string) (string, string) {
    return x, y
}

func E08() {
    a, b := swap("hello", "world")
    fmt.Println(a,b);
}
