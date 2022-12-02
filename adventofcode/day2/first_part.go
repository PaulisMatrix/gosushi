package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FirstPart() {
	file, err := os.Open("./adventofcode/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//A = X = Rock 1
	//B = Y = Paper 2
	//C = Z = Scissor 3

	var scores int

	your_scores := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}

	for scanner.Scan() {
		scores += your_scores[scanner.Text()]

	}
	fmt.Println(scores)

}
