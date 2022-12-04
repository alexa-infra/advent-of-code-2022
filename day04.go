package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"fmt"
)

func isBetween(x, a, b int) bool {
	return a <= x && x <= b
}

func main() {
	re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	p1 := 0
	p2 := 0
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		r1, _ := strconv.Atoi(match[1])
		r2, _ := strconv.Atoi(match[2])
		r3, _ := strconv.Atoi(match[3])
		r4, _ := strconv.Atoi(match[4])
		if isBetween(r1, r3, r4) && isBetween(r2, r3, r4) {
			p1++
		} else if isBetween(r3, r1, r2) && isBetween(r4, r1, r2) {
			p1++
		}
		if isBetween(r1, r3, r4) || isBetween(r2, r3, r4) {
			p2++
		} else if isBetween(r3, r1, r2) || isBetween(r4, r1, r2) {
			p2++
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
