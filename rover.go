package rover

import "fmt"

// Direction that the rover is facing
type Direction string

// Command the rover to do stuffs
type Command string

const (
	// NORTH facing
	NORTH = Direction("N")
	// SOUTH facing
	SOUTH = Direction("S")
	// WEST facing
	WEST = Direction("W")
	// EAST facing
	EAST = Direction("E")

	// FORWARD command to go forwards
	FORWARD = Command("f")
	// BACKWARD command to go backwards
	BACKWARD = Command("b")
	// LEFT command to go left
	LEFT = Command("l")
	// RIGHT command to go right
	RIGHT = Command("r")

	// LENGTH of the world
	LENGTH = 10
	// WIDTH of the world
	WIDTH = 10
)

// Rover knows where it is and where it faces
type Rover struct {
	X int
	Y int
	D Direction
}

type cmdKey struct {
	D Direction
	C Command
}

type cmdResult struct {
	X int
	Y int
}

var changes = map[cmdKey]cmdResult{

	cmdKey{NORTH, FORWARD}:  cmdResult{0, 1},
	cmdKey{NORTH, BACKWARD}: cmdResult{0, -1},

	cmdKey{SOUTH, FORWARD}:  cmdResult{0, -1},
	cmdKey{SOUTH, BACKWARD}: cmdResult{0, 1},
}

// New creates a new rover
func New(d Direction, x, y int) *Rover {

	return &Rover{
		X: x,
		Y: y,
		D: d,
	}
}

// RunCommands processes the list of commands updating the position of the rover
func (r *Rover) RunCommands(cmds []Command) {
	for _, cmd := range cmds {
		// given the current direction and command lookup the change
		c := changes[cmdKey{D: r.D, C: cmd}]
		fmt.Printf("%v : %v -> %+v\n", r.D, cmd, c)
		r.X += c.X
		r.Y += c.Y
	}
}
