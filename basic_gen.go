package main

// file generated by
// github.com.mh-cbon/lister
// do not edit

// StringSlice implements a typed slice of string
type StringSlice []string

// NewStringSlice creates a new typed slice of string
func NewStringSlice() *StringSlice {
	return &StringSlice{}
}

// Push appends every string
func (t *StringSlice) Push(x ...string) *StringSlice {
	items := *t
	items = append(items, x...)
	return t.Set(items)
}

// Unshift prepends every string
func (t *StringSlice) Unshift(x ...string) *StringSlice {
	items := *t
	items = append(x, items...)
	return t.Set(items)
}

// Pop removes then reutrns the last string.
func (t *StringSlice) Pop() string {
	var ret string
	items := *t
	if len(items) > 0 {
		ret = items[len(items)-1]
		items = append(items[:0], items[len(items)-1:]...)
		t.Set(items)
	}
	return ret
}

// Shift removes then reutrns the first string.
func (t *StringSlice) Shift() string {
	var ret string
	items := *t
	if len(items) > 0 {
		ret = items[0]
		items = append(items[:0], items[1:]...)
	}
	t.Set(items)
	return ret
}

// Index of given string. It must implements Ider interface.
func (t *StringSlice) Index(s string) int {
	ret := -1
	items := *t
	for i, item := range items {
		if s == item {
			ret = i
			break
		}
	}
	return ret
}

// RemoveAt removes a string at index i.
func (t *StringSlice) RemoveAt(i int) bool {
	items := *t
	if i < len(items) {
		items = append(items[:i], items[i+1:]...)
		t.Set(items)
		return true
	}
	return false
}

// Remove removes given string
func (t *StringSlice) Remove(s string) bool {
	if i := t.Index(s); i > -1 {
		t.RemoveAt(i)
		return true
	}
	return false
}

// InsertAt adds given string at index i
func (t *StringSlice) InsertAt(i int, s string) *StringSlice {
	items := *t
	items = append(
		items[:i],
		append(
			append(items[:0], s),
			items[i+1:]...,
		)...,
	)
	return t.Set(items)
}

// Splice removes and returns a slice of string, starting at start, ending at start+length.
// If any s is provided, they are inserted in place of the removed slice.
func (t *StringSlice) Splice(start int, length int, s ...string) []string {
	items := *t
	ret := items[start : start+length]
	items = append(items[:start], append(s, items[start+length:]...)...)
	t.Set(items)
	return ret
}

// Slice returns a copied slice of string, starting at start, ending at start+length.
func (t *StringSlice) Slice(start int, length int) []string {
	items := *t
	return items[start : start+length]
}

// Reverse the slice.
func (t *StringSlice) Reverse() *StringSlice {
	items := *t
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
	return t.Set(items)
}

// Len of the slice.
func (t *StringSlice) Len() int {
	return len(*t)
}

// Set the slice.
func (t *StringSlice) Set(x []string) *StringSlice {
	items := *t
	items = append(items[:0], x...)
	t = &items
	return t
}
