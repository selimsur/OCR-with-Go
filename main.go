package main

import (
	"fmt"
	"log"

	"github.com/otiai10/gosseract/v2"
)

func main() {

	client := gosseract.NewClient()
	defer client.Close()
	err := client.SetImage("example.png") // Photo path
	if err != nil {
		log.Fatal(err)
	}

	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Text: ")
	fmt.Println(text)
}
