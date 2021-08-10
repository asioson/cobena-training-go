// title: Relational Operators
// filename: 11172.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    var cases int
    fmt.Scan(&cases)
    for c := 0; c < cases; c++ {
        var a, b int
        fmt.Scan(&a, &b)
        if a < b { 
            println("<") 
        } else if a > b { 
            println(">") 
        } else {
            println("=") 
        }
    }
}
