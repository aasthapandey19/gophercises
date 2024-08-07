package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "csv file for ques/ans structure")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Filename not found : %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error in parsing csv file: %s\n", *csvFilename))
	}
	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
