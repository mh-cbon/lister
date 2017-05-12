// Package demo demonstrates usage of Lister.
package demo

//go:generate lister Tomate:gen/Tomates *Poireau:gen/Poireaux

// Tomate is a struct ot describe a Tomate.
type Tomate struct {
	Name   string
	Width  uint64
	Height uint64
}

// GetID of a Tomate
func (t Tomate) GetID() string {
	return t.Name
}

// Poireau is a struct ot describe a Poireau.
type Poireau struct {
	Name   string
	Width  uint64
	Height uint64
}

// GetID of a Poireau
func (t *Poireau) GetID() string {
	return t.Name
}
