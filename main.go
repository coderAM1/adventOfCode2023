package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("resources/day1part1.txt")

	if err != nil {
		fmt.Println(err)
	}

	var numbers []int
	var currentNum int = 0

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		lineStr := fileScanner.Text()
		if lineStr != "" {
			num, _ := strconv.Atoi(lineStr)
			currentNum += num
		} else {
			numbers = append(numbers, currentNum)
			currentNum = 0
		}
	}
	numbers = append(numbers, currentNum)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	highestNumber := numbers[0] + numbers[1] + numbers[2]

	fmt.Println(highestNumber)
}
