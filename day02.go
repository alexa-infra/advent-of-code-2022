package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	rock = 1
	paper = 2
	scissor = 3
	win = 6
	draw = 3
	lose = 0
)

func play(op, me int) int {
	if op == me {
		return draw
	}
	if op == rock {
		if me == scissor {
			return lose
		}
		return win
	}
	if op == paper {
		if me == rock {
			return lose
		}
		return win
	}
	if op == scissor {
		if me == paper {
			return lose
		}
		return win
	}
	return 0
}

func needs(op, outcome int) int {
	if outcome == draw {
		return op
	}
	if outcome == win {
		if op == rock {
			return paper
		}
		if op == paper {
			return scissor
		}
		if op == scissor {
			return rock
		}
	}
	if outcome == lose {
		if op == rock {
			return scissor
		}
		if op == paper {
			return rock
		}
		if op == scissor {
			return paper
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalScore := 0
	totalScore2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		var op, me int
		if line[0] == 'A' {
			op = rock
		} else if line[0] == 'B' {
			op = paper
		} else if line[0] == 'C' {
			op = scissor
		}
		if line[2] == 'X' {
			me = rock
		} else if line[2] == 'Y' {
			me = paper
		} else if line[2] == 'Z' {
			me = scissor
		}
		totalScore += me + play(op, me)

		var out int
		if line[2] == 'X' {
			out = lose
		} else if line[2] == 'Y' {
			out = draw
		} else if line[2] == 'Z' {
			out = win
		}
		totalScore2 += needs(op, out) + out
	}
	fmt.Println("Part 1:", totalScore)
	fmt.Println("Part 2:", totalScore2)
}
