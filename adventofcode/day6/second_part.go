package day6

import (
	"bufio"
	"fmt"
	"os"
	snippets "practice/snippets"
	"strings"
)

func SecondPart() {

	input, _ := os.Open("./adventofcode/day6/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	var result int
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		for i := 0; i < len(line)-14; i += 1 {
			chunk := line[i+0 : i+14]
			chunk_set := snippets.NewSet()
			for _, v := range chunk {
				chunk_set.Add(string(v))
			}
			if chunk_set.Size() == len(chunk) {
				result += (i + 14)
				break
			}

		}
		fmt.Println(result)
	}
}
