package events

import (
	"time"
)

type TickEvent struct {
	when time.Time
}

func NewTick(when time.Time) TickEvent {
	return TickEvent{when}
}

func (te TickEvent) When() time.Time {
	return te.when
}
