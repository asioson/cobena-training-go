// title: Odd Sum
// filename: 10783.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var n, a, b, sum int
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        fmt.Scan(&a, &b)
        sum = 0
        for x := a; x <= b; x++ {
            if x%2 == 1 {
                sum += x
            }
        }
        fmt.Printf("Case %d: %d\n", i, sum)
    }
}
