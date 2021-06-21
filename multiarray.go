package main

import "fmt"

func main() {
	fmt.Println("다차원 배열 연습")

	var singlearray = [100]int{50: 12, 99: 5}
	fmt.Println(singlearray)

	var doublearray = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(doublearray)

	var multiarray = [2][3][2]int{
		{{1, 2},
			{3, 4},
			{5, 6}},
		{{7, 8},
			{9, 10},
			{11, 12}},
	}
	fmt.Println(multiarray)
}
