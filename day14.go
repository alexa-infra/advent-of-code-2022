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
	data := map[coord]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		path := []coord{}
		for _, part := range parts {
			coords := strings.Split(part, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			path = append(path, coord{x, y})
		}
		for i := 1; i < len(path); i++ {
			curr := path[i-1]
			next := path[i]
			for curr != next {
				data[curr] = '#'
				if curr.x == next.x {
					if curr.y < next.y {
						curr.y++
					} else {
						curr.y--
					}
				} else {
					if curr.x < next.x {
						curr.x++
					} else {
						curr.x--
					}
				}
			}
			data[curr] = '#'
		}
	}
	var minx, maxx, miny, maxy int
	found := false
	for p, _ := range data {
		if !found {
			minx, maxx, miny, maxy = p.x, p.x, p.y, p.y
			found = true
		} else {
			minx = min(minx, p.x)
			miny = min(miny, p.y)
			maxx = max(maxx, p.x)
			maxy = max(maxy, p.y)
		}
	}
	miny -= 2
	minx -= 2
	maxx += 2
	maxy += 2
	//printData := func() {
	//	for y := miny; y <= maxy; y++ {
	//		for x := minx; x <= maxx; x++ {
	//			v, ok := data[coord{x, y}]
	//			if !ok {
	//				fmt.Print(".")
	//			} else {
	//				fmt.Print(string([]byte{ v }))
	//			}
	//		}
	//		fmt.Println()
	//	}
	//	fmt.Println()
	//}
	//fmt.Println(minx, maxx, miny, maxy)
	start := coord{500, 0}
	queue := []coord{ start }
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.y >= maxy {
			break
		}
		below := coord{curr.x, curr.y + 1}
		_, ok := data[below]
		if !ok {
			queue = append(queue, below)
			continue
		}
		left := coord{curr.x - 1, curr.y + 1}
		_, ok = data[left]
		if !ok {
			queue = append(queue, left)
			continue
		}
		right := coord{curr.x + 1, curr.y + 1}
		_, ok = data[right]
		if !ok {
			queue = append(queue, right)
			continue
		}
		data[curr] = 'o'
		queue = append(queue, start)
		//printData()
	}
	p1 := 0
	for _, v := range data {
		if v == 'o' {
			p1++
		}
	}
	fmt.Println("Part 1:", p1)
	queue = append(queue, start)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		_, ok := data[curr]
		if ok {
			break
		}
		if curr.y >= maxy - 1 {
			data[curr] = 'o'
			queue = append(queue, start)
			continue
		}
		below := coord{curr.x, curr.y + 1}
		_, ok = data[below]
		if !ok {
			queue = append(queue, below)
			continue
		}
		left := coord{curr.x - 1, curr.y + 1}
		_, ok = data[left]
		if !ok {
			queue = append(queue, left)
			continue
		}
		right := coord{curr.x + 1, curr.y + 1}
		_, ok = data[right]
		if !ok {
			queue = append(queue, right)
			continue
		}
		data[curr] = 'o'
		queue = append(queue, start)
		//printData()
	}
	p2 := 0
	for _, v := range data {
		if v == 'o' {
			p2++
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
