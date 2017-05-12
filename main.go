// Package lister generates typed slice.
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mh-cbon/lister/utils"

	"github.com/mh-cbon/astutil"
)

var name = "lister"
var version = "0.0.0"

//go:generate lister string:gen/StringSlice

func main() {

	var help bool
	var h bool
	var ver bool
	var v bool
	var outPkg string
	flag.BoolVar(&help, "help", false, "Show help.")
	flag.BoolVar(&h, "h", false, "Show help.")
	flag.BoolVar(&ver, "version", false, "Show version.")
	flag.BoolVar(&v, "v", false, "Show version.")
	flag.StringVar(&outPkg, "p", "", "Package name of the new code.")

	flag.Parse()

	if ver || v {
		showVer()
		return
	}
	if help || h {
		showHelp()
		return
	}

	if flag.NArg() < 1 {
		showHelp()
		return
	}
	args := flag.Args()

	pkgToLoad := getPkgToLoad()
	prog := astutil.GetProgram(pkgToLoad).Package(pkgToLoad)

	todos, err := utils.NewTransformsArgs(outPkg).Parse(args)
	if err != nil {
		panic(err)
	}

	filesOut := utils.NewFilesOut("github.com/mh-cbon/" + name)

	for _, todo := range todos.Args {
		fileOut := filesOut.Get(todo.ToPath)
		fileOut.PkgName = todo.ToPkgPath

		processType(&fileOut.Body, todo)
		srcName := todo.FromTypeName
		if astutil.IsBasic(srcName) == false {
			foundStruct := astutil.GetStruct(prog, astutil.GetUnpointedType(srcName))
			if foundStruct == nil {
				log.Println("Can not locate the type " + srcName)
				continue
			}
			processFilter(&fileOut.Body, foundStruct, todo)
		}
	}

	for _, f := range filesOut.Files {
		if err := f.Write(); err != nil {
			log.Println(err)
		}
	}
}

func showVer() {
	fmt.Printf("%v %v\n", name, version)
}

func showHelp() {
	showVer()
	fmt.Println()
	fmt.Println("Usage")
	fmt.Println()
	fmt.Printf("  %v [-p name] [...types]\n\n", name)
	fmt.Printf("  types:  A list of types such as src:dst.\n")
	fmt.Printf("          A type is defined by its package path and its type name,\n")
	fmt.Printf("          [pkgpath/]name\n")
	fmt.Printf("          If the Package path is empty, it is set to the package name being generated.\n")
	// fmt.Printf("          If the Package path is a directory relative to the cwd, and the Package name is not provided\n")
	// fmt.Printf("          the package path is set to this relative directory,\n")
	// fmt.Printf("          the package name is set to the name of this directory.\n")
	fmt.Printf("          Name can be a valid type identifier such as TypeName, *TypeName, []TypeName \n")
	fmt.Printf("  -p:     The name of the package output.\n")
	fmt.Println()
}

func getPkgToLoad() string {
	gopath := filepath.Join(os.Getenv("GOPATH"), "src")
	pkgToLoad, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if len(pkgToLoad) < len(gopath) || pkgToLoad[:len(gopath)] != gopath {
		panic(fmt.Errorf("unexpected gopath %q", gopath))
	}
	return pkgToLoad[len(gopath)+1:]
}

