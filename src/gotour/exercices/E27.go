package exercices

import "fmt"

type Vertex27 struct {
    X, Y int
}

var (
    p = Vertex27{1, 2}  // has type Vertex
    q = &Vertex27{1, 2}   // has type *Vertex
    r = Vertex27{X: 1}    // Y:0 is implicit 
    s = Vertex27{}        // X:0 and Y:0
)

func E27() {
    fmt.Println(p, q, r, s)
}
