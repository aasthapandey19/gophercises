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
	problems := parseLines(lines)
	for i, prob := range problems {
		fmt.Printf("Problem # %d : %s = \n", i+1, prob.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == prob.a {
			fmt.Println("Correct!")
		}

	}
}

func parseLines(lines [][]string) []problem {
	returnSlice := make([]problem, len(lines))

	for i, line := range lines {
		returnSlice[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return returnSlice
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
