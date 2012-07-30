package exercices

import "fmt"

type Vertex30 struct {
    Lat, Long float64
}

var m30  = map[string] Vertex30 {
    "Bell Labs": Vertex30{
        40.68433, -74.39967,
    },
    "Google": Vertex30{
        37.42202, -122.08408,
    },
}

func E30() {
    fmt.Println(m30)
}
