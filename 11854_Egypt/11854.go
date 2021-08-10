// title: Egypt
// filename: 11854.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var x, y, z int
    _, err := fmt.Scan(&x, &y, &z)
    for err == nil {
        if (x == 0) && (y == 0) && (z == 0) { break }
        if x > y { x, y = y, x }
        if x > z { x, z = z, x }
        if y > z { y, z = z, y }
        if x * x + y * y == z * z {
            println("right")
        } else {
            println("wrong")
        }
        _, err = fmt.Scan(&x, &y, &z)
    }
}
