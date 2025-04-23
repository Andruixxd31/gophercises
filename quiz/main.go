package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

type answers struct {
	question int
	correct  int
}

func main() {
	csvFile := flag.String("csv", "./problems.csv", "Path to csv file")
	limit := flag.Int("time", 30, "Time to run test for in seconds. Default is 30 seconds")
	flag.Parse()

	problems := readFile(*csvFile)

	done := make(chan bool)

	limitDuration := time.Duration(*limit) * time.Second
	timer := time.NewTimer(limitDuration)

	var ans = answers{}

	go playQuiz(problems, &ans, timer, done)
	select {
	case <-timer.C:
		fmt.Println("Timer done")
	case <-done:
		fmt.Println("Finished test")
	}

	printResults(&ans)
}

func playQuiz(problems []problem, ans *answers, timer *time.Timer, done chan bool) {
	for _, problem := range problems {
		fmt.Printf("%s: ", problem.question)
		ans.question++

		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == problem.answer {
			ans.correct++
		}
	}

	timer.Stop()
	done <- true
}

func readFile(filePath string) []problem {
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

	return parseLines(lines)
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

func printResults(ans *answers) {
	fmt.Printf("Quiz Done!\n problems: %d\n correct answers: %d\n", ans.question, ans.correct)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
