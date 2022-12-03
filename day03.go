package main

import (
	"bufio"
	"fmt"
	"os"
)

func intersect(a, b string) string {
	m := map[byte]bool{}
	for i := 0; i < len(a); i++ {
		m[a[i]] = true
	}
	r := []byte{}
	for i := 0; i < len(b); i++ {
		_, ok := m[b[i]]
		if ok {
			r = append(r, b[i])
		}
	}
	return string(r)
}

func byteToPriority(a byte) int {
	if a < 'a' {
		return int(a - 65 + 27)
	}
	return int(a - 97 + 1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	priority := 0
	group := make([]string, 3)
	i := 0
	groupPriority := 0
	for scanner.Scan() {
		line := scanner.Text()
		group[i] = line
		n := len(line)
		item := intersect(line[:n/2], line[n/2:])
		priority += byteToPriority(item[0])
		if i == 2 {
			set1 := intersect(group[0], group[1])
			set2 := intersect(group[1], group[2])
			set3 := intersect(set1, set2)
			groupPriority += byteToPriority(set3[0])
			i = 0
		} else {
			i++
		}
	}
	fmt.Println("Part 1:", priority)
	fmt.Println("Part 2:", groupPriority)
}
