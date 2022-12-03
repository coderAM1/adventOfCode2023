package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("resources/day3input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	var lineArray []string

	score := 0
	score2 := 0

	for fileScanner.Scan() {
		lineStr := fileScanner.Text()
		lineArray = append(lineArray, lineStr)
		strLength := len(lineStr)
		halfLength := (strLength / 2)
		firstHalf := lineStr[0:halfLength]
		secondHalf := lineStr[halfLength:strLength]
		var alreadyChecked string
		for _, element := range firstHalf {
			if strings.Contains(secondHalf, string(element)) && !strings.Contains(alreadyChecked, string(element)) {
				if element <= 91 {
					score += (int(element) - 64) + 26
				} else {
					score += int(element) - 96
				}
			}
			alreadyChecked += string(element)
		}
		fmt.Println(lineArray)
		if len(lineArray) == 3 {
			var alreadyChecked2 string
			for _, element := range lineArray[0] {
				if strings.Contains(lineArray[1], string(element)) && strings.Contains(lineArray[2], string(element)) && !strings.Contains(alreadyChecked2, string(element)) {
					if element <= 91 {
						score2 += (int(element) - 64) + 26
					} else {
						score2 += int(element) - 96
					}
				}
				alreadyChecked2 += string(element)
			}
			lineArray = nil
		}
	}

	fmt.Println(score)
	fmt.Println(score2)
}
