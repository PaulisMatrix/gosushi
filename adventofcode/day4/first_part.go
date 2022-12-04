package day4

import (
	"bufio"
	"fmt"
	"os"
)

func FirstPart() {

	input, _ := os.Open("./adventofcode/day4/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	var result int

	for sc.Scan() {
		var x0, x1, y0, y1 int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &x0, &x1, &y0, &y1)
		if x0 <= y0 && x1 >= y1 || x0 >= y0 && x1 <= y1 {
			result += 1
		}

	}

	fmt.Println(result)
}
