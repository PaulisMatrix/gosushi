package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FirstPart() {
	file, err := os.Open("./adventofcode/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxCals := 0
	currentElfCals := 0

	for scanner.Scan() {
		cal, err := strconv.Atoi(scanner.Text())
		currentElfCals += cal

		//incase found a blank line("\n")
		if err != nil {
			if currentElfCals > maxCals {
				maxCals = currentElfCals
			}
			currentElfCals = 0
		}
	}
	fmt.Println(maxCals)
}
