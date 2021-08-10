// title: Packing for Holiday
// filename: 12372.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var cases int
    fmt.Scan(&cases)
    for c := 0; c < cases; c++ {
        var L, W, H int
        fmt.Scan(&L, &W, &H)
        verdict := "good"
        if (L > 20) || (W > 20) || (H > 20) {
            verdict = "bad"
        }
        fmt.Printf("Case %d: %s\n", c+1, verdict)
    }
}
