// title: Cost Cutting
// filename: 11727.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var cases int
    fmt.Scan(&cases)
    for c := 0; c < cases; c++ {
        var x, y, z int
        fmt.Scan(&x, &y, &z)
        if x > y { x, y = y, x }
        if x > z { x, z = z, x }
        if y > z { y, z = z, y }
        fmt.Printf("Case %d: %d\n", c+1, y)
    }
}
