// title: Searching for Nessy
// filename: 11044.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var cases int
    fmt.Scan(&cases)
    for c := 0; c < cases; c++ {
        var n, m, a, b int
        fmt.Scan(&n, &m)
        a = n / 3
        b = m / 3
        println(a * b)
    }
}
