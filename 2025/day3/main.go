package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Run(``))
}

func Run(lines string) int {
	counter := 0

	for _, line := range strings.Split(lines, "\n") {

		first := 0
		firstIndex := 0
		for i, c := range line[:len(line)-1] {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println(err)
			}

			if n > first {
				first = n
				firstIndex = i
			}

		}

		second := 0
		for _, c := range line[firstIndex+1:] {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println(err)
			}

			if n > second {
				second = n
			}
		}

		d, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
		if err != nil {
			fmt.Println(err)
		}
		counter += d
	}
	return counter
}
