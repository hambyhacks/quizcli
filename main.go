package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problems struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", ".csv file that is in the format of 'question,answer'.")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		handleError(fmt.Sprintf("Failed to open the .csv file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		handleError("Failed to parse the provided .csv file.")
	}

	correct := 0
	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []Problems {
	ret := make([]Problems, len(lines))
	for i, line := range lines {
		ret[i] = Problems{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func handleError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
