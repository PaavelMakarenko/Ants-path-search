package path

import (
	"fmt"
	st "lem-in/structs"
	"os"
)

func SendAnts(group [][]string) {
	n := st.AntCount
	levels := n / len(group)
	if n%len(group) != 0 {
		levels++
	}

	var ants = make([]st.Ant, n+1)
	// ignore zero because ants start from 1
	ants[0].Ignore = true
	id := 0

	var lenGroup []int
	for _, r := range group {
		lenGroup = append(lenGroup, len(r))
	}

	index := 0
	smallestnum := 1000
	for i := 0; i < n; i++ {
		for in := range lenGroup {
			if lenGroup[in] < smallestnum {
				smallestnum = lenGroup[in]
				index = in
			}
		}
		lenGroup[index]++
		smallestnum = 1000
	}
	for i := range lenGroup {
		lenGroup[i] -= len(group[i])
	}

	for i := 0; i < st.AntCount; i++ {
		for i := range lenGroup {
			if lenGroup[i] > 0 && id <= st.AntCount {
				id++
				ants[id].Path = group[i]
				ants[id].RoomID = 0
				ants[id].Ignore = false
				lenGroup[i]--
				if id == n {
					break
				}

			}
		}
	}

	for _, p := range group {
		if len(p) == 1 && p[0] == st.EndRoom {
			for i := 1; i <= n; i++ {
				fmt.Print("L", i, "-", st.EndRoom, " ")
			}
			fmt.Println("")
			os.Exit(0)
		}
	}

	exit := false
	var taken = make(map[string]bool)
	for !exit {
		for id, ant := range ants {
			if ant.Ignore {
				continue
			}
			room := ant.Path[ant.RoomID]
			if taken[room] {
				fmt.Println()
				break
			}
			fmt.Print("L", id, "-", room, " ")
			if id == n {
				fmt.Println()
				if room == st.EndRoom {
					exit = true
				}
			}
			ants[id].RoomID++
			taken[ants[id].Previous] = false
			if room != st.EndRoom {
				taken[room] = true
				ants[id].Previous = room
			}
			if room == st.EndRoom {
				ants[id].Ignore = true
			}
		}
	}
}
