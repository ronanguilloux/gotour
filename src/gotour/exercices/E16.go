package exercices

import "fmt"

func E16() {
    sum := 1
    for ; sum < 1000; {
        sum += sum;
    }
    fmt.Println(sum)
}