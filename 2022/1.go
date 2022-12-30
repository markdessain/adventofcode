package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	allCount := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			allCount = append(allCount, currentCount)
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

	sort.Ints(allCount)

	fmt.Println(allCount[len(allCount)-1])

	topThree := 0

	for _, c := range allCount[len(allCount)-3:] {
		topThree += c
	}
	fmt.Println(topThree)

}
