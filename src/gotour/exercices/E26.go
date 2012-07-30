package exercices

import "fmt"

type Vertex26 struct {
    X int
    Y int
}

func E26() {
    p := Vertex26{1, 2}
    q := &p
    q.X = 1e9
    fmt.Println(p)
}
