package structs

var (
	// Global variables that can be accessed from other files
	StartRoom   string
	EndRoom     string
	AntCount    int
	Rooms       = make(map[string]*Room)
	Connections []Connection
	File        string
)

type Routes struct {
	Key   int
	Value []string
}

type Room struct {
	Name     string
	X, Y     string
	AntCount int
}

type Connection struct {
	From string
	To   string
}

type Ant struct {
	Path     []string
	RoomID   int
	Previous string
	Ignore   bool
}
