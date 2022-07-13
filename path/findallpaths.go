package path

import (
	st "lem-in/structs"
	"sort"
	"strconv"
)

type Graph struct {
	adjacency map[string][]string
}

var (
	counter int
	route   []string
	routes  = make(map[int][]string)
)

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

func FuncAddVertexAndEdge(g Graph) bool {
	// Adds Vertex
	for _, i := range st.Rooms {
		if g.AddVertex(i.Name) {
			return true
		} else {
			g.AddVertex(i.Name)
		}
	}
	// Adds Edge
	for _, i := range st.Connections {
		if g.AddEdge(i.From, i.To) {
			return true
		} else {
			g.AddEdge(i.From, i.To)
		}
	}

	return false
}

func (g *Graph) AddVertex(vertex string) bool {

	if _, ok := g.adjacency[vertex]; ok {
		return true
	}

	g.adjacency[vertex] = []string{}
	return false
}

func (g *Graph) AddEdge(vertex, node string) bool {

	if _, ok := g.adjacency[vertex]; !ok {
		return true
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		return true
	}

	g.adjacency[node] = append(g.adjacency[node], vertex)
	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return false
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (g Graph) CreateVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}

func FindAllPaths(starting string, ending string, g Graph, visited map[string]bool) map[int][]string {
	if starting == ending {
		counter++
		route = append(route, starting)
		routes[counter] = append(routes[counter], route...)
		route = route[:len(route)-1]
	} else {
		for _, node := range g.adjacency[starting] {
			if !visited[node] {
				route = append(route, starting)
				visited[starting] = true
				FindAllPaths(node, ending, g, visited)
				if len(route) != 0 {
					route = route[:len(route)-1]
				}
			}
		}
	}
	visited[starting] = false
	return routes
}

func SortPaths(allRoutes map[int][]string) ([][]string, []st.Routes) {
	var sortedRoutes []st.Routes
	for k, v := range allRoutes { // Add all paths to struct slice
		sortedRoutes = append(sortedRoutes, st.Routes{Key: k, Value: v})
	}

	sort.Slice(sortedRoutes, func(i, j int) bool { // Sort all paths by lenght
		return len(sortedRoutes[i].Value) < len(sortedRoutes[j].Value)
	})

	var withoutstartend []st.Routes
	for _, val := range sortedRoutes { // Delete starting and ending rooms
		val.Value = val.Value[1 : len(val.Value)-1]
		withoutstartend = append(withoutstartend, val)
	}

	var routes []string
	var pathsComboKeys [][]string
	var keys []string

	// This part creates a slice of slices with keys of non overlaping paths
	for k, val := range withoutstartend {
		routes = nil
		keys = nil
		routes = append(routes, val.Value...)
		keys = append(keys, strconv.Itoa(val.Key))
		for _, va2 := range withoutstartend {
			if val.Key == va2.Key {
				continue
			}
			if CompareTwoRoutes(routes, va2.Value) {
				routes = append(routes, va2.Value...)
				keys = append(keys, strconv.Itoa(va2.Key))
			}
		}
		if len(val.Value) == 0 {
			routes = []string{strconv.Itoa(k + 1)}
			pathsComboKeys = append(pathsComboKeys, routes)
		}
		if len(val.Value) == 1 {
			routes = []string{strconv.Itoa(k + 1)}
			pathsComboKeys = append(pathsComboKeys, routes)
		}
		pathsComboKeys = append(pathsComboKeys, keys)
	}

	// retrun combo of keys and all paths
	return pathsComboKeys, withoutstartend

}

func BestPath(routeComboKeys [][]string, routes []st.Routes) [][]string {
	var temp []int // Save the lenght of path
	var res []int
	for _, val := range routeComboKeys {
		for i := range val { // Finding lenght of all paths in combination of keys
			for _, goodRoute := range routes {

				k, _ := strconv.Atoi(val[i]) // k - key
				if goodRoute.Key == k {
					temp = append(temp, len(goodRoute.Value))
					break
				}
			}
		}
		for i := st.AntCount; i > 0; i-- { // Adding ants to routes
			_, index := LongestOrShortestRoute(temp)
			temp[index]++
		}

		l, _ := LongestOrShortestRoute(temp) // Taking the longest, because we can not use the smallest amount
		res = append(res, l)
		temp = nil

	}
	_, id := LongestOrShortestRoute(res)

	var allPaths [][]string
	for _, v := range routeComboKeys[id] {
		for _, v2 := range routes {
			check, _ := strconv.Atoi(v)
			if v2.Key == check {
				v2.Value = append(v2.Value, st.EndRoom)
				allPaths = append(allPaths, v2.Value)
			}
		}

	}
	return allPaths
}

func CompareTwoRoutes(a []string, b []string) bool {
	for _, val := range a {
		for _, val2 := range b {
			if val == val2 {
				return false
			}
		}
	}
	return true
}

func LongestOrShortestRoute(a []int) (temp int, temp2 int) {
	var temp1 int
	for i, val := range a {
		if i == 0 {
			temp = val
			temp1 = val
		} else {
			if val > temp {
				temp = val
			}
			if val < temp1 {
				temp1 = val
				temp2 = i
			}
		}
	}
	return temp, temp2
}
