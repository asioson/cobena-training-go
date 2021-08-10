// title: Horror Dash
// filename: 11799.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var cases int;
    fmt.Scan(&cases)
    for c := 0; c < cases; c++ {
        var n, answer int
        fmt.Scan(&n)
        if n > 0 {
            fmt.Scan(&answer)
            for i := 1; i < n; i++ {
                var x int
                fmt.Scan(&x)
                if x > answer { answer = x }
            }
        }
        fmt.Printf("Case %d: %d\n", c+1, answer)
    }
}
