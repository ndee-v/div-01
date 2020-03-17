package tools

//Data struct using for all parametres from
type Data struct {
	Ants  int
	Graph []Vertex
}

// Vertex struct
type Vertex struct {
	Name    string
	X, Y    string
	Type    string
	Visited bool
	Track   []*Vertex
	Next    []*Vertex
	Depth   int
	// Prev    []*Vertex
	Busy bool
}

// Ways is all ways from Start to End
type Ways struct {
	Ways [][]string
}

// Stop - is struct with params for program
type Stop struct {
	Stop       bool
	Steps      int
	Ants       int
	MaxWays    int
	Start      *Vertex
	End        *Vertex
	Rep        bool
	Total      bool
	Clear      bool
	NameOfFile string
}

//Ant struct using for visualise
type Ant struct {
	Name     string
	Where    string
	Position int
	Step     int
	Way      int
}
