// title: Hasmat The Brave Warrior
// filename: 10055.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var x, y int
    _, err := fmt.Scan(&x, &y)
    for err == nil {
        if x > y {
            println(x - y)
        } else {
            println(y - x)
        }
        _, err = fmt.Scan(&x, &y)
    }
}
