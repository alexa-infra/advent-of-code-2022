package main

import (
	"bufio"
	"os"
	"fmt"
)

type coord struct {
	x, y int
}

type figure []coord

type figures []figure

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dirs := scanner.Text()

	elems := figures{
		figure{ coord{ 0, 0 }, coord{ 1, 0 }, coord{ 2, 0 }, coord{ 3, 0 } },
		figure{ coord{ 1, 1 }, coord{ 2, 1 }, coord{ 0, 1 }, coord{ 1, 2 }, coord{ 1, 0 } },
		figure{ coord{ 0, 0 }, coord{ 1, 0 }, coord{ 2, 0 }, coord{ 2, 1 }, coord{ 2, 2 } },
		figure{ coord{ 0, 0 }, coord{ 0, 1 }, coord{ 0, 2 }, coord{ 0, 3 } },
		figure{ coord{ 0, 0 }, coord{ 1, 0 }, coord{ 0, 1 }, coord{ 1, 1 } },
	}
	data := map[coord]bool{}
	touches := func(elem figure, p coord) bool {
		for _, c := range elem {
			cc := coord{ p.x + c.x, p.y + c.y }
			if cc.x < 0 || cc.x > 6 {
				return true
			}
			if cc.y < 0 {
				return true
			}
			_, ok := data[cc]
			if ok {
				return true
			}
		}
		return false
	}
	maxY := 0
	place := func(elem figure, p coord) {
		for _, c := range elem {
			cc := coord{ p.x + c.x, p.y + c.y }
			data[cc] = true
			if cc.y + 1 > maxY {
				maxY = cc.y + 1
			}
		}
	}
	dirId := 0
	elemId := 0
	type state struct {
		dirId int
		elemId int
		hash string
	}
	cache := map[state]int{}
	getRow := func(rowId, n int) string {
		dat := make([]byte, n * 7)
		id := 0
		for j := rowId; j > rowId - n; j-- {
			for i := 0; i < 7; i++ {
				_, ok := data[coord{i, j}]
				if ok {
					dat[i + id * 7] = '#'
				} else {
					dat[i + id * 7] = '.'
				}
			}
			id++
		}
		return string(dat)
	}
	found := false
	cacheHeight := map[int]int{}
	p2 := int64(0)
	for i := 1; i < 5000; i++ {
		if !found {
			str := getRow(maxY - 1, 15)
			s := state{ dirId, elemId, str }
			prevId, ok := cache[s]
			if ok {
				prevHeight := cacheHeight[prevId]
				fmt.Println("cycle length", i - prevId, "height", maxY - prevHeight, "start", prevId, "initHeight", prevHeight)
				iter, offset := divmod(int64(1000000000000) - int64(prevId), int64(i - prevId))
				offsetHeight := cacheHeight[prevId + int(offset)]
				p2 = iter * int64(maxY - prevHeight) + int64(offsetHeight)

				found = true
			}
			cache[s] = i
		}
		elem := elems[elemId]
		elemId = (elemId + 1) % len(elems)
		p := coord{ 2, maxY + 3 }
		for {
			dir := dirs[dirId]
			dirId = (dirId + 1) % len(dirs)
			if dir == '>' && !touches(elem, coord{p.x + 1, p.y}) {
				p.x++
			}
			if dir == '<' && !touches(elem, coord{p.x - 1, p.y}) {
				p.x--
			}
			if !touches(elem, coord{p.x, p.y - 1}) {
				p.y--
			} else {
				place(elem, p)
				break
			}
		}
		cacheHeight[i] = maxY
	}
	fmt.Println("Part 1:", cacheHeight[2022])
	fmt.Println("Part 2:", p2)
	//for y := maxY; y >= 0; y-- {
	//	for x := 0; x < 7; x++ {
	//		p := coord{ x, y }
	//		_, ok := data[p]
	//		if ok {
	//			fmt.Print("#")
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Println()
	//}
}

func divmod(numerator, denominator int64) (quotient, remainder int64) {
    quotient = numerator / denominator
    remainder = numerator % denominator
    return
}
