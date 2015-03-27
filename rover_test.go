package rover

import "testing"

func TestRoverCreate(t *testing.T) {

	rover := New(NORTH, 1, 1)

	if rover.X != 1 {
		t.Error("X != 1")
	}

	if rover.Y != 1 {
		t.Error("Y != 1")
	}

	if rover.D != NORTH {
		t.Error("Should be facing NORTH")
	}

}

func TestRoverMovesForwardBackwardsNorth(t *testing.T) {

	rover := New(NORTH, 1, 1)

	rover.RunCommands([]Command{FORWARD})
	if rover.Y != 2 {
		t.Error("X != 2, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	rover.RunCommands([]Command{FORWARD, FORWARD})
	if rover.Y != 4 {
		t.Error("X != 4, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	rover.RunCommands([]Command{BACKWARD, BACKWARD})
	if rover.Y != 2 {
		t.Error("X != 2, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

}

func TestRoverMovesForwardBackwardsSouth(t *testing.T) {

	rover := New(SOUTH, 1, 5)

	rover.RunCommands([]Command{FORWARD})
	if rover.Y != 4 {
		t.Error("X != 4, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	rover.RunCommands([]Command{FORWARD, FORWARD})
	if rover.Y != 2 {
		t.Error("X != 2, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	rover.RunCommands([]Command{BACKWARD, BACKWARD})
	if rover.Y != 4 {
		t.Error("X != 4, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

}

func TestRoverLeft(t *testing.T) {

	rover := New(SOUTH, 1, 5)

	rover.RunCommands([]Command{LEFT})
	if rover.D != EAST {
		t.Error("Not facing EAST")
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

}

func TestRoverRight(t *testing.T) {

	rover := New(WEST, 1, 5)

	rover.RunCommands([]Command{RIGHT})
	if rover.D != NORTH {
		t.Error("Not facing NORTH")
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	rover.RunCommands([]Command{RIGHT, RIGHT, RIGHT})
	if rover.D != WEST {
		t.Error("Should have turned around")
	}

}

func TestRoverWrapping(t *testing.T) {

	rover := New(WEST, 0, 0)

	rover.RunCommands([]Command{FORWARD})
	if rover.X != SIZE {
		t.Errorf("Rover fell off the planet -> %+v", rover)
	}

}
