package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const helpString string = `-cvs string 
	a csv file in format "question,answer". Default is "problems.csv".
-h 
	See help menu for flag options.
`

func main() {
	helpFlag := flag.Bool("h", false, "See flag options for program")
	csvFlag := flag.String("csv", "./problems.csv", "Path to csv file") // TODO: See where to validate path
	flag.Parse()

	if *helpFlag {
		printFlags()
		os.Exit(0)
	}

	playQuiz(*csvFlag)
}

func playQuiz(filePath string) error {
	// Bonus: see how to handle a user exit
	questionsAsked, correctAnswers := 0, 0

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file: %q", err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	for {
		problem, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("Error reading file: %q", err)
		}

		answer, err := strconv.Atoi(problem[1])
		if err != nil {
			return fmt.Errorf("Wrong answer format")
		}

		inputPrompt := bufio.NewReader(os.Stdin)
		fmt.Printf("%s: ", problem[0])
		text, _ := inputPrompt.ReadString('\n')

		questionsAsked += 1

		text = strings.TrimSpace(text)

		userAnswer, err := strconv.Atoi(text)
		if err != nil {
			return fmt.Errorf("Wrong answer format: %q", err)
		}

		if userAnswer == answer {
			correctAnswers += 1
		}
	}

	fmt.Printf("Quiz Done!\n problems: %d\n correct answers: %d\n", questionsAsked, correctAnswers)
	return nil
}

// Prints all flag options
func printFlags() {
	fmt.Print(helpString)
}
