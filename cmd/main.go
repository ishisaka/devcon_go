package main

import (
	"devcon_go/pkg/hello"
	"fmt"
)

func main() {
	message := hello.SayHello()
	fmt.Println(message)
}
