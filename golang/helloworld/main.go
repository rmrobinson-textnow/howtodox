package main

import (
	"flag"
	"fmt"
)

var (
	name = flag.String("name", "", "The name to print")
)

func main() {
	flag.Parse()

	fmt.Printf("Hello world %s\n", *name)
}
