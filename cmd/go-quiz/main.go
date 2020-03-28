package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tobiade/go-exercises/quiz"
)

func main() {
	filePathPtr := flag.String("path", "problems.csv", "Where your quiz lives")
	duration := flag.Duration("duration", 30*time.Second, "Time limit for the quiz - to be specified in a parsable golang Duration.")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you ready to start the quiz? Enter 'Y' - you literally have no other option here: ")
	r, _, err := reader.ReadRune()
	if err != nil {
		return
	}
	if r == 'Y' {
		quiz.Run(*filePathPtr, *duration)
	}
}
