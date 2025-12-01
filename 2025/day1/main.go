package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(lines string, dial int) (int, int) {

	timesZero := 0

	for _, line := range strings.Split(lines, "\n") {

		direction := line[0]
		number, err := strconv.Atoi(line[1:len(line)])

		if err != nil {
			fmt.Println(err)
		}

		if direction == 'R' {
			if dial+number > 99 {
				dial = (dial + number) - 100
			} else {
				dial = dial + number
			}
		} else if direction == 'L' {
			if dial-number < 0 {
				dial = 100 - (number - dial)
			} else {
				dial = dial - number
			}
		}

		if dial == 0 {
			timesZero += 1
		}
	}

	return dial, timesZero
}
