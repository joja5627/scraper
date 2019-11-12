package scrape

import "time"

//Retry
type Retry struct {
	Pause  time.Duration
	Count     int
}

