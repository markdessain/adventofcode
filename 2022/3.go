package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./2022/data/3.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	score := 0

	priorityMap := make(map[string]int)
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	for _, c := range lowercase {
		priorityMap[string(c)] = int(c) - 96
	}

	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range uppercase {
		priorityMap[string(c)] = int(c) - 38
	}

	for scanner.Scan() {

		items := make(map[string]bool)

		line := scanner.Text()

		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]

		for _, k := range compartment1 {
			if strings.Contains(compartment2, string(k)) {
				items[string(k)] = true
			}
		}

		for item, _ := range items {
			score += priorityMap[item]
		}
	}

	fmt.Println(score)
}
