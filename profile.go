package goutil

import (
	"fmt"
	"log"
	"time"
)

//Simple timer struct
type Timer struct {
	start_time time.Time
}

// Creates a new timer instance
func NewTimer() *Timer {
	return &Timer{start_time: time.Now()}
}

//Measures and logs time elapsed since the creation of this timer instance
func (t Timer) Elapsed(action string) {
	log.Print(fmt.Sprintf("%s took %s", action, time.Since(t.start_time)))
}
