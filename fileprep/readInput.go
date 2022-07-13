package fileprep

import (
	"bufio"
	"fmt"
	st "lem-in/structs"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) bool {
	startFlag := false
	endFlag := false
	firstLine := true
	errorStatus := false
	countEndStartRooms := 0
	startConnections := 0
	endConnections := 0

	openedFile, err := os.Open("./testing/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer openedFile.Close()

	scanner := bufio.NewScanner(openedFile)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		chekingString := scanner.Text()
		st.File += chekingString + "\n"

		// First line always with the number of ants
		if firstLine {
			res, err := strconv.Atoi(chekingString)
			if err != nil || res <= 0 {
				fmt.Println("ERROR. Not valid number of ants")
				errorStatus = true
				return errorStatus
			}
			st.AntCount = res
			firstLine = false

		} else { // All after the first line. Seaching for rooms and paths

			if startFlag {
				roomSlice := strings.Split(chekingString, " ")
				if len(roomSlice) != 3 {
					fmt.Println("ERROR. Not right format of start room")
					errorStatus = true
					return errorStatus
				}
				st.Rooms[roomSlice[0]] = &st.Room{
					Name:     roomSlice[0],
					X:        roomSlice[1],
					Y:        roomSlice[2],
					AntCount: st.AntCount,
				}
				st.StartRoom = roomSlice[0]
				startFlag = false
				countEndStartRooms++
			}

			if endFlag {
				roomSlice := strings.Split(chekingString, " ")
				if len(roomSlice) != 3 {
					fmt.Println("ERROR. Not right format of end room")
					errorStatus = true
					return errorStatus
				}
				st.Rooms[roomSlice[0]] = &st.Room{
					Name:     roomSlice[0],
					X:        roomSlice[1],
					Y:        roomSlice[2],
					AntCount: 0,
				}
				st.EndRoom = roomSlice[0]
				endFlag = false
				countEndStartRooms++
			}

			switch {
			// Next line will be start room
			case chekingString == "##start":
				startFlag = true

			// Next line will be end room
			case chekingString == "##end":
				endFlag = true

			// Right fromat of the room is "[num] [X] [Y]"
			case strings.Contains(chekingString, " "):
				roomSlice := strings.Split(chekingString, " ")
				if len(roomSlice) != 3 {
					fmt.Println("ERROR. Not right format of connections")
					errorStatus = true
					return errorStatus
				}

				st.Rooms[roomSlice[0]] = &st.Room{
					Name:     roomSlice[0],
					X:        roomSlice[1],
					Y:        roomSlice[2],
					AntCount: 0,
				}

			// Right fromat of the link between rooms is "[num]-[num]"
			case strings.Contains(chekingString, "-"):
				conectionSlice := strings.Split(chekingString, "-")
				if len(conectionSlice) != 2 {
					fmt.Println("ERROR. Not right format of connections")
					errorStatus = true
					return errorStatus
				}

				if conectionSlice[0] == st.EndRoom || conectionSlice[1] == st.EndRoom {
					startConnections++
				}
				if conectionSlice[0] == st.StartRoom || conectionSlice[1] == st.StartRoom {
					endConnections++
				}

				st.Connections = append(st.Connections,
					st.Connection{
						From: conectionSlice[0],
						To:   conectionSlice[1],
					})
			}
		}
	}
	// Scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if countEndStartRooms != 2 {
		fmt.Println("ERROR. Start or end rooms were not found")
		errorStatus = true
		return errorStatus
	}
	if endConnections == 0 || startConnections == 0 {
		fmt.Println("ERROR. Start or end rooms were not found in connections")
		errorStatus = true
		return errorStatus
	}

	return errorStatus
}
