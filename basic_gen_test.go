package main

import (
	"testing"

	"github.com/mh-cbon/lister/gen"
)

func TestPush(t *testing.T) {
	s := gen.NewStringSlice()
	s.Push("", "")
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestUnshift(t *testing.T) {
	s := gen.NewStringSlice()
	s.Unshift("", "")
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestPop(t *testing.T) {
	s := gen.NewStringSlice()
	sgot := s.Unshift("first", "last").Pop()
	want := 1
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestNotPop(t *testing.T) {
	s := gen.NewStringSlice()
	sgot := s.Pop()
	want := 0
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := ""
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestShift(t *testing.T) {
	s := gen.NewStringSlice()
	sgot := s.Unshift("first", "last").Shift()
	want := 1
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "first"
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestNotShift(t *testing.T) {
	s := gen.NewStringSlice()
	sgot := s.Shift()
	want := 0
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := ""
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestIndex(t *testing.T) {
	s := gen.NewStringSlice()
	igot := s.Unshift("first", "last").Index("last")
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	iwant := 1
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestNotIndex(t *testing.T) {
	s := gen.NewStringSlice()
	igot := s.Unshift("first", "last").Index("wcwxcxcv")
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	iwant := -1
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestRemoveAt(t *testing.T) {
	s := gen.NewStringSlice()
	igot := s.Unshift("first", "last").RemoveAt(0)
	want := 1
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	iwant := true
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestNotRemoveAt(t *testing.T) {
	s := gen.NewStringSlice()
	igot := s.Unshift("first", "last").RemoveAt(-1)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	iwant := false
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestNotRemoveAt2(t *testing.T) {
	s := gen.NewStringSlice()
	igot := s.Unshift("first", "last").RemoveAt(2)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	iwant := false
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestInsertAt(t *testing.T) {
	s := gen.NewStringSlice()
	s.Unshift("first", "last").InsertAt(0, "new")
	want := 3
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "new"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	swant = "first"
	sgot = s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	swant = "last"
	sgot = s.At(2)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestNotInsertAt(t *testing.T) {
	s := gen.NewStringSlice()
	s.Unshift("first", "last").InsertAt(-1, "new")
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "first"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	swant = "last"
	sgot = s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestSplice(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Splice(0, 1)
	want := 1
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	sswant := []string{"first"}
	if sswant[0] != ssgot[0] {
		t.Errorf("want %v got %v", sswant, ssgot)
	}
	lwant := 1
	lgot := len(ssgot)
	if lwant != lgot {
		t.Errorf("want %v got %v", lwant, lgot)
	}
	swant := "last"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestSplice1(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Splice(1, 1)
	want := 1
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	sswant := []string{"last"}
	if sswant[0] != ssgot[0] {
		t.Errorf("want %v got %v", sswant, ssgot)
	}
	lwant := 1
	lgot := len(ssgot)
	if lwant != lgot {
		t.Errorf("want %v got %v", lwant, lgot)
	}
	swant := "first"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestSpliceAdd(t *testing.T) {
	s := gen.NewStringSlice()
	sgot := s.Unshift("first", "last").Splice(1, 1, "t", "r")[0]
	want := 3
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	swant = "t"
	sgot = s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	swant = "r"
	sgot = s.At(2)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestNotSplice(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Splice(-1, 1)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	iwant := 0
	igot := len(ssgot)
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestNotSplice2(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Splice(0, -5)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	iwant := 0
	igot := len(ssgot)
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestNotSplice3(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Splice(10, 5)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	iwant := 0
	igot := len(ssgot)
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestSlice(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Slice(0, 1)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	sswant := []string{"first"}
	if sswant[0] != ssgot[0] {
		t.Errorf("want %v got %v", sswant, ssgot)
	}
	lwant := 1
	lgot := len(ssgot)
	if lwant != lgot {
		t.Errorf("want %v got %v", lwant, lgot)
	}
	swant := "first"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestSlice1(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Slice(1, 1)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	sswant := []string{"last"}
	if sswant[0] != ssgot[0] {
		t.Errorf("want %v got %v", sswant, ssgot)
	}
	lwant := 1
	lgot := len(ssgot)
	if lwant != lgot {
		t.Errorf("want %v got %v", lwant, lgot)
	}
	swant := "first"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}

func TestNotSlice(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Slice(-1, 1)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	iwant := 0
	igot := len(ssgot)
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestNotSlice2(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Splice(0, -5)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	iwant := 0
	igot := len(ssgot)
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestNotSlice3(t *testing.T) {
	s := gen.NewStringSlice()
	ssgot := s.Unshift("first", "last").Slice(10, 5)
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	iwant := 0
	igot := len(ssgot)
	if iwant != igot {
		t.Errorf("want %v got %v", iwant, igot)
	}
}

func TestReverse(t *testing.T) {
	s := gen.NewStringSlice()
	s.Unshift("first", "last").Reverse()
	want := 2
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
	swant = "first"
	sgot = s.At(1)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}
func TestEmptyReverse(t *testing.T) {
	s := gen.NewStringSlice()
	s.Unshift().Reverse()
	want := 0
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestFilter(t *testing.T) {
	s := gen.NewStringSlice()
	s = s.Unshift("first", "last").Filter(func(s string) bool { return s == "last" })
	want := 1
	got := s.Len()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
	swant := "last"
	sgot := s.At(0)
	if swant != sgot {
		t.Errorf("want %v got %v", swant, sgot)
	}
}
