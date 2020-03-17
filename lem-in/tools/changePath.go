package tools

import "fmt"

//ChangeWay using for change founded path if it possible
func ChangeWay(arr []*Vertex, index int, graph []Vertex, stop *Stop) ([]*Vertex, bool) {
	p := fmt.Println
	if stop.Rep {
		p("try to change this")
		PrintWay(arr)
	}

	arr1 := arr[:index-1]
	// CleanGraph(graph)
	for len(arr1) > 0 {
		CleanGraph(graph)
		if stop.Rep {
			p("try to search from : ", arr[index-1].Name)
		}
		arr1[len(arr1)-1].Visited = true
		arr[index].Visited = true

		arr2, err := FindMinTrack(arr[index-1], stop)
		arr1[len(arr1)-1].Visited = false

		if !err {
			index--
			arr1 = arr[:index-1]
			continue
		}
		arr1 = append(arr1, arr2...)
		return arr1, true
	}
	return arr, false
}
