// Exercise 1: Quiz Game
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

// OpenCSV method
func OpenCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("<<%s>> file doesn't exist !", filename)
		return nil, err
	}

	quiz, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	if len(quiz) == 0 {
		return nil, errors.New("The quiz is empty !")
	}

	return quiz, nil
}

func processQuiz(quiz [][]string, counter *int, stop chan bool) {
	var answer string

	for index, line := range quiz {
		responseQuiz := line[1]
		fmt.Printf("Q%d: %s\n", index, line[0])
		fmt.Scanf("%s\n", &answer)

		if responseQuiz == answer {
			*counter++
		}
	}
	fmt.Println(*counter)
	stop <- true
}

func startQuiz(quiz [][]string, timer int64) {
	stop := make(chan bool)
	startC := 0
	counter := &startC

	go processQuiz(quiz, counter, stop)

	for {
		select {
		case <-stop:
			fmt.Printf("Result: %d/%d\n", *counter, len(quiz))
			return
		case <-time.After(time.Duration(timer) * time.Second):
			fmt.Printf("Result: %d/%d\n", *counter, len(quiz))
			return
		}
	}
}

func main() {
	filenameFlag := flag.String("filename", "problems.csv", "a filename")
	timerFlag := flag.Int64("timer", 30, "timer")
	flag.Parse()

	quiz, err := OpenCSV(*filenameFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	startQuiz(quiz, *timerFlag)
}
