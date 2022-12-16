package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SecondPart() {
	//Read input file
	input, _ := os.Open("./adventofcode/day10/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	registerX, cycleNumber := 1, 0

	for sc.Scan() {
		operations := strings.Fields(sc.Text())
		IncrementAndControl(&cycleNumber, &registerX)
		if operations[0] == "addx" {
			value, _ := strconv.Atoi(operations[1])
			IncrementAndControl(&cycleNumber, &registerX)
			registerX += value
		}
	}
}

func IncrementAndControl(cycleNumber, registerX *int) {
	if (*cycleNumber)%40 == 0 && *cycleNumber <= 220 {
		fmt.Println()
	}
	if *registerX-1 == *cycleNumber%40 || *registerX == *cycleNumber%40 || *registerX+1 == *cycleNumber%40 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	*cycleNumber++
}
