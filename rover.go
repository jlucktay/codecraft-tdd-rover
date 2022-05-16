package rover

// - Rover initialise at coords with facing
//   - positioning grid would be pair of signed integers
type Rover struct {
	c      Coords
	facing rune
}

type Coords struct {
	X, Y int
}

const (
	Left  rune = 'L'
	Right rune = 'R'

	Forward  rune = 'F'
	Backward rune = 'B'

	North rune = 'N'
	East  rune = 'E'
	South rune = 'S'
	West  rune = 'W'
)

func New(x, y int, face rune) Rover {
	return Rover{
		c:      Coords{X: x, Y: y},
		facing: face,
	}
}

func (r *Rover) GetFacing() rune {
	return r.facing
}

func (r *Rover) GetCoords() Coords {
	return r.c
}

func (r *Rover) Instruct(instructions ...rune) {
	for _, instruction := range instructions {
		switch instruction {
		case Left:
			r.turnLeft()
		case Right:
			r.turnRight()
		case Forward:
			r.moveForward()
		case Backward:
			r.moveBackward()
		}
	}
}

func (r *Rover) turnLeft() {
	switch r.facing {
	case North:
		r.facing = West
	case East:
		r.facing = North
	case South:
		r.facing = East
	case West:
		r.facing = South
	}
}

func (r *Rover) turnRight() {
	switch r.facing {
	case North:
		r.facing = East
	case East:
		r.facing = South
	case South:
		r.facing = West
	case West:
		r.facing = North
	}
}

func (r *Rover) moveForward() {
	switch r.facing {
	case North:
		r.c.Y++
	case East:
		r.c.X++
	case South:
		r.c.Y--
	case West:
		r.c.X--
	}
}

func (r *Rover) moveBackward() {
	switch r.facing {
	case North:
		r.c.Y--
	case East:
		r.c.X--
	case South:
		r.c.Y++
	case West:
		r.c.X++
	}
}
