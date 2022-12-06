package day6

import (
	"bufio"
	"fmt"
	"os"
	snippets "practice/snippets"
	"strings"
)

func FirstPart() {

	input, _ := os.Open("./adventofcode/day6/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	var result int
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		//fmt.Println(line[31])
		for i := 0; i < len(line)-4; i += 1 {
			chunk := line[i+0 : i+4]
			chunk_set := snippets.NewSet()
			for _, v := range chunk {
				chunk_set.Add(string(v))
			}
			if chunk_set.Size() == len(chunk) {
				result += (i + 4)
				break
			}

		}
		fmt.Println(result)
	}
}
