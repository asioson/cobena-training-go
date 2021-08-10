// title: Is This The Easiest Problem?
// filename: 11479.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var cases int
    fmt.Scan(&cases)
    for k := 0; k < cases; k++ {
        var a, b, c int
        fmt.Scan(&a, &b, &c)
        if a > b { a, b = b, a }
        if a > c { a, c = c, a }
        if b > c { b, c = c, b }
        if a < 0 {
            fmt.Printf("Case %d: Invalid\n", k+1)
        } else if (a == b) && (b == c) {
            fmt.Printf("Case %d: Equilateral\n", k+1)
        } else if a + b <= c {
            fmt.Printf("Case %d: Invalid\n", k+1)
        } else if (a == b) || (b == c) {
            fmt.Printf("Case %d: Isosceles\n", k+1)
        } else {
            fmt.Printf("Case %d: Scalene\n", k+1)
        }
    }
}
