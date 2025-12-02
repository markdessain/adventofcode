package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Run(``))
}

func Run(ranges string) int {

	counter := 0
	for _, r := range strings.Split(ranges, ",") {
		parts := strings.Split(r, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
		}
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)

			if len(s)%2 == 0 {
				if s[:len(s)/2] == s[len(s)/2:] {
					counter += i
				}
			}
		}
	}
	return counter
}
