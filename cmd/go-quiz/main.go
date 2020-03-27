package main

import (
	"flag"

	"github.com/tobiade/go-exercises/quiz"
)

func main() {
	filePathPtr := flag.String("quizPath", "problems.csv", "Where your quiz lives")
	flag.Parse()
	quiz.Run(*filePathPtr)
}
