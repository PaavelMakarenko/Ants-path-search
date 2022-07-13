package main

import (
	"fmt"
	"os"

	"lem-in/fileprep"
	"lem-in/path"
	st "lem-in/structs"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("ERROR. Invalid number of arguments")
		os.Exit(0)
	}
	status := fileprep.ReadInput(os.Args[1])
	if !status {
		g := path.NewGraph()

		if !path.FuncAddVertexAndEdge(g) {
			v := g.CreateVisited()

			allPossiblePaths := path.FindAllPaths(st.StartRoom, st.EndRoom, g, v)
			if len(allPossiblePaths) == 0 {
				fmt.Println("ERROR. No connections between start and end rooms")
				os.Exit(0)
			}

			keysOfCombinations, allPaths := path.SortPaths(allPossiblePaths)

			rightPath := path.BestPath(keysOfCombinations, allPaths)

			// Sending ants to right paths
			fmt.Println(st.File)
			path.SendAnts(rightPath)
		} else {
			fmt.Println("ERROR. Invalid data format")
		}

	}

}
