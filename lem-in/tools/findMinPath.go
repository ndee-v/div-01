package tools

import "fmt"

//FindMinTrack return first minimal path
func FindMinTrack(st *Vertex, stop *Stop) ([]*Vertex, bool) {
	p := fmt.Println
	stack := []*Vertex{}
	stack = append(stack, st)
	for len(stack) != 0 {
		if stop.Stop {
			if stack[0].Depth == stop.Steps-1 {
				if stop.Rep {
					p("depth is so big, long way don't need")
				}
				stack = stack[1:]
				continue
			}
		}
		if stack[0].Visited {
			stack = stack[1:]
			continue
		}
		for _, v := range stack[0].Next {
			switch v.Type {
			case "start":
				continue
			case "room":
				if stack[0].Type == "start" {
					if v.Busy {
						continue
					}
				}
				if v.Depth != 0 || v.Visited {
					continue
				}

				for _, k := range stack[0].Track {
					v.Track = append(v.Track, k)
				}
				v.Track = append(v.Track, stack[0])
				v.Depth = stack[0].Depth + 1
				stack = append(stack, v)
			case "end":
				tempArr := []*Vertex{}
				for _, k := range stack[0].Track {
					tempArr = append(tempArr, k)
				}
				tempArr = append(tempArr, stack[0])
				tempArr = append(tempArr, v)
				for i := range tempArr {
					if i == 1 {
						if tempArr[0].Type == "start" {
							tempArr[i].Busy = true
						}
						break
					}
				}
				return tempArr, true
			}
		}
		stack = stack[1:]
	}
	return []*Vertex{}, false
}
