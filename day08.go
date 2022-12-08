package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	type coord struct {
		x, y int
	}
	m := map[coord]int{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, ch := range line {
			c := coord{i, j}
			h, _ := strconv.Atoi(string(ch))
			m[c] = h
		}
		i++
	}
	dx := []int{ 0, 1, 0, -1 }
	dy := []int{ 1, 0, -1, 0 }
	p1 := 0
	for k, h := range m {
		visible := false
		for i := 0; i < 4; i++ {
			kk := coord{ k.x, k.y }
			for {
				kk.x += dx[i]
				kk.y += dy[i]
				hh, ok := m[kk]
				if !ok {
					visible = true
					break
				}
				if hh >= h {
					break
				}
			}
			if visible {
				break
			}
		}
		if visible {
			p1 += 1
		}
	}
	fmt.Println("Part 1:", p1)
	p2 := 0
	for k, h := range m {
		score := 1
		for i := 0; i < 4; i++ {
			kk := coord{ k.x, k.y }
			cc := 0
			for {
				kk.x += dx[i]
				kk.y += dy[i]
				hh, ok := m[kk]
				if !ok {
					break
				}
				if hh >= h {
					cc++
					break
				}
				cc++
			}
			score *= cc
		}
		if score > p2 {
			p2 = score
		}
	}
	fmt.Println("Part 2:", p2)
}
