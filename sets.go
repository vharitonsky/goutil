package goutil

type Set struct {
	contents map[interface{}]bool
}

func NewSet(elements []interface{}) *Set {
	set := Set{make(map[interface{}]bool, len(elements))}
	for _, el := range elements {
		set.contents[el] = true
	}
	return &set
}

func (s_left *Set) Equals(s_right *Set) bool {
	if len(s_left.contents) != len(s_right.contents) {
		return false
	}
	for el, _ := range s_left.contents {
		_, found := s_right.contents[el]
		if !found {
			return false
		}
	}
	return true
}

func (s_left *Set) Intersection(s_right *Set) *Set {
	set := Set{make(map[interface{}]bool, len(s_left.contents)+len(s_right.contents))}
	for el, _ := range s_left.contents {
		_, found := s_right.contents[el]
		if found {
			set.contents[el] = true
		}
	}
	return &set
}

func (s_left *Set) Union(s_right *Set) *Set {
	set := Set{make(map[interface{}]bool, len(s_left.contents)+len(s_right.contents))}
	for el, _ := range s_left.contents {
		set.contents[el] = true
	}
	for el, _ := range s_right.contents {
		set.contents[el] = true
	}
	return &set
}

func (s_left *Set) Difference(s_right *Set) *Set {
	proposed_len := len(s_left.contents) - len(s_right.contents)
	if proposed_len < 0 {
		proposed_len = 0
	}
	set := Set{make(map[interface{}]bool, proposed_len)}
	for el, _ := range s_left.contents {
		_, found := s_right.contents[el]
		if !found {
			set.contents[el] = true
		}
	}
	return &set
}

func (s *Set) Add(el interface{}) bool {
	_, found := s.contents[el]
	if found {
		return false
	}
	s.contents[el] = true
	return true
}

func (s *Set) Remove(el interface{}) bool {
	_, found := s.contents[el]
	if !found {
		return false
	}
	delete(s.contents, el)
	return true
}

func (s *Set) Pop() interface{} {
	var popped bool = false
	var popped_element interface{}
	for el, _ := range s.contents {
		popped = true
		popped_element = el
		break
	}
	if !popped {
		return nil
	}
	delete(s.contents, popped_element)
	return popped_element
}

func (s *Set) IsEmpty() bool {
	return len(s.contents) == 0
}
