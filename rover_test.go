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
	rover.New(4, 7, 'N')
}

func TestRoverInstructTurn(t *testing.T) {
	testCases := map[string]struct {
		turn, start, finish rune
	}{
		"Turn left from north":  {turn: 'L', start: 'N', finish: 'W'},
		"Turn left from east":   {turn: 'L', start: 'E', finish: 'N'},
		"Turn left from south":  {turn: 'L', start: 'S', finish: 'E'},
		"Turn left from west":   {turn: 'L', start: 'W', finish: 'S'},
		"Turn right from north": {turn: 'R', start: 'N', finish: 'E'},
		"Turn right from east":  {turn: 'R', start: 'E', finish: 'S'},
		"Turn right from south": {turn: 'R', start: 'S', finish: 'W'},
		"Turn right from west":  {turn: 'R', start: 'W', finish: 'N'},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			// Given / Arrange
			is := is.New(t)
			rov := rover.New(4, 7, tc.start)

			// When	/ Act
			rov.Instruct(tc.turn)

			// Then / Assert
			is.Equal(rov.GetFacing(), tc.finish) // not facing correct final direction
		})
	}
}

func TestRoverMove(t *testing.T) {
	testCases := map[rune][]struct {
		finish  rover.Coords
		moveDir rune
	}{
		'N': {
			{moveDir: 'F', finish: rover.Coords{0, 1}},
			{moveDir: 'B', finish: rover.Coords{0, -1}},
		},

		'E': {
			{moveDir: 'F', finish: rover.Coords{1, 0}},
			{moveDir: 'B', finish: rover.Coords{-1, 0}},
		},

		'S': {
			{moveDir: 'F', finish: rover.Coords{0, -1}},
			{moveDir: 'B', finish: rover.Coords{0, 1}},
		},

		'W': {
			{moveDir: 'F', finish: rover.Coords{-1, 0}},
			{moveDir: 'B', finish: rover.Coords{1, 0}},
		},
	}

	for facing, outerCase := range testCases {
		for _, innerCase := range outerCase {
			desc := fmt.Sprintf("Face %v and move %v",
				string(facing), string(innerCase.moveDir))

			t.Run(desc, func(t *testing.T) {
				// Arrange
				is := is.New(t)
				rov := rover.New(0, 0, facing)

				// Act
				rov.Instruct(innerCase.moveDir)

				// Assert
				is.Equal(rov.GetCoords(), innerCase.finish) // hasn't moved correctly
			})
		}
	}
}

func TestRoverMultipleInstructions(t *testing.T) {
	// Arrange
	is := is.New(t)
	rov := rover.New(0, 0, 'N')

	// Act
	rov.Instruct('F', 'R', 'F', 'F')

	// Assert
	is.Equal(rov.GetCoords(), rover.Coords{2, 1})
}
