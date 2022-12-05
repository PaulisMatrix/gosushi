package day5

import (
	"bufio"
	"fmt"
	"os"
	"practice/snippets"
)

func FirstPart() {
	//read the file
	input, _ := os.Open("./adventofcode/day5/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)

	//slice of stacks
	stacks := make([]*snippets.Stack, 0)

	for i := 0; i < 9; i++ {
		stacks = append(stacks, snippets.NewStack())
	}

	//Parsing the input
	sc.Scan()
	for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, r := range sc.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].AddToBottom(r)
			}
		}
		sc.Scan()
	}
	//empty line
	sc.Scan()
	for sc.Scan() {
		var toMove, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

		//Move elements one by one
		for move := 0; move < toMove; move++ {
			stacks[to-1].Push(stacks[from-1].Pop(1))
		}
	}

	for _, s := range stacks {
		fmt.Print(string(s.Pop(1)[0]))
	}

}
