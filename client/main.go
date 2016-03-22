package main

import (
	"fmt"
	"time"

	"github.com/willemvds/esgame/events"
	"github.com/willemvds/esgame/game"
)

func main() {
	var err error
	game := game.New()

	msmv := events.NewMouseMove(10, 20)
	fmt.Printf("%v\n", msmv)

	tickChan := time.Tick(50 * time.Millisecond)
	for now := range tickChan {
		tick := events.NewTick(now)
		game.Tick(tick.When())
		if err = game.MakeMoveX(); err != nil {
			fmt.Println(err)
		}
	}
}