func processFilter(dest io.Writer, s *ast.StructType, todo utils.TransformArg) {

	srcName := todo.FromTypeName
	destName := todo.ToTypeName

	destConcrete := astutil.GetUnpointedType(destName)

	props := astutil.StructProps(s)

	newStructProps := ""
	for _, prop := range props {
		//todo: find a way to detect if the type implements Eq or something like this.
		propType := prop["type"]
		if !astutil.IsArrayType(propType) {
			propName := prop["name"]
			newStructProps += fmt.Sprintf("By%v func(%v) func (%v) bool", propName, propType, srcName)
			newStructProps += "\n"
		}
	}

	if newStructProps != "" {
		fmt.Fprintf(dest, "// Filter%v provides filters for a struct.\n", destConcrete)
		fmt.Fprintf(dest, `var Filter%v = struct {`, destConcrete)
		fmt.Fprintln(dest)
		fmt.Fprintln(dest, newStructProps+"\n")
		fmt.Fprintln(dest, "}{")
		for _, prop := range props {
			//todo: find a way to detect if the type implements Eq or something like this.
			propType := prop["type"]
			if !astutil.IsArrayType(propType) {
				propName := prop["name"]
				fmt.Fprintf(dest, `By%v: func(v %v) func(%v) bool {`, propName, propType, srcName)
				fmt.Fprintf(dest, `	return func(o %v) bool {`, srcName)
				fmt.Fprintf(dest, `	return o.%v==v`, propName)
				fmt.Fprintf(dest, `	}`)
				fmt.Fprintf(dest, `},`)
				fmt.Fprintln(dest, "")
			}
		}
		fmt.Fprintln(dest)
		fmt.Fprintln(dest, "}")
	}
}

