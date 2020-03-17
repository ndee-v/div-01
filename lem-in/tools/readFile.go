package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//ReadArgs - reading arguments, then  readin file and creating Graph
func ReadArgs(arr []string) (Data, Stop) {
	p := fmt.Println
	ex := os.Exit
	if len(arr) == 0 {
		PrintToTerm("./help.txt")
		ex(0)
	}
	var (
		rep, total, clear bool
	)
	for i := 0; i < len(arr); {
		arr[i] = strings.ToLower(arr[i])
		if arr[i][0] == '-' {
			if len(arr[i]) != 1 {
				if arr[i][1] == '-' {
					switch arr[i][2:] {
					case "help":
						PrintToTerm("./help.txt")
						ex(0)
					case "total":
						total = true
					case "clear":
						clear = true
					case "report":
						rep = true
					case "about":
						PrintToTerm("./about.txt")
						ex(0)
					default:
						p("Error! incorrect flag: ", arr[i][2:], ";")
						PrintToTerm("./help.txt")
						ex(0)
					}
				} else {
					for _, v := range arr[i][1:] {
						switch v {
						case 'h':
							PrintToTerm("./help.txt")
							ex(0)
						case 't':
							total = true
						case 'c':
							clear = true
						case 'r':
							rep = true
						case 'a':
							PrintToTerm("./about.txt")
							ex(0)
						default:
							p("Error! incorrect flag: ", string(v), ";")
							PrintToTerm("./help.txt")
							os.Exit(0)
						}
					}
				}
			}
			arr = append(arr[:i], arr[i+1:]...)
			i-- // form the remove item index to start iterate next item
			continue
		}
		i++
	}
	if len(arr) == 0 {
		PrintToTerm("/help.txt")
		ex(0)
	}

	var byteArr []byte
	var e error
	var strArr []string
	temp := strings.Split(arr[0], "/")
	if len(temp) > 1 {
		if strings.Index(temp[0], "http") != -1 {
			arr[0] = strings.ReplaceAll(arr[0], "github", "raw.githubusercontent")
			arr[0] = strings.ReplaceAll(arr[0], "/blob/master/", "/master/")
			f, _ := exec.Command("curl", arr[0]).Output()
			strArr = strings.Split(string(f), "\n")
		} else {
			byteArr, e = ioutil.ReadFile(arr[0])
			strArr = strings.Split(string(byteArr), "\n")
		}
	} else {
		byteArr, e = ioutil.ReadFile("./examples/" + os.Args[1])
		strArr = strings.Split(string(byteArr), "\n")
	}
	if e != nil {
		p(e.Error())
		ex(0)
	}

	var DataBase Data
	antsBool := false
	startBool := false
	endBool := false
	if !clear {
		for _, v := range strArr {
			p(v)
		}
		p()
	}

	for i := 0; i < len(strArr); {
		if strArr[i] == "" {
			i++
			continue
		}
		if !antsBool {
			ants, e := strconv.Atoi(strArr[i])
			if e != nil {
				if rep {
					p("no ants found")
				}
				p("ERROR: invalid data format")
				ex(0)
			}
			DataBase.Ants = ants
			antsBool = true
			i++
			continue
		}
		if strArr[i] == "##start" || strArr[i] == "##end" {

			if startBool && strArr[i] == "##start" {
				if rep {
					p("start rooms more than one")
				}
				p("ERROR: invalid data format")
				ex(0)
			}
			if endBool && strArr[i] == "##end" {
				if rep {
					p("end rooms more than one")
				}
				p("ERROR: invalid data format")
				ex(0)
			}
			if i == len(strArr)-1 {
				if rep {
					p("no more strings after ", strArr[i], ";")
				}
				p("ERROR: invalid data format")
				ex(0)
			}
			tempArr := strings.Split(strArr[i+1], " ")
			if len(tempArr) != 3 {
				if rep {
					p("incorrect ", strArr[i], " room params: ", tempArr, ";")
				}
				p("ERROR: invalid data format")
				ex(0)
			}
			V := &Vertex{Name: tempArr[0], Type: strArr[i][2:], X: tempArr[1], Y: tempArr[2]}
			DataBase.Graph = append(DataBase.Graph, *V)

			if strArr[i] == "##start" {
				startBool = true
			}
			if strArr[i] == "##end" {
				endBool = true
			}
			i += 2
			continue
		}

		tempArr := strings.Split(strArr[i], " ")
		if len(tempArr) == 3 {
			V := &Vertex{Name: tempArr[0], Type: "room", X: tempArr[1], Y: tempArr[2]}
			DataBase.Graph = append(DataBase.Graph, *V)
			i++
			continue
		}
		if len(tempArr) != 1 {
			i++
			continue
		}
		tunnel := strings.Split(tempArr[0], "-")
		if len(tunnel) != 2 {
			//commets like ##comment can be here
			if tunnel[0][0] == '#' {
				if rep {
					p("comment: ", tunnel[0], ";")
				}
			} else {
				if rep {
					p("incorrect tunnel", tunnel, ";")
				}
			}
			i++
			continue
		}
		if tunnel[0] == tunnel[1] {
			if rep {
				p("tunnel type is loop", tunnel, ";")
			}
			i++
			continue
		}
		for index := range DataBase.Graph {
			if DataBase.Graph[index].Name == tunnel[0] {
				for index2 := range DataBase.Graph {
					if DataBase.Graph[index2].Name == tunnel[1] {
						//check for dublicate tunnels
						if uniq(DataBase.Graph[index].Next, &DataBase.Graph[index2]) {
							DataBase.Graph[index].Next = append(DataBase.Graph[index].Next, &DataBase.Graph[index2])
							DataBase.Graph[index2].Next = append(DataBase.Graph[index2].Next, &DataBase.Graph[index])
						}
					}
				}
			}
		}
		i++
	}
	if !antsBool || !startBool || !endBool {
		if rep {
			p("incorrect input: no number of ants or no start room or no end room", antsBool, startBool, endBool)
		}
		p("ERROR: invalid data format")
		ex(0)
	}
	startV := &Vertex{}
	endV := &Vertex{}
	for i := range DataBase.Graph {
		switch DataBase.Graph[i].Type {
		case "start":
			startV = &DataBase.Graph[i]
		case "end":
			endV = &DataBase.Graph[i]
		}
	}
	stop := Stop{Stop: false, Ants: DataBase.Ants, Start: startV, End: endV, Rep: rep, Total: total, Clear: clear, NameOfFile: arr[0]}
	if len(startV.Next) > len(endV.Next) {
		stop.MaxWays = len(endV.Next)
	} else {
		stop.MaxWays = len(startV.Next)
	}
	return DataBase, stop
}

//uniq needs for checking duplicates vertex
func uniq(arr []*Vertex, i *Vertex) bool {
	for _, v := range arr {
		if v == i {
			return false
		}
	}
	return true
}
