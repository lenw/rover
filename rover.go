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

	// SIZE of the world
	SIZE = 10
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
	D Direction
	X int
	Y int
}

var changes = map[cmdKey]cmdResult{

	cmdKey{NORTH, FORWARD}:  cmdResult{NORTH, 0, 1},
	cmdKey{NORTH, BACKWARD}: cmdResult{NORTH, 0, -1},
	cmdKey{NORTH, LEFT}:     cmdResult{WEST, 0, 0},
	cmdKey{NORTH, RIGHT}:    cmdResult{EAST, 0, 0},

	cmdKey{SOUTH, FORWARD}:  cmdResult{SOUTH, 0, -1},
	cmdKey{SOUTH, BACKWARD}: cmdResult{SOUTH, 0, 1},
	cmdKey{SOUTH, LEFT}:     cmdResult{EAST, 0, 0},
	cmdKey{SOUTH, RIGHT}:    cmdResult{WEST, 0, 0},

	cmdKey{WEST, FORWARD}:  cmdResult{WEST, -1, 0},
	cmdKey{WEST, BACKWARD}: cmdResult{WEST, 1, 0},
	cmdKey{WEST, LEFT}:     cmdResult{SOUTH, 0, 0},
	cmdKey{WEST, RIGHT}:    cmdResult{NORTH, 0, 0},

	cmdKey{EAST, FORWARD}:  cmdResult{EAST, 1, 0},
	cmdKey{EAST, BACKWARD}: cmdResult{EAST, -1, 0},
	cmdKey{EAST, LEFT}:     cmdResult{NORTH, 0, 0},
	cmdKey{EAST, RIGHT}:    cmdResult{SOUTH, 0, 0},
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
		r.D = c.D
		r.X += c.X
		r.Y += c.Y

		if r.X < 0 {
			r.X = SIZE
		}
		if r.Y < 0 {
			r.Y = SIZE
		}
		if r.X > SIZE {
			r.X = 0
		}
		if r.Y > SIZE {
			r.Y = 0
		}

	}
}
