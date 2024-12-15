package c2d

type C struct {
	X, Y int
}

func (c *C) Add(other C) {
	c.X += other.X
	c.Y += other.Y
}

func (c *C) Subtract(other C) {
	c.X -= other.X
	c.Y -= other.Y
}

func (c C) Plus(other C) C {
	return C{c.X + other.X, c.Y + other.Y}
}

func (c C) Minus(other C) C {
	return C{c.X - other.X, c.Y - other.Y}
}

func (c C) Neighbours(lenX, lenY int, includeDiagonal bool) []C {
	var ns []C
	if includeDiagonal {
		ns = make([]C, 8)
	} else {
		ns = make([]C, 4)
	}
	if c.X > 0 {
		ns[0] = C{c.X - 1, c.Y}
	}
	if c.X < lenX-1 {
		ns[1] = C{c.X + 1, c.Y}
	}
	if c.Y > 0 {
		ns[2] = C{c.X, c.Y - 1}
	}
	if c.Y < lenY-1 {
		ns[3] = C{c.X, c.Y + 1}
	}
	if includeDiagonal {
		if c.X > 0 && c.Y > 0 {
			ns[4] = C{c.X - 1, c.Y - 1}
		}
		if c.X < lenX-1 && c.Y > 0 {
			ns[5] = C{c.X + 1, c.Y - 1}
		}
		if c.X < lenX-1 && c.Y < lenY-1 {
			ns[6] = C{c.X + 1, c.Y + 1}
		}
		if c.X > 0 && c.Y < lenY-1 {
			ns[7] = C{c.X - 1, c.Y + 1}
		}
	}
	return ns
}

func (c C) AllNeighbours(includeDiagonal bool) []C {
	var ns []C
	if includeDiagonal {
		ns = make([]C, 8)
	} else {
		ns = make([]C, 4)
	}
	ns[0] = C{c.X - 1, c.Y}
	ns[1] = C{c.X + 1, c.Y}
	ns[2] = C{c.X, c.Y - 1}
	ns[3] = C{c.X, c.Y + 1}
	if includeDiagonal {
		ns[4] = C{c.X - 1, c.Y - 1}
		ns[5] = C{c.X + 1, c.Y - 1}
		ns[6] = C{c.X + 1, c.Y + 1}
		ns[7] = C{c.X - 1, c.Y + 1}
	}
	return ns
}

func (c C) RotateRight() C {
	return C{c.Y, -c.X}
}

func (c C) RotateLeft() C {
	return C{-c.Y, c.X}
}

// Rotate45Right rotates a vector like a square around origo
// Example C{1, 1}.Rotate45Right() // {C{1, 0}}
func (c C) Rotate45Right() C {
	if c.X == 0 {
		return C{c.Y, c.Y}
	}
	if c.Y == 0 {
		return C{c.X, -c.X}
	}
	if c.X == c.Y {
		return C{c.X, 0}
	} else {
		return C{0, c.Y}
	}
}

// Rotate45Left rotates a vector like a square around origo
// Example C{1, 1}.Rotate45Left() // {C{0, 1}}
func (c C) Rotate45Left() C {
	if c.X == 0 {
		return C{-c.Y, c.Y}
	}
	if c.Y == 0 {
		return C{c.X, c.X}
	}
	if c.X == c.Y {
		return C{0, c.Y}
	} else {
		return C{c.X, 0}
	}
}

func (c C) ManhattanDistance(other C) int {
	return absDiffInt(c.X, other.X) + absDiffInt(c.Y, other.Y)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
