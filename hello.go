package main

import (
	"fmt"
	"log"

	greetings "github.com/ncor/gl-em1"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("ensi")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
