package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	type command struct {
		name string
		value int
	}
	cmds := []command{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if parts[0] == "noop" {
			cmds = append(cmds, command{ "noop", 0 })
		} else if parts[0] == "addx" {
			val, _ := strconv.Atoi(parts[1])
			cmds = append(cmds, command{ "addx", val })
		}
	}
	regx := 1
	n := len(cmds)
	cycle := 1
	cmdid := 0
	current := command{ "", 0 }
	counter := 0
	p1 := 0
	for {
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			p1 += cycle * regx
		}
		row := (cycle - 1) % 40
		if regx + 1 >= row && regx - 1 <= row {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if cycle == 40 || cycle == 80 || cycle == 120 || cycle == 160 || cycle == 200 || cycle == 240 {
			fmt.Println()
		}
		if counter > 0 {
			counter--
			if counter == 0 {
				if current.name == "addx" {
					regx += current.value
				}
			}
		} else {
			if cmdid >= n {
				break
			}
			current = cmds[cmdid]
			cmdid++
			if current.name == "addx" {
				counter = 1
			}
		}
		cycle++
	}
	fmt.Println()
	fmt.Println("Part 1:", p1)
}
