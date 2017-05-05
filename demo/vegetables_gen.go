package demo

// file generated by
// github.com/mh-cbon/lister
// do not edit

// Tomates implements a typed slice of Tomate
type Tomates struct{ items []Tomate }

// NewTomates creates a new typed slice of Tomate
func NewTomates() *Tomates {
	return &Tomates{items: []Tomate{}}
}

// Push appends every Tomate
func (t *Tomates) Push(x ...Tomate) *Tomates {
	t.items = append(t.items, x...)
	return t
}

// Unshift prepends every Tomate
func (t *Tomates) Unshift(x ...Tomate) *Tomates {
	t.items = append(x, t.items...)
	return t
}

// Pop removes then returns the last Tomate.
func (t *Tomates) Pop() Tomate {
	var ret Tomate
	if len(t.items) > 0 {
		ret = t.items[len(t.items)-1]
		t.items = append(t.items[:0], t.items[len(t.items)-1:]...)
	}
	return ret
}

// Shift removes then returns the first Tomate.
func (t *Tomates) Shift() Tomate {
	var ret Tomate
	if len(t.items) > 0 {
		ret = t.items[0]
		t.items = append(t.items[:0], t.items[1:]...)
	}
	return ret
}

// Index of given Tomate. It must implements Ider interface.
func (t *Tomates) Index(s Tomate) int {
	ret := -1
	for i, item := range t.items {
		if s.GetID() == item.GetID() {
			ret = i
			break
		}
	}
	return ret
}

// Contains returns true if s in is t.
func (t *Tomates) Contains(s Tomate) bool {
	return t.Index(s) > -1
}

// RemoveAt removes a Tomate at index i.
func (t *Tomates) RemoveAt(i int) bool {
	if i >= 0 && i < len(t.items) {
		t.items = append(t.items[:i], t.items[i+1:]...)
		return true
	}
	return false
}

// Remove removes given Tomate
func (t *Tomates) Remove(s Tomate) bool {
	if i := t.Index(s); i > -1 {
		t.RemoveAt(i)
		return true
	}
	return false
}

// InsertAt adds given Tomate at index i
func (t *Tomates) InsertAt(i int, s Tomate) *Tomates {
	if i < 0 || i >= len(t.items) {
		return t
	}
	res := []Tomate{}
	res = append(res, t.items[:0]...)
	res = append(res, s)
	res = append(res, t.items[i:]...)
	t.items = res
	return t
}

// Splice removes and returns a slice of Tomate, starting at start, ending at start+length.
// If any s is provided, they are inserted in place of the removed slice.
func (t *Tomates) Splice(start int, length int, s ...Tomate) []Tomate {
	var ret []Tomate
	for i := 0; i < len(t.items); i++ {
		if i >= start && i < start+length {
			ret = append(ret, t.items[i])
		}
	}
	if start >= 0 && start+length <= len(t.items) && start+length >= 0 {
		t.items = append(
			t.items[:start],
			append(s,
				t.items[start+length:]...,
			)...,
		)
	}
	return ret
}

// Slice returns a copied slice of Tomate, starting at start, ending at start+length.
func (t *Tomates) Slice(start int, length int) []Tomate {
	var ret []Tomate
	if start >= 0 && start+length <= len(t.items) && start+length >= 0 {
		ret = t.items[start : start+length]
	}
	return ret
}

// Reverse the slice.
func (t *Tomates) Reverse() *Tomates {
	for i, j := 0, len(t.items)-1; i < j; i, j = i+1, j-1 {
		t.items[i], t.items[j] = t.items[j], t.items[i]
	}
	return t
}

// Len of the slice.
func (t *Tomates) Len() int {
	return len(t.items)
}

// Set the slice.
func (t *Tomates) Set(x []Tomate) *Tomates {
	t.items = append(t.items[:0], x...)
	return t
}

// Get the slice.
func (t *Tomates) Get() []Tomate {
	return t.items
}

