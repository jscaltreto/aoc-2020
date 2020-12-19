package main

type Ship struct {
	x   int
	y   int
	dir int
}

func (s *Ship) turnRight(deg int) {
	pos := deg / 90
	s.dir = (s.dir + pos) % len(DIRECTIONS)
}
func (s *Ship) turnLeft(deg int) {
	pos := (-deg / 90) + 4
	s.dir = (s.dir + pos) % len(DIRECTIONS)
}
func (s *Ship) move(dir byte, dist int) {
	if dir == FORWARD {
		dir = DIRECTIONS[s.dir]
	}
	switch dir {
	case RIGHT:
		s.turnRight(dist)
	case LEFT:
		s.turnLeft(dist)
	case NORTH:
		s.y -= dist
	case SOUTH:
		s.y += dist
	case EAST:
		s.x += dist
	case WEST:
		s.x -= dist
	}
}
