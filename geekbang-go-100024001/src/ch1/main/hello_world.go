package main

import (
	"fmt"
	"os"
)

func main() {
	// go run hello_world.go go
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world!", os.Args[1])
	}
	// os.Exit(-1)
}
