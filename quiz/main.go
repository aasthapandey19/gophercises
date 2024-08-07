package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "csv file for ques/ans structure")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Filename not found : %s", *csvFilename))
	}
	_ = file
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
