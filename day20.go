package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"github.com/gammazero/deque"
)

type pair struct {
	x, y int
}

func main() {
	var data []int
	scanner := bufio.NewScanner(os.Stdin)
	deque := deque.New[pair]()
	i := 0
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		data = append(data, val)
		deque.PushBack(pair{i, val})
		i++
	}
	n := len(data)
	mix := func() {
		for i, v := range data {
			j := deque.Index(func(a pair) bool {
				return a.x == i && a.y == v
			})
			deque.Rotate(j)
			pp := deque.PopFront()
			deque.Rotate(v)
			deque.PushFront(pp)
		}
	}
	getSum := func() int {
		ndata := make([]int, n)
		for i := 0; i < n; i++ {
			ndata[i] = deque.At(i).y
		}
		zeroPos := indexOf(ndata, 0)
		sum := 0
		for i := 1; i <= 3; i++ {
			sum += ndata[(zeroPos + i * 1000) % n]
		}
		return sum
	}
	mix()
	fmt.Println("Part 1:", getSum())
	deque.Clear()
	for i := 0; i < n; i++ {
		data[i] *= 811589153
		deque.PushBack(pair{i, data[i]})
	}
	for k := 0; k < 10; k++ {
		mix()
	}
	fmt.Println("Part 2:", getSum())
}

func indexOf[T comparable](arr []T, val T) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
