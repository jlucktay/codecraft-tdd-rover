package rover_test

import (
	"fmt"
	"testing"

	"github.com/matryer/is"

	rover "go.jlucktay.dev/tdd-rover"
)

// - Rover initialise at coords with facing
//   - positioning grid would be pair of signed integers
func TestRoverInit(t *testing.T) {
	rover.New(4, 7, rover.North)
}

func TestRoverInstructTurn(t *testing.T) {
	testCases := map[string]struct {
		turn, start, finish rune
	}{
		"Turn left from north":  {turn: rover.Left, start: rover.North, finish: rover.West},
		"Turn left from east":   {turn: rover.Left, start: rover.East, finish: rover.North},
		"Turn left from south":  {turn: rover.Left, start: rover.South, finish: rover.East},
		"Turn left from west":   {turn: rover.Left, start: rover.West, finish: rover.South},
		"Turn right from north": {turn: rover.Right, start: rover.North, finish: rover.East},
		"Turn right from east":  {turn: rover.Right, start: rover.East, finish: rover.South},
		"Turn right from south": {turn: rover.Right, start: rover.South, finish: rover.West},
		"Turn right from west":  {turn: rover.Right, start: rover.West, finish: rover.North},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			// Given / Arrange
			is := is.New(t)
			rov := rover.New(4, 7, tc.start)

			// When	/ Act
			rov.Instruct(tc.turn)

			// Then / Assert
			is.Equal(rov.GetFacing(), tc.finish) // Rover is not facing correct final direction
		})
	}
}

func TestRoverMove(t *testing.T) {
	testCases := map[rune][]struct {
		moveDir rune
		finish  rover.Coords
	}{
		rover.North: {
			{moveDir: rover.Forward, finish: rover.Coords{0, 1}},
			{moveDir: rover.Backward, finish: rover.Coords{0, -1}},
		},

		rover.East: {
			{moveDir: rover.Forward, finish: rover.Coords{1, 0}},
			{moveDir: rover.Backward, finish: rover.Coords{-1, 0}},
		},

		rover.South: {
			{moveDir: rover.Forward, finish: rover.Coords{0, -1}},
			{moveDir: rover.Backward, finish: rover.Coords{0, 1}},
		},

		rover.West: {
			{moveDir: rover.Forward, finish: rover.Coords{-1, 0}},
			{moveDir: rover.Backward, finish: rover.Coords{1, 0}},
		},
	}

	for facing, outerCase := range testCases {
		for _, innerCase := range outerCase {
			desc := fmt.Sprintf("Face %v and move %v", string(facing), string(innerCase.moveDir))

			t.Run(desc, func(t *testing.T) {
				// Arrange
				is := is.New(t)
				rov := rover.New(0, 0, facing)

				// Act
				rov.Instruct(innerCase.moveDir)

				// Assert
				is.Equal(rov.GetCoords(), innerCase.finish) // Rover hasn't moved correctly
			})
		}
	}
}

func TestRoverMultipleInstructions(t *testing.T) {
	// Arrange
	is := is.New(t)
	rov := rover.New(0, 0, rover.North)

	// Act
	rov.Instruct(rover.Forward, rover.Right, rover.Forward, rover.Forward)

	// Assert
	is.Equal(rov.GetCoords(), rover.Coords{2, 1}) // Rover hasn't followed multiple instructions properly
}
