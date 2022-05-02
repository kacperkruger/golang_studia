package ant_simulator_src

type Ant struct {
	x    int
	y    int
	leaf *Leaf
}

func (a *Ant) Move(dx int, dy int) {
	a.x += dx
	a.y += dy
	if IsCarryingLeaf(*a) {
		a.leaf.Move(dx, dy)
	}
}

func (a *Ant) GetX() int {
	return a.x
}

func (a *Ant) GetY() int {
	return a.y
}

func (a *Ant) PickUpLeaf(leaf *Leaf) {
	a.leaf = leaf
}

func (a *Ant) DropLeaf() {
	a.leaf = &Leaf{0, 0}
}

func IsCarryingLeaf(a Ant) bool {
	if a.leaf.x != 0 && a.leaf.y != 0 {
		return true
	}
	return false
}
