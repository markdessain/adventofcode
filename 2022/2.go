package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./2022/data/2.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	score := 0
	matchingscore := 0

	winningMap := make(map[string]int)
	winningMap["A X"] = 3 + 1
	winningMap["A Y"] = 6 + 2
	winningMap["A Z"] = 0 + 3
	winningMap["B X"] = 0 + 1
	winningMap["B Y"] = 3 + 2
	winningMap["B Z"] = 6 + 3
	winningMap["C X"] = 6 + 1
	winningMap["C Y"] = 0 + 2
	winningMap["C Z"] = 3 + 3

	matchingMap := make(map[string]int)
	matchingMap["A X"] = 0 + 3
	matchingMap["A Y"] = 3 + 1
	matchingMap["A Z"] = 6 + 2
	matchingMap["B X"] = 0 + 1
	matchingMap["B Y"] = 3 + 2
	matchingMap["B Z"] = 6 + 3
	matchingMap["C X"] = 0 + 2
	matchingMap["C Y"] = 3 + 3
	matchingMap["C Z"] = 6 + 1

	for scanner.Scan() {
		line := scanner.Text()

		score += winningMap[line]
		matchingscore += matchingMap[line]
	}

	fmt.Println(score)
	fmt.Println(matchingscore)

}
