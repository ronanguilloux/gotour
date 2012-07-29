package exercices

import(
    "fmt"
    "math"
)

func pow21(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func E21() {
    fmt.Println(
        pow21(3, 2, 10),
        pow21(3, 3, 20),
    )
}
