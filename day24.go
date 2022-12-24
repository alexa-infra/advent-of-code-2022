package main

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
)

func main() {
	type coord struct {
		x, y int
	}
	re := regexp.MustCompile(`^#+(\.)#+$`)
	scanner := bufio.NewScanner(os.Stdin)
	var start, end coord
	var hasStart, hasEnd bool
	j := 0
	data := map[coord][]byte{}
	for scanner.Scan() {
		line := scanner.Bytes()
		if re.Match(line) {
			for i, ch := range line {
				p := coord{ i, j }
				if ch == '.' {
					data[p] = []byte{}
					if !hasStart {
						start = p
						hasStart = true
					} else if !hasEnd {
						end = p
						hasEnd = true
					}
				}
			}
		} else {
			for i, ch := range line {
				p := coord{ i, j }
				if ch == '.' {
					data[p] = []byte{}
				} else if ch != '#' {
					data[p] = []byte{ ch }
				}
			}
		}
		j++
	}
	moveBlizzards := func() {
		newData := map[coord][]byte{}
		for p, arr := range data {
			_, ok := newData[p]
			if !ok {
				newData[p] = []byte{}
			}
			for _, ch := range arr {
				pp := p
				neg := coord{0, 0}
				if ch == '>' {
					pp = coord{p.x + 1, p.y}
					neg = coord{-1, 0}
				}
				if ch == '<' {
					pp = coord{p.x - 1, p.y}
					neg = coord{1, 0}
				}
				if ch == '^' {
					pp = coord{p.x, p.y - 1}
					neg = coord{0, 1}
				}
				if ch == 'v' {
					pp = coord{p.x, p.y + 1}
					neg = coord{0, -1}
				}
				_, ok1 := data[pp]
				if !ok1 {
					for {
						test := coord{pp.x + neg.x, pp.y + neg.y}
						if _, ok := data[test]; ok {
							pp = test
						} else {
							break
						}
					}
				}
				m, ok2 := newData[pp]
				if !ok2 {
					newData[pp] = []byte{ch}
				} else {
					newData[pp] = append(m, ch)
				}
			}
		}
		data = newData
	}
	diff := []coord{ coord{ 1, 0 }, coord{ -1, 0 }, coord{ 0, 1 }, coord{ 0, -1 }, coord{ 0, 0 } }
	getFreeAdj := func(p coord) []coord{
		arr := []coord{}
		for _, d := range diff {
			pp := coord{p.x + d.x, p.y + d.y }
			v, ok := data[pp]
			if ok && len(v) == 0 {
				arr = append(arr, pp)
			}
		}
		return arr
	}
	moveElf := func(start, end coord) int {
		state := []coord{ start }
		n := 0
		for len(state) > 0 {
			newState := []coord{}
			found := false
			moveBlizzards()
			for _, current := range state {
				if current == end {
					found = true
					break
				}
				for _, a := range getFreeAdj(current) {
					if !contains(newState, a) {
						newState = append(newState, a)
					}
				}
			}
			if found {
				break
			}
			state = newState
			n++
		}
		return n
	}
	p1 := moveElf(start, end)
	fmt.Println("Part 1:", p1)
	p2 := p1
	p2 += moveElf(end, start) + 1
	p2 += moveElf(start, end) + 1
	fmt.Println("Part 2:", p2)
}

func contains[T comparable](arr []T, x T) bool {
	for _, y := range arr {
		if x == y {
			return true
		}
	}
	return false
}
