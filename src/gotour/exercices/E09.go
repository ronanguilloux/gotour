package exercices

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4/9
    y = sum - x
    return
}

func E09() {
    fmt.Println(split(17))
}
