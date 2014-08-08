package goutil

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	go func() {
		ch1 <- "1"
		ch2 <- "2"
		ch1 <- "3"
		ch2 <- "4"
	}()
	ch_list := []chan interface{}{ch1, ch2}
	out := MergeChannels(ch_list)
	out1 := <-out
	if out1 != "1" {
		t.Error("First message from channel should be 1, got", out1)
	}
	out2 := <-out
	if out2 != "2" {
		t.Error("Second message from channel should be 2, got", out2)
	}
	out3 := <-out
	if out3 != "4" {
		t.Error("Third message from channel should be 3, got", out3)
	}
	out4 := <-out
	if out4 != "3" {
		t.Error("Last message from channel should be 4, got", out4)
	}
}

func TestReadLines(t *testing.T) {
	lines := []string{}
	for line := range ReadLines("test_lines.txt") {
		lines = append(lines, line)
	}
	if lines[0] != "First Line" {
		t.Error("First Line was expected, got", lines[0])
	}
	if lines[1] != "Second Line" {
		t.Error("Second Line was expected got", lines[1])
	}
	if len(lines) != 2 {
		t.Error("There should only be 2 lines, got", len(lines))
	}
}

func TestSliceChannel(t *testing.T) {
	fibGenerator := func() (ch chan interface{}) {
		ch = make(chan interface{})
		go func() {
			last := 0
			current := 1
			for {
				last += current
				last, current = current, last
				ch <- current
			}
		}()
		return
	}
	fibs := []int{}
	for i := range SliceChannel(fibGenerator(), 4) {
		fibs = append(fibs, i.(int))
	}
	if len(fibs) != 4 {
		t.Error("Slice len should be 4")
	}
	if fibs[0] != 1 {
		t.Error("First element should be 1, got", fibs[0])
	}
	if fibs[1] != 2 {
		t.Error("Second element should be 2 got", fibs[1])
	}
	if fibs[2] != 3 {
		t.Error("Last element should be 3 got", fibs[2])
	}
	if fibs[3] != 5 {
		t.Error("Last element should be 5 got", fibs[3])
	}
}

func TestTimer(t *testing.T) {
	timer := NewTimer()
	defer timer.Elapsed("Testing")
}

func TestSet(t *testing.T) {
	strings_a := []interface{}{"a", "b", "c", "d"}
	strings_b := []interface{}{"b", "c", "d", "f", "g"}
	set_of_strings_a := NewSet(strings_a)
	set_of_strings_b := NewSet(strings_b)
	set_of_strings_c := set_of_strings_a.Intersection(set_of_strings_b)
	if !set_of_strings_c.Equals(NewSet([]interface{}{"b", "c", "d"})) {
		t.Error("Intersection should be 'c', 'd'", set_of_strings_c)
	}
	set_of_strings_d := set_of_strings_a.Union(set_of_strings_b)
	if !set_of_strings_d.Equals(NewSet([]interface{}{"a", "b", "c", "d", "f", "g"})) {
		t.Error("Union should be 'a' 'b' c' 'd' 'f' 'g'", set_of_strings_d)
	}
	set_of_strings_e := set_of_strings_a.Difference(set_of_strings_b)
	if !set_of_strings_e.Equals(NewSet([]interface{}{"a"})) {
		t.Error("Difference should be 'a' 'b'", set_of_strings_d)
	}
	popped := set_of_strings_e.Pop()
	if popped != "a" {
		t.Error("Popped should be 'a'", popped)
	}
	if !set_of_strings_e.IsEmpty() {
		t.Error("Set should be empty", set_of_strings_e)
	}
}
