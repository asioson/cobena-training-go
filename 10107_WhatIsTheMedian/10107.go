// title: What is the Median?
// filename: 10107.go
// allan.sioson@gmail.com

package main

import (
    "fmt"
    "sort"
)

func main() {
    var x []int
    var a, n int
    _, err := fmt.Scan(&a)
    for err == nil {
        x = append(x, a)
        sort.Ints(x)
        n = len(x)
        if n % 2 == 1 {
            println(x[n >> 1])
        } else {
            println((x[(n >> 1) - 1] + x[(n >> 1)]) / 2)
        }
        _, err = fmt.Scan(&a)
    }
}
