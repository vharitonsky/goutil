package goutil

import (
	"fmt"
	"log"
	"runtime"
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
func (t *Timer) Elapsed(action string) {
	log.Print(fmt.Sprintf("%s took %s", action, time.Since(t.start_time)))
}

type MemProfiler struct {
	start_mem uint64
}

func NewMemProfiler() *MemProfiler {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)
	return &MemProfiler{start_mem: mem.Alloc}
}

func (m *MemProfiler) Used(action string) {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)
	log.Print(fmt.Sprintf("%s allocated %f Kb", action, float64(mem.Alloc-m.start_mem)/1024.0))
}
