package main

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
	"strconv"
	"math"
)

type point struct {
	x, y int
}

func main() {
	re := regexp.MustCompile(`(\d+|\w)`)
	file, _ := os.Open("data/day22-2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := map[point]byte{}
	j := 0
	first := true
	var start point
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		for i, ch := range scanner.Bytes() {
			if ch != ' ' {
				p := point{i, j}
				data[p] = ch
				if first {
					start = p
					first = false
				}
			}
		}
		j++
	}
	diff := []point{ point{ 1, 0 }, point{ 0, 1 }, point{ -1, 0 }, point{ 0, -1 } }
	//dirStr := []byte{ '>', 'v', '<', '^' }
	scanner.Scan()
	matches := re.FindAllStringSubmatch(scanner.Text(), -1)
	getWrap2D := func(cur point, diffId int) (point, int) {
		dirOp := diff[(diffId + 2) % 4]
		for {
			next := point{cur.x + dirOp.x, cur.y + dirOp.y}
			_, ok := data[next]
			if !ok {
				break
			} else {
				cur = next
			}
		}
		return cur, diffId
	}

	type wrapFunc func(point, int) (point, int)
	move := func(cur point, diffId int, wrap wrapFunc) (point, int) {
		for _, match := range matches {
			val, err := strconv.Atoi(match[1])
			if err != nil {
				if match[1] == "L" {
					diffId = (diffId - 1 + 4) % 4
				}
				if match[1] == "R" {
					diffId = (diffId + 1) % 4
				}
			} else {
				dir := diff[diffId]
				for i := 0; i < val; i++ {
					next := point{cur.x + dir.x, cur.y + dir.y}
					ch, ok := data[next]
					if !ok {
						cur2, newDirId := wrap(cur, diffId)
						ch2 := data[cur2]
						if ch2 == '.' {
							cur = cur2
							diffId = newDirId
							dir = diff[diffId]
							//data[cur] = dirStr[diffId]
						} else {
							break
						}
					} else if ch == '#' {
						break
					} else if ch == '.' {
						cur = next
						//data[cur] = dirStr[diffId]
					}
				}
			}
		}
		return cur, diffId
	}
 	cur, diffId := move(start, 0, getWrap2D)
	p1 := 1000 * (cur.y + 1) + 4 * (cur.x + 1) + diffId
	fmt.Println("Part 2:", p1)

	n := len(data)
	side := int(math.Sqrt(float64(n / 6)))
	maxx, maxy := 0, 0
	for p := range data {
		maxx = max(maxx, p.x)
		maxy = max(maxy, p.y)
	}

	faceToCoord := map[int]point{}
	coordToFace := map[point]int{}
	for y := 0; y <= maxy / side; y++ {
		for x := 0; x <= maxx / side; x++ {
			p := point{x * side, y * side}
			_, ok := data[p]
			if ok {
				faceId := len(faceToCoord)
				faceToCoord[faceId] = p
				coordToFace[p] = faceId
				//fmt.Print(faceId)
			} else {
				//fmt.Print(".")
			}
		}
		//fmt.Println()
	}
	missing := 6 * 4
	joint := make([][]point, 6)
	for faceId, p := range faceToCoord {
		joint[faceId] = make([]point, 4)
		for dirId, dir := range diff {
			pp := point{p.x + side * dir.x, p.y + side * dir.y}
			nextFace, ok := coordToFace[pp]
			if ok {
				joint[faceId][dirId] = point{nextFace, dirId}
				missing--
			} else {
				joint[faceId][dirId] = point{-1, -1}
			}
		}
	}
	for missing > 0 {
		for faceId := 0; faceId < 6; faceId++ {
			for dirId := 0; dirId < 4; dirId++ {
				if joint[faceId][dirId].x != -1 {
					continue
				}
				next := joint[faceId][(dirId + 3) % 4]
				if next.x == -1 {
					continue
				}
				nextFaceId, nextDirId := next.x, next.y
				next2 := joint[nextFaceId][(nextDirId + 1) % 4]
				if next2.x == -1 {
					continue
				}
				nextFaceId, nextDirId = next2.x, next2.y
				joint[faceId][dirId] = point{nextFaceId, (nextDirId + 3) % 4}
				missing--
			}
		}
	}
	//for fid, face := range joint {
	//	for did, d := range face {
	//		dirs := ">v<^"
	//		fmt.Println(fid, string(dirs[did]), "--", string(dirs[d.y]), d.x)
	//	}
	//}
	//fmt.Println(n, side, maxx / side + 1, maxy / side + 1)
	type pointDir struct {
		p point
		d int
	}
	glue := map[pointDir]pointDir{}
	for fid, face := range joint {
		fp1 := faceToCoord[fid]
		for did, d := range face {
			fp2 := faceToCoord[d.x]
			f1 := []point{}
			for i := 0; i < side; i++ {
				p := point{0, 0}
				if did == 0 {
					p = point{side - 1, i}
				} else if did == 1 {
					p = point{side - i - 1, side - 1}
				} else if did == 2 {
					p = point{0, side - i - 1}
				} else if did == 3 {
					p = point{i, 0}
				}
				f1 = append(f1, point{p.x + fp1.x, p.y + fp1.y})
			}
			f2 := []point{}
			for i := 0; i < side; i++ {
				p := point{0, 0}
				if d.y == 0 {
					p = point{0, i}
				} else if d.y == 1 {
					p = point{side - i - 1, 0}
				} else if d.y == 2 {
					p = point{side - 1, side - i - 1}
				} else if d.y == 3 {
					p = point{i, side - 1}
				}
				f2 = append(f2, point{p.x + fp2.x, p.y + fp2.y})
			}
			for i := 0; i < side; i++ {
				a := pointDir{ f1[i], did }
				b := pointDir{ f2[i], d.y }
				glue[a] = b
			}
		}
	}

	getWrap3D := func(cur point, diffId int) (point, int) {
		wrapPoint := pointDir{cur, diffId}
		telep := glue[wrapPoint]
		return telep.p, telep.d
	}
 	cur, diffId = move(start, 0, getWrap3D)
	p2 := 1000 * (cur.y + 1) + 4 * (cur.x + 1) + diffId
	fmt.Println("Part 2:", p2)
	//for j := 0; j <= maxy; j++ {
	//	for i := 0; i <= maxx; i++ {
	//		if i == 10 && j == 3 {
	//			fmt.Print("x")
	//			continue
	//		}
	//		if i == cur.x && j == cur.y {
	//			fmt.Print("x")
	//			continue
	//		}
	//		ch, ok := data[point{i, j}]
	//		if !ok {
	//			fmt.Print(" ")
	//		} else {
	//			fmt.Print(string(ch))
	//		}
	//	}
	//	fmt.Print("\n")
	//}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
