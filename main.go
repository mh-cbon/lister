// Package lister is a generator to generate typed slice.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var name = "lister"
var version = "0.0.0"

func main() {

	var help bool
	var h bool
	var ver bool
	var v bool
	var p string
	flag.BoolVar(&help, "help", false, "Show help.")
	flag.BoolVar(&h, "h", false, "Show help.")
	flag.BoolVar(&ver, "version", false, "Show version.")
	flag.BoolVar(&v, "v", false, "Show version.")
	flag.StringVar(&p, "p", os.Getenv("GOPACKAGE"), "Package name of the new code.")

	flag.Parse()

	if ver || v {
		showVer()
		return
	}
	if help || h {
		showHelp()
		return
	}

	if flag.NArg() < 2 {
		panic("wrong usage")
	}
	args := flag.Args()

	dest := os.Stdout

	o := args[0]
	restargs := args[1:]

	if o != "-" {
		f, err := os.Create(o)
		if err != nil {
			panic(err)
		}
		dest = f
		defer func() {
			f.Close()
			exec.Command("go", "fmt", args[0]).Run()
		}()
	}

	fmt.Fprintf(dest, "package %v\n\n", p)
	fmt.Fprintln(dest, `// file generated by`)
	fmt.Fprintln(dest, `// github.com.mh-cbon/`+name)
	fmt.Fprintln(dest, `// do not edit`)
	fmt.Fprintln(dest, ``)

	for _, todo := range restargs {
		srcName, destName := splitTypeArg(todo)
		res := processType(destName, srcName)
		io.Copy(dest, &res)
	}
}

func splitTypeArg(todo string) (src string, dest string) {
	y := strings.Split(todo, ":")
	if len(y) != 2 {
		panic("wrong name " + todo)
	}
	return y[0], y[1]
}

func showVer() {
	fmt.Printf("%v %v\n", name, version)
}

func showHelp() {
	showVer()
	fmt.Println()
	fmt.Println("Usage")
	fmt.Println()
	fmt.Printf("	%v [-p name] [out] [...types]\n\n", name)
	fmt.Printf("	out: 	Output destination of the results, use '-' for stdout.\n")
	fmt.Printf("	types:	A list of types such as src:dst.\n")
	fmt.Printf("	-p:			The name of the package output.\n")
	fmt.Println()
}

func processType(destName, srcName string) bytes.Buffer {

	var b bytes.Buffer
	dest := &b

	fmt.Fprintf(dest, `// %v implements a typed slice of %v`, getUnpointedType(destName), srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `type %v []%v`, getUnpointedType(destName), srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// New%v creates a new typed slice of %v`, destName, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func New%v() %v {
 return &%v{}
}`, getUnpointedType(destName), getPointedType(destName), getUnpointedType(destName))

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Push appends every %v`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Push(x ...%v) %v {
 items := *t
 items = append(items, x...)
 return t.Set(items)
}`, getPointedType(destName), srcName, getPointedType(destName))

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Unshift prepends every %v`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Unshift(x ...%v) %v {
 items := *t
 items = append(x, items...)
 return t.Set(items)
}`, getPointedType(destName), srcName, getPointedType(destName))

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Pop removes then reutrns the last %v.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Pop() %v {
 var ret %v
 items := *t
 if len(items)>0 {
  ret = items[len(items)-1]
  items = append(items[:0], items[len(items)-1:]...)
  t.Set(items)
 }
 return ret
}`, getPointedType(destName), srcName, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Shift removes then reutrns the first %v.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Shift() %v {
  var ret %v
  items := *t
  if len(items)>0 {
    ret = items[0]
    items = append(items[:0], items[1:]...)
  }
  t.Set(items)
  return ret
}`, getPointedType(destName), srcName, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Index of given %v. It must implements Ider interface.`, srcName)
	fmt.Fprintln(dest, "")
	if isBasic(srcName) == false {
		fmt.Fprintf(dest, `func (t %v) Index(s %v) int {
	  ret := -1
	  items := *t
	  for i,item:= range items {
			if s.GetID()==item.GetID() {
				ret = i
				break
			}
	  }
	  return ret
	}`, getPointedType(destName), srcName)
	} else if isAPointedType(srcName) && isBasic(srcName) { // needed ?
		fmt.Fprintf(dest, `func (t %v) Index(s %v) int {
	  ret := -1
	  items := *t
	  for i,item:= range items {
			if *s==*item {
				ret = i
				break
			}
	  }
	  return ret
	}`, getPointedType(destName), srcName)
	} else {
		fmt.Fprintf(dest, `func (t %v) Index(s %v) int {
	  ret := -1
	  items := *t
	  for i,item:= range items {
			if s==item {
				ret = i
				break
			}
	  }
	  return ret
	}`, getPointedType(destName), srcName)
	}

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// RemoveAt removes a %v at index i.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) RemoveAt(i int) bool {
  items := *t
  if i<len(items) {
    items = append(items[:i], items[i+1:]...)
	  t.Set(items)
		return true
  }
  return false
}`, getPointedType(destName))

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Remove removes given %v`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Remove(s %v) bool {
  if i := t.Index(s); i > -1 {
    t.RemoveAt(i)
		return true
  }
  return false
}`, getPointedType(destName), srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// InsertAt adds given %v at index i`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) InsertAt(i int, s %v) %v {
  items := *t
  items = append(
    items[:i],
    append(
      append(items[:0], s),
      items[i+1:]...
    )...,
  )
  return t.Set(items)
}`, getPointedType(destName), srcName, getPointedType(destName))

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Splice removes and returns a slice of %v, starting at start, ending at start+length.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `// If any s is provided, they are inserted in place of the removed slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Splice(start int, length int, s ...%v) []%v {
		items := *t
  ret := items[start:start+length]
  items = append(items[:start], append(s, items[start+length:]...)...)
  t.Set(items)
  return ret
}`, getPointedType(destName), srcName, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Slice returns a copied slice of %v, starting at start, ending at start+length.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Slice(start int, length int) []%v {
  items := *t
  return items[start:start+length]
}`, getPointedType(destName), srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Reverse the slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Reverse() %v {
  items := *t
  for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
    items[i], items[j] = items[j], items[i]
  }
  return t.Set(items)
}`, getPointedType(destName), getPointedType(destName))

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Len of the slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Len() int {
  return len(*t)
}`, getPointedType(destName))

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Set the slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Set(x []%v) %v {
	items := *t
  items = append(items[:0], x...)
	t = &items
	return t
}`, getPointedType(destName), srcName, getPointedType(destName))

	fmt.Fprintln(dest, "")

	return b
}

func isAPointedType(t string) bool {
	return t[0] == '*'
}

func getUnpointedType(t string) string {
	if isAPointedType(t) {
		return t[1:]
	}
	return t
}
func getPointedType(t string) string {
	if !isAPointedType(t) {
		t = "*" + t
	}
	return t
}

func isBasic(t string) bool {
	if isAPointedType(t) {
		t = t[1:]
	}
	//go:generate lister basic_gen.go string:StringSlice
	basicTypes := NewStringSlice().Push(
		"string",
		"int",
		"uint",
		"int8",
		"uint8",
		"int16",
		"uint16",
		"int32",
		"uint32",
		"int64",
		"uint64",
		"float",
		"float64",
		"ufloat",
		"ufloat64",
	)
	return basicTypes.Index(t) > -1
}
