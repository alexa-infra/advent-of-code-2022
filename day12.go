package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := make([]byte, len(scanner.Bytes()))
		copy(row, scanner.Bytes())
		data = append(data, row)
	}
	m, n := len(data), len(data[0])
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	type coord struct {
		x, y int
	}
	var start, end coord
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if data[i][j] == 'S' {
				data[i][j] = 'a'
				start = coord{i, j}
			}
			if data[i][j] == 'E' {
				data[i][j] = 'z'
				end = coord{i, j}
			}
		}
	}
	run := func(start coord, cond func(coord, byte) bool, step func(byte, byte) bool) int {
		type state struct {
			p coord
			n int
		}
		queue := []state{state{start, 0}}
		visited := map[coord]int{start: 0}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			currentH := data[current.p.x][current.p.y]
			if cond(current.p, currentH) {
				return current.n
			}
			for i := 0; i < 4; i++ {
				p := coord{current.p.x + dx[i], current.p.y + dy[i]}
				if p.x >= 0 && p.x < m && p.y >= 0 && p.y < n {
					h := data[p.x][p.y]
					if step(currentH, h) {
						if v, ok := visited[p]; !ok || v > current.n+1 {
							visited[p] = current.n + 1
							queue = append(queue, state{p, current.n + 1})
						}
					}
				}
			}
		}
		return -1
	}
	p1 := run(start, func(p coord, h byte) bool { return p == end }, func(a, b byte) bool { return b-1 <= a })
	fmt.Println("Part 1:", p1)
	p2 := run(end, func(p coord, h byte) bool { return h == 'a' }, func(a, b byte) bool { return a-1 <= b })
	fmt.Println("Part 2:", p2)
}
