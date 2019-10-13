package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFile := flag.String("csv", "", ".csv file for creating the quiz")
	timeDuration := flag.Int("duration", 30, "time duration for the quiz")
	flag.Parse()
	generateQuiz(csvFile, timeDuration)
}

func generateQuiz(filename *string, duration *int) {
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

	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	correctAnswers := 0
	for i, problem := range problems {

		fmt.Printf("Problem %d. %s \n", i+1, problem.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTimes Up! \nYou got %d out of %d answers right\n", correctAnswers, len(problems))
			os.Exit(1)
		case answer := <-answerCh:
			if answer == problem.answer {
				correctAnswers++
			}
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
