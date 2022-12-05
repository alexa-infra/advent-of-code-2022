package main

import (
	"bufio"
	"os"
	"strings"
	"regexp"
	"strconv"
	"fmt"
)

func main() {
	re := regexp.MustCompile(`move (\d+) from (\d) to (\d)`)
	scanner := bufio.NewScanner(os.Stdin)
	stacks := map[int][]byte{}
	stacks2 := map[int][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "[") >= 0 {
			last := 0
			for {
				idx := strings.Index(line[last:], "[")
				if idx < 0 {
					break
				}
				cargo := line[last + idx + 1]
				stackIdx := (last + idx) / 4 + 1
				stack, ok := stacks[stackIdx]
				if ok {
					newStack := []byte{cargo}
					stacks[stackIdx] = append(newStack, stack...)
				} else {
					stacks[stackIdx] = []byte{cargo}
				}
				last += idx + 1
			}
		} else if strings.Index(line, "move") >= 0 {
			if len(stacks2) == 0 {
				for k, v := range stacks {
					newStack := []byte{}
					stacks2[k] = append(newStack, v...)
				}
			}
			match := re.FindStringSubmatch(line)
			count, _ := strconv.Atoi(match[1])
			src, _ := strconv.Atoi(match[2])
			dst, _ := strconv.Atoi(match[3])
			srcStack := stacks[src]
			dstStack := stacks[dst]
			for i := 0; i < count; i++ {
				cargo := srcStack[len(srcStack)-1]
				srcStack = srcStack[:len(srcStack)-1]
				dstStack = append(dstStack, cargo)
			}
			stacks[src] = srcStack
			stacks[dst] = dstStack

			srcStack = stacks2[src]
			dstStack = stacks2[dst]
			stacks2[dst] = append(dstStack, srcStack[len(srcStack) - count:]...)
			stacks2[src] = srcStack[:len(srcStack)-count]
		}
	}
	p1 := []byte{}
	for i := 1; i <= len(stacks); i++ {
		stack := stacks[i]
		p1 = append(p1, stack[len(stack)-1])
	}
	fmt.Println("Part 1:", string(p1))
	p2 := []byte{}
	for i := 1; i <= len(stacks2); i++ {
		stack := stacks2[i]
		p2 = append(p2, stack[len(stack)-1])
	}
	fmt.Println("Part 2:", string(p2))
}
