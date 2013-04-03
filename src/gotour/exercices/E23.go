package exercices

import(
    "math/cmplx"
    "fmt"
)

var (
   ToBe bool = false
   MaxInt uint64 = 1<<64 - 1
   z23 complex128 = cmplx.Sqrt(-5+12i)
)

func E23() {
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z23, z23)
}