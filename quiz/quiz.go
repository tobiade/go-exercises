package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type quizEntry struct {
	question string
	answer   string
}

//Run the quiz program.
func Run(filePath string, d time.Duration) {
	quiz, err := getQuiz(filePath)
	if err != nil {
		return
	}
	c := make(chan int)
	go runQuiz(quiz, c)

	var correctAnswers int
	timeHasRunOut := false
	for {
		select {
		case correctAnswers = <-c:
		case <-time.After(d):
			timeHasRunOut = true
			fmt.Println("\nYou ran out of time!")
		}
		if correctAnswers == len(quiz) || timeHasRunOut {
			break
		}
	}
	fmt.Printf("Correct answers: %d, Total number of questions: %d\n", correctAnswers, len(quiz))
}

func runQuiz(quiz []quizEntry, c chan int) {
	var correct int
	fmt.Println("Welcome to the most amazing quiz ever!")
	reader := bufio.NewReader(os.Stdin)
	for _, entry := range quiz {
		fmt.Printf("%s: ", entry.question)
		userAnswer, _ := reader.ReadString('\n')
		if strings.TrimSpace(userAnswer) == entry.answer {
			correct++
		}
		c <- correct
	}
}

func getQuiz(filePath string) ([]quizEntry, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)

	entries, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	var quiz []quizEntry
	for _, entry := range entries {
		quiz = append(quiz, quizEntry{entry[0], entry[1]})
	}
	return quiz, nil
}
