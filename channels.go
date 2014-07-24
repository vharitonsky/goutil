package goutil

import (
	"sync"
)

func MergeChannels(cs []chan interface{}) chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	output := func(c chan interface{}) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func SliceChannel(ch chan interface{}, sliceLen uint) (out chan interface{}) {
	out = make(chan interface{})
	go func() {
		for i := 0; i < int(sliceLen); i++ {
			temp := <-ch
			out <- temp
		}
		close(out)
	}()
	return
}
