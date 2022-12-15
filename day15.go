package main

import (
	"bufio"
	"os"
	"strconv"
	"regexp"
	"fmt"
	"sort"
)

func main() {
	type coord struct {
		x, y int
	}
	type pair struct {
		sensor, beacon coord
		dist int
	}
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	scanner := bufio.NewScanner(os.Stdin)
	data := []pair{}
	beacons := map[coord]bool{}
	scanner.Scan()
	y1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	mm, _ := strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		r1, _ := strconv.Atoi(match[1])
		r2, _ := strconv.Atoi(match[2])
		r3, _ := strconv.Atoi(match[3])
		r4, _ := strconv.Atoi(match[4])
		sensor := coord{r1, r2}
		beacon := coord{ r3, r4 }
		dist := abs(sensor.x - beacon.x) + abs(sensor.y - beacon.y)
		data = append(data, pair{ sensor, beacon, dist })
		beacons[beacon] = true
	}
	type interval struct {
		a, b int
		valid bool
	}
	intervals := []interval{}
	for _, d := range data {
		dist := d.dist
		//dist1 := abs(d.sensor.x - p.x) + abs(d.sensor.y - y1) <= dist
		//abs(d.sensor.x - p.x) <= dist - Y
		//d.sensor.x - p.x < 0 -> p.x - d.sensor.x <= dist - Y
		//                        p.x <= dist - Y + d.sensor.x
		//                        p.x > d.sensor.x
		//d.sensor.x - p.x >= 0 -> d.sensor.x - p.x <= dist - Y
		//                         p.x >= d.sensor.x - dist + Y
		//                         p.x <= d.sensor.x
		distY := abs(d.sensor.y - y1)
		a1, b1 := d.sensor.x + 1, dist - distY + d.sensor.x
		a2, b2 := d.sensor.x - dist + distY, d.sensor.x
		if a1 <= b1 {
			intervals = append(intervals, interval{a1, b1, true})
		}
		if a2 <= b2 {
			intervals = append(intervals, interval{a2, b2, true})
		}
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].a == intervals[j].a {
			return intervals[i].b < intervals[j].b
		}
		return intervals[i].a < intervals[j].a
	})
	p1 := 0
	last := 0
	for i := 0; i < len(intervals); i++ {
		curr := &intervals[i]
		valid := true
		if i != 0 && curr.a <= last {
			if curr.b > last {
				curr.a = last + 1
			} else {
				valid = false
			}
		}
		if valid {
			last = curr.b
			p1 += curr.b - curr.a + 1
		}
	}
	for b, _ := range beacons {
		if b.y == y1 {
			p1--
		}
	}
	fmt.Println("Part 1:", p1)
	test := func(p coord) bool {
		if p.x < 0 || p.x > mm || p.y < 0 || p.y > mm {
			return false
		}
		for _, d2 := range data {
			dist1 := abs(d2.sensor.x - p.x) + abs(d2.sensor.y - p.y)
			if dist1 <= d2.dist {
				return false
			}
		}
		return true
	}
	p := coord{0, 0}
	for _, d1 := range data {
		p.x = d1.sensor.x + d1.dist + 1
		p.y = d1.sensor.y
		found := false
		for p.x >= d1.sensor.x {
			p.x--
			p.y++
			if test(p) {
				found = true
				break
			}
		}
		if found {
			break
		}
		p.x = d1.sensor.x
		p.y = d1.sensor.y + d1.dist + 1
		for p.y >= d1.sensor.y {
			p.x--
			p.y--
			if test(p) {
				found = true
				break
			}
		}
		if found {
			break
		}
		p.x = d1.sensor.x - d1.dist - 1
		p.y = d1.sensor.y
		for p.y >= d1.sensor.y {
			p.x++
			p.y--
			if test(p) {
				found = true
				break
			}
		}
		if found {
			break
		}
		p.x = d1.sensor.x
		p.y = d1.sensor.y - d1.dist - 1
		for p.y >= d1.sensor.y {
			p.x++
			p.y++
			if test(p) {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	p2 := p.x * 4000000 + p.y
	fmt.Println("Part 2:", p2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
