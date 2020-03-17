package tools

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//CleanGraph using for cleaning vertex before using algorythm
func CleanGraph(data []Vertex) {
	for i := range data {
		// data[i].Visited = false
		data[i].Depth = 0
		// data[i].Prev = []*Vertex{}
		data[i].Track = []*Vertex{}
	}
}

//PrintWay using for image path in terminal window
func PrintWay(arr []*Vertex) {
	for i, v := range arr {
		fmt.Print(v.Name)
		if i != len(arr)-1 {
			fmt.Print("-")
		} else {
			fmt.Println(".")
		}
	}
}

//PrintToTerm using for print text file to terminal window
func PrintToTerm(str string) {
	temp := strings.Split(str, "/")
	var byteArr []byte
	var e error
	if len(temp) > 1 {
		byteArr, e = ioutil.ReadFile(str)
	} else {
		byteArr, e = ioutil.ReadFile("./examples/" + str)
	}
	if e == nil {
		fmt.Print(string(byteArr))
	} else {
		fmt.Println(e.Error())
	}
}

//Filter using for sorting path from min to max by lenght and deleting empty paths
func Filter(arr [][]*Vertex) [][]*Vertex {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for k := 0; k < len(arr); k++ {
			if i == k {
				continue
			}
			if len(arr[i]) < len(arr[k]) {
				temp := arr[i]
				arr[i] = arr[k]
				arr[k] = temp
			}
		}
	}
	for len(arr[0]) == 0 {
		arr = arr[1:]
	}
	return arr
}
