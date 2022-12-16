package main

import (
	"bufio"
	"os"
	"strconv"
	"regexp"
	"strings"
	"fmt"
)

type node struct {
	name string
	rate int
	leads []*node
}

func main() {
	re := regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)
	scanner := bufio.NewScanner(os.Stdin)
	nodes := map[string]*node{}
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		name := match[1]
		rate, _ := strconv.Atoi(match[2])
		leads := strings.Split(match[3], ", ")
		curr, ok := nodes[name]
		if !ok {
			curr = &node{ name, rate, []*node{} }
			nodes[name] = curr
		} else {
			curr.rate = rate
		}
		for _, lead := range leads {
			next, ok := nodes[lead]
			if !ok {
				next = &node{ lead, 0, []*node{} }
				nodes[lead] = next
			}
			curr.leads = append(curr.leads, next)
		}
	}
	nonZero := []string{}
	for name, n := range nodes {
		if n.rate > 0 {
			nonZero = append(nonZero, name)
		}
	}
	type state struct {
		path []string
		time int
		pressure int
	}
	cache := map[string]int{}
	dijkstraCached := func (source, target string) (dist int) {
		key := source + target
		dist, ok := cache[key]
		if !ok {
			dist = dijkstra(nodes, source, target)
			cache[key] = dist
		}
		return dist
	}
	searchPath := func(steps int, nameSet []string) (maxPressure int) {
		states := []state{ state{ []string{"AA"}, steps, 0 } }
		for len(states) > 0 {
			current := states[0]
			states = states[1:]
			if current.pressure > maxPressure {
				maxPressure = current.pressure
			}
			for _, name := range nameSet {
				if !contains(current.path, name) {
					last := current.path[len(current.path)-1]
					dist := dijkstraCached(last, name)
					time := current.time - dist - 1
					if time > 0 {
						pressure := current.pressure + time * nodes[name].rate
						newPath := []string{}
						newPath = append(newPath, current.path...)
						newPath = append(newPath, name)
						states = append(states, state{ newPath, time, pressure })
					}
				}
			}
		}
		return maxPressure
	}
	p1 := searchPath(30, nonZero)
	fmt.Println("Part 1:", p1)
	p2 := 0
	for _, combi := range combinations(nonZero, len(nonZero) / 2) {
		pressure1 := searchPath(26, combi)
		nonUsed := exclude(nonZero, combi)
		pressure2 := searchPath(26, nonUsed)
		if pressure1 + pressure2 > p2 {
			p2 = pressure1 + pressure2
		}
	}
	fmt.Println("Part 2:", p2)
}

func exclude(a, b []string) []string {
	ret := []string{}
	for _, x := range a {
		if !contains(b, x) {
			ret = append(ret, x)
		}
	}
	return ret
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

func removeAt(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

func combinations[T any](set []T, n int) (subsets [][]T) {
	length := len(set)
	for bits := 1; bits < (1 << length); bits++ {
		if n > 0 && countBits(bits) != n {
			continue
		}
		var subset []T
		for i := 0; i < length; i++ {
			if (bits >> i) & 1 == 1 {
				subset = append(subset, set[i])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

func countBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

func dijkstra(graph map[string]*node, source, target string) int {
	dist := map[string]int{}
	queue := []string{}
	inf := 100000
	for name := range graph {
		queue = append(queue, name)
		dist[name] = inf
	}
	dist[source] = 0
	for len(queue) > 0 {
		minName := queue[0]
		minDist := dist[minName]
		minId := 0
		for i, q := range queue {
			if dist[q] < minDist {
				minName = q
				minDist = dist[q]
				minId = i
			}
		}
		queue = removeAt(queue, minId)
		minNode := graph[minName]
		for _, lead := range minNode.leads {
			alt := dist[minName] + 1
			if alt < dist[lead.name] {
				dist[lead.name] = alt
			}
		}
	}
	return dist[target]
}
