package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func SecondPart(numelfs int) {
	file, err := os.Open("./adventofcode/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calsList := []int{}

	scanner := bufio.NewScanner(file)

	sumval := 0
	for scanner.Scan() {
		cal, err := strconv.Atoi(scanner.Text())
		sumval += cal

		//incase found a blank line("\n")
		if err != nil {
			calsList = append(calsList, sumval)
			sumval = 0
		}
	}
	//sort in reverse
	sort.Sort(sort.Reverse(sort.IntSlice(calsList)))
	result := 0
	for _, v := range calsList[:numelfs] {
		result += v
	}
	fmt.Println(result)
}