// At return the item at index i.
func (t *Tomates) At(i int) Tomate {
	return t.items[i]
}

// Filter return a new Tomates with all items satisfying f.
func (t *Tomates) Filter(filters ...func(Tomate) bool) *Tomates {
	ret := NewTomates()
	for _, i := range t.items {
		ok := true
		for _, f := range filters {
			ok = ok && f(i)
			if !ok {
				break
			}
		}
		if ok {
			ret.Push(i)
		}
	}
	return ret
}

// Map return a new Tomates of each items modified by f.
func (t *Tomates) Map(mappers ...func(Tomate) Tomate) *Tomates {
	ret := NewTomates()
	for _, i := range t.items {
		val := i
		for _, m := range mappers {
			val = m(val)
		}
		ret.Push(val)
	}
	return ret
}

// First returns the first value or default.
func (t *Tomates) First() Tomate {
	var ret Tomate
	if len(t.items) > 0 {
		ret = t.items[0]
	}
	return ret
}

// Last returns the last value or default.
func (t *Tomates) Last() Tomate {
	var ret Tomate
	if len(t.items) > 0 {
		ret = t.items[len(t.items)-1]
	}
	return ret
}

// Empty returns true if the slice is empty.
func (t *Tomates) Empty() bool {
	return len(t.items) == 0
}

var FilterTomates = struct {
	ByName   func(string) func(Tomate) bool
	ByWidth  func(uint64) func(Tomate) bool
	ByHeight func(uint64) func(Tomate) bool
}{
	ByName:   func(v string) func(Tomate) bool { return func(o Tomate) bool { return o.Name == v } },
	ByWidth:  func(v uint64) func(Tomate) bool { return func(o Tomate) bool { return o.Width == v } },
	ByHeight: func(v uint64) func(Tomate) bool { return func(o Tomate) bool { return o.Height == v } },
}

// Poireaux implements a typed slice of *Poireau
type Poireaux struct{ items []*Poireau }

// NewPoireaux creates a new typed slice of *Poireau
func NewPoireaux() *Poireaux {
	return &Poireaux{items: []*Poireau{}}
}

// Push appends every *Poireau
func (t *Poireaux) Push(x ...*Poireau) *Poireaux {
	t.items = append(t.items, x...)
	return t
}

// Unshift prepends every *Poireau
func (t *Poireaux) Unshift(x ...*Poireau) *Poireaux {
	t.items = append(x, t.items...)
	return t
}

// Pop removes then returns the last *Poireau.
func (t *Poireaux) Pop() *Poireau {
	var ret *Poireau
	if len(t.items) > 0 {
		ret = t.items[len(t.items)-1]
		t.items = append(t.items[:0], t.items[len(t.items)-1:]...)
	}
	return ret
}

// Shift removes then returns the first *Poireau.
func (t *Poireaux) Shift() *Poireau {
	var ret *Poireau
	if len(t.items) > 0 {
		ret = t.items[0]
		t.items = append(t.items[:0], t.items[1:]...)
	}
	return ret
}

// Index of given *Poireau. It must implements Ider interface.
func (t *Poireaux) Index(s *Poireau) int {
	ret := -1
	for i, item := range t.items {
		if s.GetID() == item.GetID() {
			ret = i
			break
		}
	}
	return ret
}

// Contains returns true if s in is t.
func (t *Poireaux) Contains(s *Poireau) bool {
	return t.Index(s) > -1
}

// RemoveAt removes a *Poireau at index i.
func (t *Poireaux) RemoveAt(i int) bool {
	if i >= 0 && i < len(t.items) {
		t.items = append(t.items[:i], t.items[i+1:]...)
		return true
	}
	return false
}

// Remove removes given *Poireau
func (t *Poireaux) Remove(s *Poireau) bool {
	if i := t.Index(s); i > -1 {
		t.RemoveAt(i)
		return true
	}
	return false
}

