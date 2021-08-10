// title: Back to High School Physics
// filename: 10071.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
	var v, t int
	_, err := fmt.Scan(&v, &t)
	for err == nil {
		println(2 * v * t)
		_, err = fmt.Scan(&v, &t)
	}
}
