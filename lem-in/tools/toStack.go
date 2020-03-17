package tools

import "fmt"

//ToStack func using for adding founded new path to all founded paths
func ToStack(stack [][]*Vertex, way []*Vertex, graph []Vertex, stop *Stop) ([][]*Vertex, bool) {
	p := fmt.Println
	for index := range stack {
		for i := 1; i < len(stack[index])-1; i++ {
			for j := 1; j < len(way)-1; j++ {
				if stack[index][i] == way[j] {
					if stop.Rep {
						p("duplicate vertex: -->", way[j].Name, "<--")
						PrintWay(stack[index])
						PrintWay(way)
						p("----------")
					}
					temp, e := ChangeWay(way, j, graph, stop)
					if !e {
						if stop.Rep {
							p("can't change way, trying to change before added way")
						}
						temp2, e2 := ChangeWay(stack[index], i, graph, stop)
						if !e2 {
							if stop.Rep {
								p("can't to change before added way")
							}

							return stack, false
						}
						if stop.Rep {
							p("way was changed to:")
							PrintWay(temp2)
							p("___________")
							p("saving old way to temp")
						}
						tempArr := stack[index]
						if stop.Rep {
							PrintWay(tempArr)
							p("delete this way from stack")
						}
						stack[index] = nil
						if stop.Rep {
							p("try to add new(changed) way to stack")
						}
						stack, er := ToStack(stack, temp2, graph, stop)
						if !er {
							if stop.Rep {
								p("Can't to add new way to stack")
								p("Moving back temp way to stack")
							}
							stack[index] = tempArr
							return stack, false
						}
						if stop.Rep {
							p("changed way added to stack")
							p("try to add new way to stack")
						}
						stack, er2 := ToStack(stack, way, graph, stop)
						if !er2 {
							if stop.Rep {
								p("Can't add new way to changed stack")
								p("return false")
							}
							return stack, false
						}
						if stop.Rep {
							p("new way added to changed stack")
						}
						return stack, true
					}
					if stop.Rep {
						p("way changed to:")
						PrintWay(temp)
						p("_______________")
					}
					stack, er := ToStack(stack, temp, graph, stop)
					if !er {
						// stack[index] = tempArr
						if stop.Rep {
							p("Can not add changed way")
						}
						return stack, false

					}
					return stack, true
				}
			}
		}
	}
	stack = append(stack, way)
	if stop.Rep {
		PrintWay(way)
		p("^appended")
		p("at this time stack has his ways:")
		for i := range stack {
			PrintWay(stack[i])
		}
		p("^^^^^^^^")
	}
	return stack, true
}
