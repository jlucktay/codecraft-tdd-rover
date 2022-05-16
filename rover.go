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
		case 'L':
			r.turnLeft()
		case 'R':
			r.turnRight()
		case 'F':
			r.moveForward()
		case 'B':
			r.moveBackward()
		}
	}
}

func (r *Rover) turnLeft() {
	switch r.facing {
	case 'N':
		r.facing = 'W'
	case 'E':
		r.facing = 'N'
	case 'S':
		r.facing = 'E'
	case 'W':
		r.facing = 'S'
	}
}

func (r *Rover) turnRight() {
	switch r.facing {
	case 'N':
		r.facing = 'E'
	case 'E':
		r.facing = 'S'
	case 'S':
		r.facing = 'W'
	case 'W':
		r.facing = 'N'
	}
}

func (r *Rover) moveForward() {
	switch r.facing {
	case 'N':
		r.c.Y++
	case 'E':
		r.c.X++
	case 'S':
		r.c.Y--
	case 'W':
		r.c.X--
	}
}

func (r *Rover) moveBackward() {
	switch r.facing {
	case 'N':
		r.c.Y--
	case 'E':
		r.c.X--
	case 'S':
		r.c.Y++
	case 'W':
		r.c.X++
	}
}
