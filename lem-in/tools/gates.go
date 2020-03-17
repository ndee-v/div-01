package tools

import (
	"fmt"
	"os"
	"strconv"
)

//OpenGates prints ants and them position at each steps
func OpenGates(ways [][]*Vertex, stop *Stop) {
	p := fmt.Println
	pf := fmt.Print
	if !stop.Stop {
		p("ERROR: invalid data format")
		os.Exit(0)
	}
	ants := stop.Ants
	antName := 1
	antQueue := make([]*Ant, ants)
	lenQueue := make([]int, len(ways))
	for ants != 0 {
		for i := 0; i < len(ways); {
			if ants == 0 {
				break
			}
			ants--
			ant := &Ant{Name: "L" + strconv.Itoa(antName), Where: ways[i][1].Type, Position: lenQueue[i], Step: 1, Way: i}

			antQueue[antName-1] = ant
			lenQueue[i]++
			antName++
			if i < len(ways)-1 {
				if lenQueue[i]+len(ways[i][1:]) > lenQueue[i+1]+len(ways[i+1][1:]) {
					i++
				}
			} else {
				i++
			}
		}
	}
	p()
	onTheWay := true
	for onTheWay {
		for _, v := range antQueue {
			switch v.Where {
			case "room":
				onTheWay = true
				if v.Position == 0 {
					pf(v.Name, "-", ways[v.Way][v.Step].Name, " ")
					v.Step++
					v.Where = ways[v.Way][v.Step].Type
				} else {
					v.Position--
				}
			case "end":
				pf(v.Name, "-", ways[v.Way][v.Step].Name, " ")
				v.Where = ""
				onTheWay = false
			}
		}
		p()
	}
	p()
	if stop.Total || stop.Rep {
		p("Total of steps: ", stop.Steps, ", number of ants: ", stop.Ants, ", number of paths : ", len(ways), ";")
	}
}
