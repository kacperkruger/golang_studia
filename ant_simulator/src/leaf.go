package ant_simulator_src

type Leaf struct {
	x int
	y int
}

func (l *Leaf) Move(dx int, dy int) {
	l.x += dx
	l.y += dy
}

func (l *Leaf) GetX() int {
	return l.x
}

func (l *Leaf) GetY() int {
	return l.y
}
