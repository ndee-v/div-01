package main

import (
	"fmt"
	"os"

	tools "./tools"
)

func main() {
	p := fmt.Println
	//opening txt file
	dataBase, stop := tools.ReadArgs(os.Args[1:])
	if stop.Rep {
		p("total of Vertex :", len(dataBase.Graph))
		p("ants amount: ", dataBase.Ants)
	}

	ways := [][]*tools.Vertex{}
	for tools.MoreWays(ways, &stop) {
		tools.CleanGraph(dataBase.Graph)
		way, e := tools.FindMinTrack(stop.Start, &stop)
		if e {
			if stop.Rep {
				p("new way found")
				tools.PrintWay(way)
				p("try to add this way to stack of ways")
			}
			ways, e = tools.ToStack(ways, way, dataBase.Graph, &stop)
			if !e {
				break
			}
			ways = tools.Filter(ways)
		} else {
			if stop.Rep {
				p("no ways found")
			}
			break
		}
	}

	ways = tools.Filter(ways)
	if stop.Rep {
		p("=====total ways after filtering===")
		for i := range ways {
			p("way# ", i+1, " lenght: ", len(ways[i][1:]), ";")
			tools.PrintWay(ways[i])
		}
		p("=====total ways===:", len(ways), "; minimum Steps: ", stop.Steps, ";")
		p("ready to open gates!", "\n")
	}

	tools.OpenGates(ways, &stop)

}
