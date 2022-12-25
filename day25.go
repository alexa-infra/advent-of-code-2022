package main

import (
	"bufio"
	"os"
	"fmt"
	//"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	p1 := 0
	for scanner.Scan() {
		//n, _ := strconv.Atoi(scanner.Text())
		//str := convertSnafuBack(n)
		//fmt.Println(str, convertSnafu(str))
		n := convertSnafu(scanner.Text())
		p1 += n
	}
	fmt.Println("Part 1:", convertSnafuBack(p1))
}

func convertSnafu(str string) int {
	alpha := map[byte]int{ '0': 0, '1': 1, '2': 2, '-': -1, '=': -2 }
	n := 1
	r := 0
	for i := len(str) - 1; i >= 0; i-- {
		ch := str[i]
		val := alpha[ch]
		r += val * n
		n *= 5
	}
	return r
}

func convertSnafuBack(val int) string {
	alpha := map[int]byte{ 0: '0', 1: '1', 2: '2', -1: '-', -2: '=' }
	r := []byte{}
	for val > 0 {
		i, nval := val % 5, val / 5
		if i == 4 {
			nval++
			i = -1
		} else if i == 3 {
			nval++
			i = -2
		}
		r = append([]byte{alpha[i]}, r...)
		val = nval
	}
	return string(r)
}
