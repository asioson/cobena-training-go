// title: Language Detection
// filename: 12250.go
// allan.sioson@gmail.com

package main

import "fmt"

func main() {
	country := map[string]string {
        "HELLO": "ENGLISH",
		"HOLA": "SPANISH", 
        "HALLO": "GERMAN", 
        "BONJOUR": "FRENCH",
        "CIAO": "ITALIAN",
        "ZDRAVSTVUJTE": "RUSSIAN",
    }
	var mesg string
	count := 0
	_, err := fmt.Scan(&mesg)
	for err == nil {
		count++
		if lang, ok := country[mesg]; ok {
			fmt.Printf("CASE %d: %s\n", count, lang)
		} else if mesg == "#" {
            break
        } else {
			fmt.Printf("CASE %d: UNKNOWN\n", count)
		}
		_, err = fmt.Scan(&mesg)
	}
}
