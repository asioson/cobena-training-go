// title: Zapping
// filename: 12468.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var a, b, d1, d2 int
    for true {
        fmt.Scan(&a, &b)
        if (a < 0) && (b < 0) { break }
        if a < b { a, b = b, a }
        d1 = a - b
        d2 = b + 100 - a
        if d1 < d2 {
            println(d1)
        } else {
            println(d2)
        }
    }
}
