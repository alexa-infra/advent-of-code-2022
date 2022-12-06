package main

import (
	"bufio"
	"os"
	"fmt"
)

func isUnique(arr []rune) bool {
	m := map[rune]int{}
	for _, ch := range arr {
		_, ok := m[ch]
		if ok {
			return false
		}
		m[ch] = 1
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := []rune{}
	arr2 := []rune{}
	p1 := -1
	p2 := -1
	scanner.Scan()
	for i, ch := range scanner.Text() {
		if i < 4 {
			arr = append(arr, ch)
		} else {
			arr = arr[1:]
			arr = append(arr, ch)
		}
		if p1 == -1 && len(arr) == 4 && isUnique(arr) {
			p1 = i + 1
		}
		if i < 14 {
			arr2 = append(arr2, ch)
		} else {
			arr2 = arr2[1:]
			arr2 = append(arr2, ch)
		}
		if p2 == -1 && len(arr2) == 14 && isUnique(arr2) {
			p2 = i + 1
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
