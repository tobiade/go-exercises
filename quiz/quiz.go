package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type quizEntry struct {
	question string
	answer   string
}

//Run the quuiz program.
func Run(filePath string) {
	quiz, err := getQuiz(filePath)
	if err != nil {
		return
	}
	var correct int
	var wrong int
	fmt.Println("Welcome to the most amazing quiz ever!")
	reader := bufio.NewReader(os.Stdin)
	for _, entry := range quiz {
		fmt.Printf("%s: ", entry.question)
		userAnswer, _ := reader.ReadString('\n')
		if strings.TrimSpace(userAnswer) == entry.answer {
			correct++
		} else {
			wrong++
		}
	}
	fmt.Printf("Correct answers: %d, Wrong answers: %d\n", correct, wrong)
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
