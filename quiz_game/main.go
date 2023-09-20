package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitTimer := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open %s file\n", *csvFilename)
		os.Exit(1)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Erro parse CSV file.")
	}

	problems := parseProblems(lines)

	timer := time.NewTimer(time.Duration(*limitTimer) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d => %s= ", i, p.question)
		answerChannel := make(chan string)
		go func() {
			var answerUser string
			fmt.Scanf("%s", &answerUser)
			answerChannel <- answerUser
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYour scored %d of %d\n", correct, len(problems))
			return
		case answerUser := <-answerChannel:
			if answerUser == p.answer {
				correct++
			}
		}
	}
	fmt.Printf("Your scored %d of %d\n", correct, len(problems))
}

func parseProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, p := range lines {
		problems[i] = problem{
			question: p[0],
			answer:   strings.TrimSpace(p[1]),
		}
	}
	return problems
}
