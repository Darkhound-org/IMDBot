// Input your api key from http://www.omdbapi.com/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var val string
	fmt.Println("Enter your api key: ")
	fmt.Scanln(&val)
	f, err := os.OpenFile("data", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(val)); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully written api key to data file..!")
}
