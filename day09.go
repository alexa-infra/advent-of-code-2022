package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	type coord struct {
		x, y int
	}
	shouldMove := func(head, tail coord) bool {
		return abs(head.x - tail.x) > 1 || abs(head.y - tail.y) > 1
	}
	diff := map[string]coord{
		"U": coord{0, 1},
		"D": coord{0, -1},
		"R": coord{1, 0},
		"L": coord{-1, 0},
	}
	//printStep := func(dir string, times int, chain []coord) {
	//	fmt.Println(dir, times)
	//	maxX, maxY, minX, minY := 0, 0, 0, 0
	//	for i := 0; i < len(chain); i++ {
	//		c := chain[i]
	//		minX = min(minX, c.x)
	//		minY = min(minY, c.y)
	//		maxX = max(maxX, c.x)
	//		maxY = max(maxY, c.y)
	//	}
	//	minX -= 2
	//	maxX += 2
	//	minY -= 2
	//	maxY += 2
	//	for j := maxY; j >= minY; j-- {
	//		for i := minX; i < maxX; i++ {
	//			if i == 0 && j == 0 {
	//				fmt.Print("s")
	//				continue
	//			}
	//			found := false
	//			for k := 0; k < len(chain); k++ {
	//				c := chain[k]
	//				if c.x == i && c.y == j {
	//					fmt.Print(k)
	//					found = true
	//					break
	//				}
	//			}
	//			if !found {
	//				fmt.Print(".")
	//			}
	//		}
	//		fmt.Println()
	//	}
	//	fmt.Println()
	//}
	move := func(dir string, times int, chain []coord, visited map[coord]bool) {
		d := diff[dir]
		for i := 0; i < times; i++ {
			chain[0] = coord{ chain[0].x + d.x, chain[0].y + d.y }
			for j := 1; j < len(chain); j++ {
				if shouldMove(chain[j - 1], chain[j]) {
					dx, dy := 0, 0
					if abs(chain[j - 1].x - chain[j].x) >= 1 {
						if chain[j].x < chain[j - 1].x {
							dx += 1
						} else {
							dx -= 1
						}
					}
					if abs(chain[j - 1].y - chain[j].y) >= 1 {
						if chain[j].y < chain[j - 1].y {
							dy += 1
						} else {
							dy -= 1
						}
					}
					chain[j] = coord{ chain[j].x + dx, chain[j].y + dy }
					if j == len(chain) - 1 {
						visited[chain[j]] = true
					}
				}
			}
			//printStep(dir, times, chain)
		}
	}
	type op struct {
		dir string
		times int
	}
	ops := []op{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		dir := parts[0]
		times, _ := strconv.Atoi(parts[1])
		ops = append(ops, op{ dir, times })
	}
	knots := []coord{ coord{0, 0}, coord{0, 0} }
	visited := map[coord]bool{}
	visited[coord{0, 0}] = true
	for _, x := range ops {
		move(x.dir, x.times, knots, visited)
	}
	p1 := len(visited)
	fmt.Println("Part 1:", p1)

	knots = []coord{}
	for i := 0; i < 10; i++ {
		knots = append(knots, coord{0, 0})
	}
	visited = map[coord]bool{}
	visited[coord{0, 0}] = true
	for _, x := range ops {
		move(x.dir, x.times, knots, visited)
	}
	p2 := len(visited)
	fmt.Println("Part 2:", p2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
