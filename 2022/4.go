package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2022/data/4.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	score := 0

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, ",")

		parts1 := strings.Split(parts[0], "-")
		parts2 := strings.Split(parts[1], "-")

		parts1start, err := strconv.Atoi(parts1[0])
		if err != nil {
			fmt.Println(err)
		}

		parts1end, err := strconv.Atoi(parts1[1])
		if err != nil {
			fmt.Println(err)
		}

		parts2start, err := strconv.Atoi(parts2[0])
		if err != nil {
			fmt.Println(err)
		}

		parts2end, err := strconv.Atoi(parts2[1])
		if err != nil {
			fmt.Println(err)
		}


		if parts1start <= parts2start && parts1end >= parts2end {
			score += 1
		} else if parts2start <= parts1start && parts2end >= parts1end {
			score += 1
		}

	}

	fmt.Println(score)
}
