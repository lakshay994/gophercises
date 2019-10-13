package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("csv", "", ".csv file for creating the quiz")
	flag.Parse()
	generateQuiz(csvFile)
}

func generateQuiz(filename *string) {
	file, err := os.Open(*filename)
	if err != nil {
		exit("No such file found")
	}

	reader := csv.NewReader(file)
	content, err := reader.ReadAll()
	if err != nil {
		exit("Errors while reading file")
	}

	problems := parseFileContent(content)

	correctAnswers := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d. %s \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correctAnswers++
		}
	}

	fmt.Printf("You got %d out of %d answers right.\n", correctAnswers, len(problems))
}

func parseFileContent(content [][]string) []Content {

	parsed := make([]Content, len(content))
	for i, line := range content {
		parsed[i] = Content{
			question: line[0],
			answer:   line[1],
		}
	}

	return parsed
}

type Content struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
