package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SecondPart() {
	file, err := os.Open("./adventofcode/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//A = X = Rock 1
	//B = Y = Paper 2
	//C = Z = Scissor 3

	//X = lose
	//Y = draw
	//Z = win

	var scores int

	your_scores := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}

	for scanner.Scan() {
		scores += your_scores[scanner.Text()]

	}
	fmt.Println(scores)

}
