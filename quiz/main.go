package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "./problems.csv", "Path to csv file")
	flag.Parse()

	playQuiz(*csvFile)
}

func playQuiz(filePath string) {
	// Bonus: see how to handle a user exit
	questionsAsked, correctAnswers := 0, 0

	file, err := os.Open(filePath)
	if err != nil {
		exit(fmt.Sprintf("Error opening file: %q", err))
	}

	defer file.Close()

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse csv file: %v", err))
	}

	problems := parseLines(lines)

	for _, problem := range problems {
		fmt.Printf("%s: ", problem.question)
		questionsAsked += 1

		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == problem.answer {
			correctAnswers += 1
		}
	}

	fmt.Printf("Quiz Done!\n problems: %d\n correct answers: %d\n", questionsAsked, correctAnswers)
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
