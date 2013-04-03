package exercices

import "fmt"

type Vertex29 struct {
    Lat, Long float64
}

var m29 map[string]Vertex29

func E29() {
    m29 = make(map[string]Vertex29)
    m29["Bell Labs"] = Vertex29{
        40.68433, 74.39967,
    }
    fmt.Println(m29["Bell Labs"])
}
