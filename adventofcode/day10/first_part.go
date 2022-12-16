package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FirstPart() {
	input, _ := os.Open("./adventofcode/day10/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	registerX, cycleNumber := 1, 0
	var result int

	for sc.Scan() {
		operations := strings.Fields(sc.Text())
		incrementAndControl(&cycleNumber, &registerX, &result)
		if operations[0] == "addx" {
			value, _ := strconv.Atoi(operations[1])
			incrementAndControl(&cycleNumber, &registerX, &result)
			registerX += value
		}
	}
	fmt.Println(result)
}

func incrementAndControl(cycleNumber, registerX, finalValue *int) {
	*cycleNumber++
	if (*cycleNumber-20)%40 == 0 && *cycleNumber <= 220 {
		*finalValue += *registerX * *cycleNumber
	}
}
