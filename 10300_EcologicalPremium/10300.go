// title: Ecological Premium
// filename: 10300.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var n, f, total, farmSize, count, score int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        total = 0
        fmt.Scan(&f)
        for j := 0; j < f; j++ {
            fmt.Scan(&farmSize, &count, &score)
            total += farmSize * score
        }
        println(total)
    }
}
