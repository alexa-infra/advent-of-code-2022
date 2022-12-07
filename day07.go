package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	type node struct {
		name string
		size int
		isDir bool
		children []*node
		parent *node
	}
	newFile := func(name string, size int, parent *node) *node {
		return &node{ name, size, false, []*node{}, parent }
	}
	newDir := func(name string, parent *node) *node {
		return &node{ name, 0, true, []*node{}, parent }
	}
	root := newDir("", nil)
	scanner := bufio.NewScanner(os.Stdin)
	listing := false
	var currentNode *node
	currentNode = root
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				listing = false
				if parts[2] == "/" {
					currentNode = root
				} else if parts[2] == ".." {
					currentNode = currentNode.parent
				} else {
					var child *node
					for _, node := range currentNode.children {
						if node.name == parts[2] {
							child = node
							break
						}
					}
					if child == nil {
						child = newDir(parts[2], currentNode)
						currentNode.children = append(currentNode.children, child)
					}
					currentNode = child
				}
			} else if parts[1] == "ls" {
				listing = true
			}
		} else if listing {
			var child *node
			for _, node := range currentNode.children {
				if node.name == parts[1] {
					child = node
					break
				}
			}
			if parts[0] == "dir" {
				if child == nil {
					child = newDir(parts[1], currentNode)
					currentNode.children = append(currentNode.children, child)
				}
			} else {
				size, _ := strconv.Atoi(parts[0])
				if child == nil {
					child = newFile(parts[1], size, currentNode)
					currentNode.children = append(currentNode.children, child)
				}
			}
		} else {
			fmt.Println("Unknown", line)
		}
	}
	dirs := map[string]int{}
	var travel func(*node, string)int
	travel = func(cur *node, path string) int {
		size := 0
		if cur.isDir {
			for _, child := range cur.children {
				size += travel(child, path + "/" + child.name)
			}
			cur.size = size
			dirs[path] = size
		} else {
			size += cur.size
		}
		return size
	}
	travel(root, "")
	p1 := 0
	for _, v := range dirs {
		if v < 100000 {
			p1 += v
		}
	}
	fmt.Println("Part 1:", p1)
	p2 := 70000000
	freeSize := 70000000 - root.size
	for _, v := range dirs {
		if freeSize + v >= 30000000 {
			if v < p2 {
				p2 = v
			}
		}
	}
	fmt.Println("Part 2:", p2)
}
