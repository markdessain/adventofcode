package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2022/data/5.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	stacks := make(map[int][]string)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.HasPrefix(line, " 1") {
			break
		}

		stackCount := (len(line) + 1) / 4

		for i := 0; i < stackCount; i++ {

			value := line[(i*4)+1 : (i*4)+2]

			if value != " " {
				stacks[i+1] = append([]string{value}, stacks[i+1]...)
			}
		}

	}

	scanner.Scan()

	fmt.Println(stacks)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.Replace(line, "move ", "", 1)
		parts := strings.Split(line, " from ")

		count, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
		}

		parts = strings.Split(parts[1], " to ")

		from, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
		}
		to, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < count; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	keys := []int{}

	for k := range stacks {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		l := stacks[k]
		fmt.Print(l[len(l)-1])
	}

	fmt.Println("")
}
