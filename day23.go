package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	type coord struct {
		x, y int
	}
	j := 0
	elves := map[coord]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for i, ch := range scanner.Text() {
			if ch == '#' {
				p := coord{i, j}
				elves[p] = 0
			}
		}
		j++
	}
	nswe := [][]coord {
		[]coord{ coord{ -1, -1 }, coord{ 0, -1 }, coord{ 1, -1 } }, // N
		[]coord{ coord{ -1,  1 }, coord{ 0,  1 }, coord{ 1,  1 } }, // S
		[]coord{ coord{ -1, -1 }, coord{ -1, 0 }, coord{ -1, 1 } }, // W
		[]coord{ coord{  1, -1 }, coord{  1, 0 }, coord{  1, 1 } }, // E
	}
	play := func(round int) (moved int) {
		moves := map[coord][]coord{}
		for p := range elves {
			shouldMove := false
			for i := 0; i < 4; i++ {
				for _, diff := range nswe[i] {
					pp := coord{ p.x + diff.x, p.y + diff.y }
					_, ok := elves[pp]
					if ok {
						shouldMove = true
						break
					}
				}
				if shouldMove {
					break
				}
			}
			if !shouldMove {
				continue
			}
			for i := 0; i < 4; i++ {
				id := (i + round % 4) % 4
				canMove := true
				for _, diff := range nswe[id] {
					pp := coord{ p.x + diff.x, p.y + diff.y }
					_, ok := elves[pp]
					if ok {
						canMove = false
						break
					}
				}
				if canMove {
					dir := nswe[id][1]
					next := coord{ p.x + dir.x, p.y + dir.y }
					m, ok := moves[next]
					if ok {
						moves[next] = append(m, p)
					} else {
						moves[next] = []coord{ p }
					}
					break
				}
			}
		}
		moved = 0
		for next, origList := range moves {
			if len(origList) == 1 {
				p := origList[0]
				delete(elves, p)
				elves[next] = 0
				moved++
			}
		}
		return moved
	}
	for round := 0; round < 10; round++ {
		play(round)
	}
	minx, maxx, miny, maxy := 0, 0, 0, 0
	first := true
	for p := range elves {
		if first {
			minx, maxx, miny, maxy = p.x, p.x, p.y, p.y
			first = false
			continue
		}
		minx = min(minx, p.x)
		miny = min(miny, p.y)
		maxx = max(maxx, p.x)
		maxy = max(maxy, p.y)
	}
	p1 := 0
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			p := coord{x, y}
			_, ok := elves[p]
			if !ok {
				p1 += 1
			}
		}
	}
	//for y := miny; y <= maxy; y++ {
	//	for x := minx; x <= maxx; x++ {
	//		p := coord{x, y}
	//		_, ok := elves[p]
	//		if !ok {
	//			fmt.Print(".")
	//		} else {
	//			fmt.Print("#")
	//		}
	//	}
	//	fmt.Println()
	//}
	fmt.Println("Part 1:", p1)
	p2 := 0
	for round := 10; round < 100000; round++ {
		moved := play(round)
		if moved == 0 {
			p2 = round + 1
			break
		}
	}
	fmt.Println("Part 2:", p2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
