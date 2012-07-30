package exercices

import "fmt"

type Vertex28 struct {
    X, Y int
}

func E28() {
    v := new(Vertex28)
    fmt.Println(v)
    v.X, v.Y = 11, 9
    fmt.Println(v)
}
