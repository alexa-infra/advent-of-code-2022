package main

import (
	"bufio"
	"os"
	"regexp"
	"fmt"
	"strconv"
)

type monkey struct {
	name string
	hasValue bool
	value int
	left, right string
	op string
}

func NewMonkey(name string) *monkey {
	return &monkey{name, false, 0, "", "", ""}
}

func (m *monkey) SetValue(value int) {
	m.hasValue = true
	m.value = value
	m.left = ""
	m.right = ""
	m.op = ""
}

func (m *monkey) SetOp(left, right string, op string) {
	m.hasValue = false
	m.value = 0
	m.left = left
	m.right = right
	m.op = op
}

func main() {
	re1 := regexp.MustCompile(`^(\w+): (\w+) (.) (\w+)$`)
	re2 := regexp.MustCompile(`^(\w+): (\d+)$`)
	scanner := bufio.NewScanner(os.Stdin)
	monkeys := map[string]*monkey{}
	getMonkey := func(name string)*monkey {
		mon, ok := monkeys[name]
		if !ok {
			mon = NewMonkey(name)
			monkeys[name] = mon
		}
		return mon
	}
	for scanner.Scan() {
		line := scanner.Text()
		if re1.MatchString(line) {
			match := re1.FindStringSubmatch(line)
			name1, name2, opName, name3 := match[1], match[2], match[3], match[4]
			mon1 := getMonkey(name1)
			mon2 := getMonkey(name2)
			mon3 := getMonkey(name3)
			mon1.SetOp(mon2.name, mon3.name, opName)
		} else if re2.MatchString(line) {
			match := re2.FindStringSubmatch(line)
			name, value := match[1], match[2]
			intVal, _ := strconv.Atoi(value)
			mon := getMonkey(name)
			mon.SetValue(intVal)
		}
	}
	var resolve func(string) int
	resolve = func(name string) int {
		m := monkeys[name]
		if m.hasValue {
			return m.value
		}
		v, a, b := 0, resolve(m.left), resolve(m.right)
		if m.op == "+" {
			v = a + b
		} else if m.op == "*" {
			v = a * b
		} else if m.op == "-" {
			v = a - b
		} else if m.op == "/" {
			v = a / b
		}
		return v
	}
	fmt.Println("Part 1:", resolve("root"))
	findDirect := func(name string)*monkey {
		m, ok := monkeys[name]
		if ok {
			return m
		}
		return nil
	}
	findIndirect := func(name string, except string) *monkey {
		for _, m := range monkeys {
			if !m.hasValue && m.name != except {
				if m.left == name || m.right == name {
					return m
				}
			}
		}
		return nil
	}
	reverseFirst := func(m *monkey)*monkey {
		if m.op == "*" {
			return &monkey{m.left, false, 0, m.name, m.right, "/"}
		}
		if m.op == "/" {
			return &monkey{m.left, false, 0, m.name, m.right, "*"}
		}
		if m.op == "-" {
			return &monkey{m.left, false, 0, m.name, m.right, "+"}
		}
		if m.op == "+" {
			return &monkey{m.left, false, 0, m.name, m.right, "-"}
		}
		return nil
	}
	reverseSecond := func(m *monkey)*monkey {
		if m.op == "*" {
			return &monkey{m.right, false, 0, m.name, m.left, "/"}
		}
		if m.op == "/" {
			return &monkey{m.right, false, 0, m.left, m.name, "/"}
		}
		if m.op == "-" {
			return &monkey{m.right, false, 0, m.left, m.name, "-"}
		}
		if m.op == "+" {
			return &monkey{m.right, false, 0, m.name, m.left, "-"}
		}
		return nil
	}
	var last string
	searchName := "humn"
	for {
		delete(monkeys, searchName)
		node := findIndirect(searchName, last)
		if node.name == "root" {
			var node2 *monkey
			if node.left == searchName {
				node2 = findDirect(node.right)
			}
			if node.right == searchName {
				node2 = findDirect(node.left)
			}
			delete(monkeys, node2.name)
			monkeys[searchName] = node2
			node2.name = searchName
			break
		}
		var rnode *monkey
		if node.left == searchName {
			rnode = reverseFirst(node)
		}
		if node.right == searchName {
			rnode = reverseSecond(node)
		}
		delete(monkeys, node.name)
		monkeys[searchName] = rnode
		last = rnode.name
		searchName = node.name
	}
	fmt.Println("Part 2:", resolve("humn"))
}
