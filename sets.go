package goutil

type Set struct {
	contents map[interface{}]bool
}

func NewSet(elements []interface{}) *Set {
	set = Set{make(map[interface{}]bool, len(elements))}
	for el := range elements {
		set.contents[el] = true
	}
	return &set
}

func (s_left *Set) Intersection(s_right *Set) *Set {
	set = Set{make(map[interface{}]bool, len(s_left.contents)+len(s_right.contents))}
	for el, _ := range s_left.contents {
		_, found = s_right[el]
		if found {
			set.contents[el] = true
		}
	}
	return &set
}

func (s_left *Set) Union(s_right *Set) *Set {
	set = Set{make(map[interface{}]bool, len(s_left.contents)+len(s_right.contents))}
	for el, _ := range s_left.contents {
		set.contents[el] = true
	}
	for el, _ := range s_right.contents {
		set.contents[el] = true
	}
	return &set
}

func (s_left *Set) Difference(s_right *Set) *Set {
	set = Set{make(map[interface{}]bool, len(s_left.contents)+len(s_right.contents))}
	for el, _ := range s_left.contents {
		_, found = s_right.contents[el]
		if !found {
			set.contents[el] = true
		}
	}
	return &set
}
