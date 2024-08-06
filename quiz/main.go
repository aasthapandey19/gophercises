package main

import (
	"flag"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "csv file for ques/ans structure")
	flag.Parse()

	_ = csvFilename

}
