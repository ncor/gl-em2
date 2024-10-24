package main

import (
	"fmt"

	greetings "github.com/ncor/gl-em1"
)

func main() {
	message := greetings.Hello("ensi")
	fmt.Println(message)
}
