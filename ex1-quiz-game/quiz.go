package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}

	problems := parseLines(lines)

	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", index+1, problem.question)
		var answer string
		fmt.Scanf("%s/n", &answer)
		fmt.Printf("You answer is: %v\n", answer)
	}

	// fmt.Println(lines)
	// fmt.Println(problems)
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return ret
}

type Problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
