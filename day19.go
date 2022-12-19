package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"fmt"
)

type cost struct {
	ore, clay, obsi int
}

func (a cost) addCost(b cost) cost {
	return cost{ a.ore + b.ore, a.clay + b.clay, a.obsi + b.obsi }
}

func (a cost) subCost(b cost) cost {
	return cost{ a.ore - b.ore, a.clay - b.clay, a.obsi - b.obsi }
}

func (a cost) addInt(ore, clay, obsi int) cost {
	return cost{ a.ore + ore, a.clay + clay, a.obsi + obsi }
}

type blueprint struct {
	id int
	robotOre, robotClay, robotObsi, robotGeod cost
}

type state struct {
	time int
	res, robots cost
}

func main() {
	re := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	scanner := bufio.NewScanner(os.Stdin)
	var blueprints []blueprint
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		var bp blueprint
		bp.id, _ = strconv.Atoi(match[1])
		bp.robotOre.ore, _ = strconv.Atoi(match[2])
		bp.robotClay.ore, _ = strconv.Atoi(match[3])
		bp.robotObsi.ore, _ = strconv.Atoi(match[4])
		bp.robotObsi.clay, _ = strconv.Atoi(match[5])
		bp.robotGeod.ore, _ = strconv.Atoi(match[6])
		bp.robotGeod.obsi, _ = strconv.Atoi(match[7])
		blueprints = append(blueprints, bp)
	}
	getMaxGeode := func(bp blueprint, initTime int) int {
		maxOre := 0
		maxOre = max(maxOre, bp.robotOre.ore)
		maxOre = max(maxOre, bp.robotClay.ore)
		maxOre = max(maxOre, bp.robotObsi.ore)
		maxOre = max(maxOre, bp.robotGeod.ore)
		cache := map[state]int{}
		var getGeode func(time int, res cost, robots cost) int
		getGeode = func(time int, res cost, robots cost) int {
			s := state{ time, res, robots }
			r, ok := cache[s]
			if ok {
				return r
			}
			if time == 0 {
				cache[s] = 0
				return 0
			}
			newRes := res.addCost(robots)
			if res.obsi >= bp.robotGeod.obsi && res.ore >= bp.robotGeod.ore {
				// always build geode robot if it's possible
				r = (time - 1) + getGeode(time - 1, newRes.subCost(bp.robotGeod), robots)
			} else {
				built := false
				if res.ore >= bp.robotOre.ore {
					built = true
					r = max(r, getGeode(time - 1, newRes.subCost(bp.robotOre), robots.addInt(1, 0, 0)))
				}
				if res.ore >= bp.robotClay.ore {
					built = true
					r = max(r, getGeode(time - 1, newRes.subCost(bp.robotClay), robots.addInt(0, 1, 0)))
				}
				if res.ore >= bp.robotObsi.ore && res.clay >= bp.robotObsi.clay {
					built = true
					r = max(r, getGeode(time - 1, newRes.subCost(bp.robotObsi), robots.addInt(0, 0, 1)))
				}
				// check "do nothing" only if:
				// 1. there is really nothing to build on this step
				// 2: if after producing Ore robot there will be not enough ore to build any robot on the next step
				if !built || newRes.subCost(bp.robotOre).ore < maxOre {
					r = max(r, getGeode(time - 1, newRes, robots))
				}
			}
			cache[s] = r
			return r
		}
		return getGeode(initTime, cost{0, 0, 0}, cost{1, 0, 0})
	}
	sumEff := 0
	for _, bp := range blueprints {
		sumEff += bp.id * getMaxGeode(bp, 24)
	}
	fmt.Println("Part 1:", sumEff)
	mult := 1
	for i, bp := range blueprints {
		if i >= 3 {
			break
		}
		mult *= getMaxGeode(bp, 32)
	}
	fmt.Println("Part 2:", mult)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
