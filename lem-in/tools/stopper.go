package tools

import "fmt"

//MoreWays func reads graph and ways and cheking for find more ways
func MoreWays(ways [][]*Vertex, stop *Stop) bool {
	p := fmt.Println
	if len(ways) > 0 {
		ants := stop.Ants
		arrQueue := make([]int, len(ways))
		for ants != 0 {
			for i := 0; i < len(ways); {
				if ants == 0 {
					break
				}
				ants--
				arrQueue[i]++
				if i != len(ways)-1 {
					if arrQueue[i]+len(ways[i])-1 > arrQueue[i+1]+len(ways[i+1])-1 {
						i++
						continue
					}
				} else {
					i = 0
					continue
				}
				continue
			}
		}
		minSteps := arrQueue[0] + len(ways[0][1:]) - 1

		if !stop.Stop {
			stop.Steps = minSteps
			stop.Stop = true
		}
		if minSteps < stop.Steps {
			stop.Steps = minSteps
		}
		if stop.Rep {
			for i, v := range arrQueue {
				p("queue for way #", i+1, "len of queue : ", v, ";")
			}
		}

		for i, v := range arrQueue {
			if v == 1 {
				if stop.Rep {
					p("most small queue found")
				}
				stop.Steps = len(ways[i])
				return false
			}
		}

		if stop.Steps == len(ways[0][1:])*stop.Ants {
			stop.Steps = 1
			if stop.Rep {
				p("shortest way found")
			}
			return false
		}
	}
	if stop.Rep {
		if stop.Stop {
			p("at this time min steps = ", stop.Steps)
		}
	}

	count := 0
	for i := range ways {
		if ways[i] != nil {
			count++
		}
	}
	if count == stop.MaxWays {
		if stop.Rep {
			p("amount of ways == max  ", count, ":", stop.MaxWays)
		}
		return false
	}
	if stop.Rep {
		p("try to find more ways")
	}
	return true
}