func processType(dest io.Writer, todo utils.TransformArg) {

	srcName := todo.FromTypeName
	destName := todo.ToTypeName

	destPointed := astutil.GetPointedType(destName)
	destConcrete := astutil.GetUnpointedType(destName)
	srcIsPointer := astutil.IsAPointedType(srcName)
	srcIsBasic := astutil.IsBasic(srcName)

	fmt.Fprintf(dest, `// %v implements a typed slice of %v`, destConcrete, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `type %v struct {items []%v}`, destConcrete, srcName)
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// New%v creates a new typed slice of %v`, destConcrete, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func New%v() %v {
 return &%v{items: []%v{}}
}`, destConcrete, destPointed, destConcrete, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Push appends every %v`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Push(x ...%v) %v {
 t.items = append(t.items, x...)
 return t
}`, destPointed, srcName, destPointed)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Unshift prepends every %v`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Unshift(x ...%v) %v {
	t.items = append(x, t.items...)
	return t
}`, destPointed, srcName, destPointed)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Pop removes then returns the last %v.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Pop() %v {
 var ret %v
 if len(t.items)>0 {
  ret = t.items[len(t.items)-1]
  t.items = append(t.items[:0], t.items[len(t.items)-1:]...)
 }
 return ret
}`, destPointed, srcName, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Shift removes then returns the first %v.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Shift() %v {
  var ret %v
  if len(t.items)>0 {
    ret = t.items[0]
    t.items = append(t.items[:0], t.items[1:]...)
  }
  return ret
}`, destPointed, srcName, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Index of given %v. It must implements Ider interface.`, srcName)
	fmt.Fprintln(dest, "")
	if srcIsBasic == false {
		fmt.Fprintf(dest, `func (t %v) Index(s %v) int {
	  ret := -1
	  for i,item:= range t.items {
			if s.GetID()==item.GetID() {
				ret = i
				break
			}
	  }
	  return ret
	}`, destPointed, srcName)
	} else if srcIsPointer && srcIsBasic { // needed ?
		fmt.Fprintf(dest, `func (t %v) Index(s %v) int {
	  ret := -1
	  for i,item:= range t.items {
			if *s==*item {
				ret = i
				break
			}
	  }
	  return ret
	}`, destPointed, srcName)
	} else {
		fmt.Fprintf(dest, `func (t %v) Index(s %v) int {
	  ret := -1
	  for i,item:= range t.items {
			if s==item {
				ret = i
				break
			}
	  }
	  return ret
	}`, destPointed, srcName)
	}
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Contains returns true if s in is t.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Contains(s %v) bool {
  return t.Index(s)>-1
}`, destPointed, srcName)
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// RemoveAt removes a %v at index i.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) RemoveAt(i int) bool {
  if i>=0 && i<len(t.items) {
    t.items = append(t.items[:i], t.items[i+1:]...)
		return true
  }
  return false
}`, destPointed)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Remove removes given %v`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Remove(s %v) bool {
  if i := t.Index(s); i > -1 {
    t.RemoveAt(i)
		return true
  }
  return false
}`, destPointed, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// InsertAt adds given %v at index i`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) InsertAt(i int, s %v) %v {
	if i<0 || i >= len(t.items) {
		return t
	}
	res := []%v{}
	res = append(res, t.items[:0]...)
	res = append(res, s)
	res = append(res, t.items[i:]...)
	t.items = res
	return t
}`, destPointed, srcName, destPointed, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Splice removes and returns a slice of %v, starting at start, ending at start+length.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `// If any s is provided, they are inserted in place of the removed slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Splice(start int, length int, s ...%v) []%v {
	var ret []%v
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
}`, destPointed, srcName, srcName, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Slice returns a copied slice of %v, starting at start, ending at start+length.`, srcName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Slice(start int, length int) []%v {
  var ret []%v
	if start >= 0 && start+length <= len(t.items) && start+length >= 0 {
		ret = t.items[start:start+length]
	}
	return ret
}`, destPointed, srcName, srcName)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Reverse the slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Reverse() %v {
  for i, j := 0, len(t.items)-1; i < j; i, j = i+1, j-1 {
    t.items[i], t.items[j] = t.items[j], t.items[i]
  }
  return t
}`, destPointed, destPointed)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Len of the slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Len() int {
  return len(t.items)
}`, destPointed)

	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Set the slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Set(x []%v) %v {
  t.items = append(t.items[:0], x...)
	return t
}`, destPointed, srcName, destPointed)
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Get the slice.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Get() []%v {
	return t.items
}`, destPointed, srcName)
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// At return the item at index i.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) At(i int) %v {
	return t.items[i]
}`, destPointed, srcName)
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Filter return a new %v with all items satisfying f.`, destName)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Filter(filters ...func(%v) bool) %v {
	ret := New%v()
	for _, i := range t.items {
		ok := true
		for _, f := range filters {
			ok = ok && f(i)
			if ! ok {
				break
			}
		}
		if ok {
			ret.Push(i)
		}
	}
	return ret
}`, destPointed, srcName, destPointed, destConcrete)
	fmt.Fprintln(dest, "")

	// todod: handle more cases like ArayType etc.
	fmt.Fprintf(dest, `// Map return a new %v of each items modified by f.`, destName)
	fmt.Fprintln(dest, "")
	if astutil.IsAPointedType(srcName) {
		fmt.Fprintf(dest, `func (t %v) Map(mappers ...func(%v) %v) %v {
		ret := New%v()
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
	}`, destPointed, srcName, srcName, destPointed, destConcrete)
	} else {
		fmt.Fprintf(dest, `func (t %v) Map(mappers ...func(%v) %v) %v {
		ret := New%v()
		for _, i := range t.items {
			val := i
			for _, m := range mappers {
				val = m(val)
			}
			ret.Push(val)
		}
		return ret
	}`, destPointed, srcName, srcName, destPointed, destConcrete)
	}
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// First returns the first value or default.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) First() %v {
	var ret %v
	if len(t.items)>0 {
		ret = t.items[0]
	}
	return ret
}`, destPointed, srcName, srcName)
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Last returns the last value or default.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Last() %v {
	var ret %v
	if len(t.items)>0 {
		ret = t.items[len(t.items)-1]
	}
	return ret
}`, destPointed, srcName, srcName)
	fmt.Fprintln(dest, "")

	fmt.Fprintf(dest, `// Empty returns true if the slice is empty.`)
	fmt.Fprintln(dest, "")
	fmt.Fprintf(dest, `func (t %v) Empty() bool {
	return len(t.items)==0
}`, destPointed)
	fmt.Fprintln(dest, "")

	fmt.Fprintln(dest, "")
}
