package exercices

import(
    "fmt"
    "math"
)

func pow22(x, n , lim float64) float64 {
    if v:= math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func E22() {
    fmt.Println(
        pow22(3, 2, 10),
        pow22(3, 3, 20),
    )
}
