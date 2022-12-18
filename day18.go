package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	x, y, z int
}

func main() {
	cubes := map[cube]bool{}
	diffs := []cube{
		cube{1, 0, 0},
		cube{-1, 0, 0},
		cube{0, 1, 0},
		cube{0, -1, 0},
		cube{0, 0, 1},
		cube{0, 0, -1},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		c := cube{x, y, z}
		cubes[c] = true
	}
	countNotConnectedSides := 0
	notConnected := map[cube]bool{}
	firstCube := cube{0, 0, 0}
	firstFound := false
	for c := range cubes {
		if firstFound {
			firstCube = c
			firstFound = true
		}
		for _, d := range diffs {
			cc := cube{c.x + d.x, c.y + d.y, c.z + d.z}
			_, ok := cubes[cc]
			if !ok {
				countNotConnectedSides++
				notConnected[cc] = true
			}
		}
	}
	fmt.Println("Part 1:", countNotConnectedSides)
	maxX, minX, maxY, minY, maxZ, minZ := firstCube.x, firstCube.x, firstCube.y, firstCube.y, firstCube.z, firstCube.z
	for c := range cubes {
		maxX = max(maxX, c.x)
		minX = min(minX, c.x)
		maxY = max(maxY, c.y)
		minY = min(minY, c.y)
		maxZ = max(maxZ, c.z)
		minZ = min(minZ, c.z)
	}
	minX -= 1
	minY -= 1
	minZ -= 1
	maxX += 1
	maxY += 1
	maxZ += 1
	start := cube{minX, minY, minZ}
	queue := []cube{start}
	water := map[cube]bool{start: true}
	exteriorSides := 0
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for _, d := range diffs {
			cc := cube{c.x + d.x, c.y + d.y, c.z + d.z}
			if cc.x < minX || cc.x > maxX || cc.y < minY || cc.y > maxY || cc.z < minZ || cc.z > maxZ {
				continue
			}
			_, ok1 := water[cc]
			if ok1 {
				continue
			}
			_, ok2 := cubes[cc]
			if ok2 {
				exteriorSides++
				continue
			}
			water[cc] = true
			queue = append(queue, cc)
		}
	}
	fmt.Println("Part 2:", exteriorSides)
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