// InsertAt adds given *Poireau at index i
func (t *Poireaux) InsertAt(i int, s *Poireau) *Poireaux {
	if i < 0 || i >= len(t.items) {
		return t
	}
	res := []*Poireau{}
	res = append(res, t.items[:0]...)
	res = append(res, s)
	res = append(res, t.items[i:]...)
	t.items = res
	return t
}

// Splice removes and returns a slice of *Poireau, starting at start, ending at start+length.
// If any s is provided, they are inserted in place of the removed slice.
func (t *Poireaux) Splice(start int, length int, s ...*Poireau) []*Poireau {
	var ret []*Poireau
	for i := 0; i < len(t.items); i++ {
		if i >= start && i < start+length {
			ret = append(ret, t.items[i])
		}
	}
	if start >= 0 && start+length <= len(t.items) && start+length >= 0 {
		t.items = append(
			t.items[:start],
			append(s,
				t.items[start+length:]...,
			)...,
		)
	}
	return ret
}

// Slice returns a copied slice of *Poireau, starting at start, ending at start+length.
func (t *Poireaux) Slice(start int, length int) []*Poireau {
	var ret []*Poireau
	if start >= 0 && start+length <= len(t.items) && start+length >= 0 {
		ret = t.items[start : start+length]
	}
	return ret
}

// Reverse the slice.
func (t *Poireaux) Reverse() *Poireaux {
	for i, j := 0, len(t.items)-1; i < j; i, j = i+1, j-1 {
		t.items[i], t.items[j] = t.items[j], t.items[i]
	}
	return t
}

// Len of the slice.
func (t *Poireaux) Len() int {
	return len(t.items)
}

// Set the slice.
func (t *Poireaux) Set(x []*Poireau) *Poireaux {
	t.items = append(t.items[:0], x...)
	return t
}

// Get the slice.
func (t *Poireaux) Get() []*Poireau {
	return t.items
}

// At return the item at index i.
func (t *Poireaux) At(i int) *Poireau {
	return t.items[i]
}

// Filter return a new Poireaux with all items satisfying f.
func (t *Poireaux) Filter(filters ...func(*Poireau) bool) *Poireaux {
	ret := NewPoireaux()
	for _, i := range t.items {
		ok := true
		for _, f := range filters {
			ok = ok && f(i)
			if !ok {
				break
			}
		}
		if ok {
			ret.Push(i)
		}
	}
	return ret
}

// Map return a new Poireaux of each items modified by f.
func (t *Poireaux) Map(mappers ...func(*Poireau) *Poireau) *Poireaux {
	ret := NewPoireaux()
	for _, i := range t.items {
		val := i
		for _, m := range mappers {
			val = m(val)
			if val == nil {
				break
			}
		}
		if val != nil {
			ret.Push(val)
		}
	}
	return ret
}

// First returns the first value or default.
func (t *Poireaux) First() *Poireau {
	var ret *Poireau
	if len(t.items) > 0 {
		ret = t.items[0]
	}
	return ret
}

// Last returns the last value or default.
func (t *Poireaux) Last() *Poireau {
	var ret *Poireau
	if len(t.items) > 0 {
		ret = t.items[len(t.items)-1]
	}
	return ret
}

// Empty returns true if the slice is empty.
func (t *Poireaux) Empty() bool {
	return len(t.items) == 0
}

var FilterPoireaux = struct {
	ByName   func(string) func(*Poireau) bool
	ByWidth  func(uint64) func(*Poireau) bool
	ByHeight func(uint64) func(*Poireau) bool
}{
	ByName:   func(v string) func(*Poireau) bool { return func(o *Poireau) bool { return o.Name == v } },
	ByWidth:  func(v uint64) func(*Poireau) bool { return func(o *Poireau) bool { return o.Width == v } },
	ByHeight: func(v uint64) func(*Poireau) bool { return func(o *Poireau) bool { return o.Height == v } },
}
