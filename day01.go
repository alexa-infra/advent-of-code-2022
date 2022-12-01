package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	currentCalories := 0
	elves := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			calories, _ := strconv.Atoi(line)
			currentCalories += calories
		} else {
			elves = append(elves, currentCalories)
			currentCalories = 0
		}
	}
	elves = append(elves, currentCalories)
	sort.Ints(elves)
	n := len(elves)
	fmt.Println("Part 1:", elves[n-1])
	r := elves[n-1] + elves[n-2] + elves[n-3]
	fmt.Println("Part 2:", r)
}
