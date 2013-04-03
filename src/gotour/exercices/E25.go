package exercices

import "fmt"

type Vertex25 struct{
    X int
    Y int
}

func E25() {
    v := Vertex25{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
