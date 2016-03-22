package events

var eventIdChan chan int

func init() {
	eventIdChan = make(chan int)
	go generateEventIds()
}

func generateEventIds() {
	id := 0
	for {
		id++
		eventIdChan <- id
	}
}

func nextEventId() int {
	return <-eventIdChan
}
