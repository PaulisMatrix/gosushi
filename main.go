package main

import (
	snippets "practice/snippets"
	//websockets "practice/websockets"
	//cards "practice/cards"
)

func main() {

	//day2.FirstPart()
	//day2.SecondPart()

	myset1 := snippets.NewSet()
	myset1.AddMulti([]string{"a", "b", "c", "d"})

	myset2 := snippets.NewSet()
	myset2.AddMulti([]string{"e", "f", "g", "d"})

	inter := myset1.Intersection(myset2)

	inter.Display()

	union := myset1.Union(myset2)

	union.Display()

	diff := myset1.Difference(myset2)

	diff.Display()

}
