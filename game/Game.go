package game

import (
	"fmt"
	"time"
)

type Game struct {
	ticks int
}

func New() Game {
	return Game{}
}

func (g *Game) Tick(when time.Time) {
	g.ticks++
	fmt.Printf("[game] tick %d (%v)\n", g.ticks, time.Since(when))
}

func (g *Game) MakeMoveX() error {
	if g.ticks < 10 {
		return fmt.Errorf("Invalid move. %d ticks. Can only move after 10", g.ticks)
	}

	fmt.Println("[game] Move X has been made")
	return nil
}
