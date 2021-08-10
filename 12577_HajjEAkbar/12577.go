// title: Hajj-e-Akbar
// filename: 12577.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
    h := map[string]string {
        "Hajj":"Hajj-e-Akbar",
        "Umrah":"Hajj-e-Asghar",
    }
    c := 0
    for true {
        var word string
        fmt.Scan(&word)
        if word == "*" { break }
        c++
        fmt.Printf("Case %d: %s\n", c, h[word])
    }
}
