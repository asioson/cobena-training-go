// title: Division of Nlognia
// filename: 11498.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var K, ox, oy, locx, locy int
    for true {
        fmt.Scan(&K)
        if K == 0 { break }
        fmt.Scan(&ox, &oy)
        for k := 0; k < K; k++ {
            fmt.Scan(&locx, &locy)
            x := locx - ox
            y := locy - oy
            if (x == 0) || (y == 0) {
                println("divisia")
            } else if (x > 0) && (y > 0) {
                println("NE")
            } else if (x > 0) && (y < 0) {
                println("SE")
            } else if (x < 0) && (y > 0) {
                println("NO")
            } else {
                println("SO")
            }
        }
    }
}
