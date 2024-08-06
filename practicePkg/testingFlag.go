package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String("name", " defualt name", "helper exp")
	age := flag.Int("age", 0, "helper age")
	verbose := flag.Bool("verbose", false, "enable verbose function")

	flag.Parse()

	if *verbose {
		fmt.Println("verbose mode enabled")
	}
	fmt.Printf("Your name is %s and age is %d", *name, *age)
}
