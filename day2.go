package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("resources/day2input.txt")

	if err != nil {
		fmt.Println(err)
	}

	currentScore := 0

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		lineStr := fileScanner.Text()
		opponentsChoice := string(lineStr[0])
		outcome := string(lineStr[2])
		yourChoice := whatShouldYouChoose(opponentsChoice, outcome)
		currentScore = currentScore + characterValue(yourChoice)
		currentScore = currentScore + didYouWin(opponentsChoice, yourChoice)
	}

	fmt.Println(currentScore)
}

func whatShouldYouChoose(opponent, outcome string) string {
	if outcome == "X" {
		switch opponent {
		case "A":
			return "Z"
		case "B":
			return "X"
		case "C":
			return "Y"
		}
	} else if outcome == "Y" {
		switch opponent {
		case "A":
			return "X"
		case "B":
			return "Y"
		case "C":
			return "Z"
		}
	} else if outcome == "Z" {
		switch opponent {
		case "A":
			return "Y"
		case "B":
			return "Z"
		case "C":
			return "X"
		}
	}
	return ""
}

func characterValue(a string) int {
	switch a {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}

func didYouWin(opponent, yourself string) int {
	if yourself == "X" {
		switch opponent {
		case "A":
			return 3
		case "B":
			return 0
		case "C":
			return 6
		}
	} else if yourself == "Y" {
		switch opponent {
		case "A":
			return 6
		case "B":
			return 3
		case "C":
			return 0
		}
	} else if yourself == "Z" {
		switch opponent {
		case "A":
			return 0
		case "B":
			return 6
		case "C":
			return 3
		}
	}
	return 0
}
