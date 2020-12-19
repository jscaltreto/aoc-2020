package main

type Waypoint struct {
	x int
	y int
}

func (w *Waypoint) rotate(pos int) {
	x := w.x
	y := w.y
	switch pos % 4 {
	case 1:
		w.x = y
		w.y = -x
	case 2:
		w.x = -x
		w.y = -y
	case 3:
		w.x = -y
		w.y = x
	}
}
func (w *Waypoint) turnRight(deg int) {
	pos := deg / 90
	w.rotate(pos)
}
func (w *Waypoint) turnLeft(deg int) {
	pos := (-deg / 90) + 4
	w.rotate(pos)
}
func (w *Waypoint) move(dir byte, dist int) {
	switch dir {
	case RIGHT:
		w.turnRight(dist)
	case LEFT:
		w.turnLeft(dist)
	case NORTH:
		w.y -= dist
	case SOUTH:
		w.y += dist
	case EAST:
		w.x += dist
	case WEST:
		w.x -= dist
	}
}
