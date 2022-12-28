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

	for scanner.Scan() {
		line := scanner.Text()

		score += winningMap[line]
	}

	fmt.Println(score)

}
