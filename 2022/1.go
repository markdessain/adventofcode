package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./2022/data/1.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	index := 0
	currentCount := 0
	maxCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentCount > maxCount {
				maxCount = currentCount
			}
			currentCount = 0
			index += 1
		} else {
			count, err := strconv.Atoi(line)

			if err != nil {
				fmt.Println(err)
			} else {
				currentCount += count
			}
		}
	}

	fmt.Println(maxCount)
	fmt.Println(index)

}
