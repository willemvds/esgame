package events

type MouseMoveEvent struct {
	id   int
	x, y int
}

func NewMouseMove(x, y int) MouseMoveEvent {
	id := nextEventId()
	return MouseMoveEvent{id, x, y}
}
