package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type node struct {
	list     []node
	hasValue bool
	value    int
}

func makeListNode() node {
	return node{[]node{}, false, 0}
}

func makeIntNode(value int) node {
	return node{[]node{}, true, value}
}

func parse(expr []rune) (node, int) {
	buf := []rune{}
	root := makeListNode()
	idx := 0
	for idx < len(expr) {
		switch expr[idx] {
		case '[':
			sub, end := parse(expr[idx+1:])
			root.list = append(root.list, sub)
			idx += end
		case ']', ',':
			if len(buf) > 0 {
				value, _ := strconv.Atoi(string(buf))
				root.list = append(root.list, makeIntNode(value))
				buf = buf[:0]
			}
			if expr[idx] == ']' {
				return root, idx + 1
			}
		case ' ':
		default:
			buf = append(buf, expr[idx])
		}
		idx++
	}
	return root, idx
}

func cmp(left, right node) int {
	if !left.hasValue && !right.hasValue {
		for i := 0; i < max(len(left.list), len(right.list)); i++ {
			if i >= len(left.list) {
				return -1
			}
			if i >= len(right.list) {
				return 1
			}
			r := cmp(left.list[i], right.list[i])
			if r != 0 {
				return r
			}
		}
		return 0
	}
	if !left.hasValue && right.hasValue {
		n := makeListNode()
		n.list = append(n.list, right)
		return cmp(left, n)
	}
	if left.hasValue && !right.hasValue {
		n := makeListNode()
		n.list = append(n.list, left)
		return cmp(n, right)
	}
	if left.value == right.value {
		return 0
	}
	if left.value < right.value {
		return -1
	}
	return 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	p1 := 0
	i := 1
	nodes := []node{}
	for scanner.Scan() {
		line1 := scanner.Text()
		left, _ := parse([]rune(line1))
		scanner.Scan()
		line2 := scanner.Text()
		right, _ := parse([]rune(line2))
		scanner.Scan()
		if cmp(left, right) < 0 {
			p1 += i
		}
		i++
		nodes = append(nodes, left)
		nodes = append(nodes, right)
	}
	fmt.Println("Part 1:", p1)
	extra1, _ := parse([]rune("[[2]]"))
	extra2, _ := parse([]rune("[[6]]"))
	nodes = append(nodes, extra1)
	nodes = append(nodes, extra2)
	sort.Slice(nodes, func(i, j int) bool {
		if cmp(nodes[i], nodes[j]) < 0 {
			return true
		}
		return false
	})
	p2 := 1
	for i := 0; i < len(nodes); i++ {
		if cmp(nodes[i], extra1) == 0 {
			p2 *= i + 1
		}
		if cmp(nodes[i], extra2) == 0 {
			p2 *= i + 1
		}
	}
	fmt.Println("Part 2:", p2)
}
