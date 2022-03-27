package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problems struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", ".csv file that is in the format of 'question,answer'.")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds.")
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

	// Initialize problems and randomize the order using rand.Shuffle()
	problems := parseLines(lines)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})

	// Implement timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
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
