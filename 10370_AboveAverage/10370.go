// title: Above Average
// filename: 10370.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var cases int
    fmt.Scan(&cases)
    for c := 0; c < cases; c++ {
        var total, count, n, a int
        var x []int

        total = 0
        fmt.Scan(&n)
        for i := 0; i < n; i++ {
            fmt.Scan(&a)
            x = append(x, a)
            total += a
        }
        count = 0
        for _, a = range x {
            if a * n > total { count++ }
        }
        fmt.Printf("%.3f%%\n", (float32(count) * 100.0) / float32(n))
    }
}
