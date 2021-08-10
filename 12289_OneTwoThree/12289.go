// title: One Two Three
// filename: 12289.go
// allan.sioson@gmail.com

package main

import "fmt"

func getValue(w string) int {
    if len(w) == 5 { return 3 }
    one := "one"
    mismatch := 0
    for i := 0; i < 3; i++ {
       if w[i] != one[i] { mismatch++ } 
       if mismatch > 1 { return 2 }
    }
    return 1
}

func main() {
    var word string
    var cases int
    fmt.Scan(&cases)
    for c := 0; c < cases; c++ {
        fmt.Scan(&word)
        println(getValue(word))
    }
}
