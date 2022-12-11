package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func splitEmptyLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(splitEmptyLine)
	r1 := regexp.MustCompile(`Monkey (\d+):
  Starting items: (.+)
  Operation: new = old (.) (\d+|old)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)`)
	r2 := regexp.MustCompile(`\d+`)
	type monkey struct {
		id         int
		startItems []int
		items      []int
		operation  func(int)int
		test       int
		positive   int
		negative   int
		counter    int
	}
	monkeys := []monkey{}
	for scanner.Scan() {
		m1 := r1.FindStringSubmatch(scanner.Text())
		monkeyId, _ := strconv.Atoi(m1[1])
		m2 := r2.FindAllString(m1[2], -1)
		items := []int{}
		for _, m := range m2 {
			item, _ := strconv.Atoi(m)
			items = append(items, item)
		}
		ref, _ := strconv.Atoi(m1[4])
		var op func(int)int
		if m1[3] == "+" {
			op = func(old int) int { return old + ref; }
		} else if m1[3] == "*" {
			if m1[4] == "old" {
				op = func(old int) int { return old * old; }
			} else {
				op = func(old int) int { return old * ref; }
			}
		}
		div, _ := strconv.Atoi(m1[5])
		positive, _ := strconv.Atoi(m1[6])
		negative, _ := strconv.Atoi(m1[7])
		monkeys = append(monkeys, monkey{monkeyId, items, items, op, div, positive, negative, 0})
	}
	run := func(rounds int, deduce func(int) int) int {
		for i := 0; i < len(monkeys); i++ {
			current := &monkeys[i]
			current.counter = 0
			current.items = make([]int, len(current.startItems))
			copy(current.items, current.startItems)
		}
		for i := 0; i < rounds; i++ {
			for j := 0; j < len(monkeys); j++ {
				current := &monkeys[j]
				for _, item := range current.items {
					item = current.operation(item)
					item = deduce(item)
					var target *monkey
					if item % current.test == 0 {
						target = &monkeys[current.positive]
					} else {
						target = &monkeys[current.negative]
					}
					target.items = append(target.items, item)
					current.counter++
				}
				current.items = current.items[:0]
			}
		}
		counters := make([]int, len(monkeys))
		for i := 0; i < len(monkeys); i++ {
			counters[i] = monkeys[i].counter
		}
		sort.Sort(sort.Reverse(sort.IntSlice(counters)))
		return counters[0] * counters[1]
	}
	p1 := run(20, func(v int) int { return v / 3 })
	fmt.Println("Part 1:", p1)
	lcm := 1
	for i := 0; i < len(monkeys); i++ {
		lcm *= monkeys[i].test
	}
	p2 := run(10000, func(v int) int { return v % lcm })
	fmt.Println("Part 2:", p2)
}
