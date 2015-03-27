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

	cmds := []Command{FORWARD}
	rover.RunCommands(cmds)
	if rover.Y != 2 {
		t.Error("X != 2, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	cmds = []Command{FORWARD, FORWARD}
	rover.RunCommands(cmds)
	if rover.Y != 4 {
		t.Error("X != 4, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	cmds = []Command{BACKWARD, BACKWARD}
	rover.RunCommands(cmds)
	if rover.Y != 2 {
		t.Error("X != 2, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

}

func TestRoverMovesForwardBackwardsSouth(t *testing.T) {

	rover := New(SOUTH, 1, 5)

	cmds := []Command{FORWARD}
	rover.RunCommands(cmds)
	if rover.Y != 4 {
		t.Error("X != 4, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	cmds = []Command{FORWARD, FORWARD}
	rover.RunCommands(cmds)
	if rover.Y != 2 {
		t.Error("X != 2, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

	cmds = []Command{BACKWARD, BACKWARD}
	rover.RunCommands(cmds)
	if rover.Y != 4 {
		t.Error("X != 4, got ", rover.Y)
	}
	if rover.X != 1 {
		t.Error("X should not change.")
	}

}
