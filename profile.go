package goutil

import (
	"fmt"
	"log"
	"time"
)

type Timer struct {
	start_time time.Time
}

func NewTimer() *Timer {
	return &Timer{start_time: time.Now()}
}

func (t Timer) Elapsed(action string) {
	log.Print(fmt.Sprintf("%s took %s", action, time.Since(t.start_time)))
}
