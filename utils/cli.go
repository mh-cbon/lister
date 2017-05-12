package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// FilesOut ...
type FilesOut struct {
	Files         []*FileOut
	GeneratorName string
}

// NewFilesOut ...
func NewFilesOut(name string) *FilesOut {
	return &FilesOut{GeneratorName: name}
}

// Get the file handler matching path s
func (f *FilesOut) Get(s string) *FileOut {
	for _, p := range f.Files {
		if p.Path == strings.ToLower(s) {
			return p
		}
	}
	r := &FileOut{GeneratorName: f.GeneratorName, Path: s}
	f.Files = append(f.Files, r)
	return r
}

// FileOut ...
type FileOut struct {
	GeneratorName string
	PkgName       string
	Path          string
	Body          bytes.Buffer
}

func (f *FileOut) Write() error {
	o := f.Path
	dest := os.Stdout
	if o != "-" {
		os.MkdirAll(filepath.Dir(o), os.ModePerm)
		f, err := os.Create(o)
		if err != nil {
			panic(err)
		}
		dest = f
		defer func() {
			f.Close()
			exec.Command("go", "fmt", o).Run()
		}()
	}

	fmt.Fprintf(dest, "package %v\n\n", f.PkgName)
	fmt.Fprintln(dest, `// file generated by`)
	fmt.Fprintln(dest, `// `+f.GeneratorName)
	fmt.Fprintln(dest, `// do not edit`)
	fmt.Fprintln(dest, ``)

	_, err := io.Copy(dest, &f.Body)
	return err
}

// TransformArgs parse cli args.
type TransformArgs struct {
	PkgBase string
	Args    []TransformArg
}

// NewTransformsArgs ...
func NewTransformsArgs(outPkg string) TransformArgs {
	if outPkg == "" {
		outPkg = os.Getenv("GOPACKAGE")
	}
	return TransformArgs{PkgBase: outPkg}
}

// TransformArg is a parsed cli arg.
type TransformArg struct {
	FromPkgPath  string
	FromTypeName string
	ToPkgPath    string
	ToTypeName   string
	ToPath       string
}

func (t TransformArg) String() string {
	ret := ""
	ret += fmt.Sprintln("todo from pkg path:", t.FromPkgPath)
	ret += fmt.Sprintln("todo from type:", t.FromTypeName)
	ret += fmt.Sprintln("todo to pkg path:", t.ToPkgPath)
	ret += fmt.Sprintln("todo to type:", t.ToTypeName)
	ret += fmt.Sprintln("todo save path:", t.ToPath)
	return ret
}

// Parse cli arguments.
func (t TransformArgs) Parse(args []string) (TransformArgs, error) {
	for _, arg := range args {
		y := strings.Split(arg, ":")
		if len(y) != 2 {
			return t, fmt.Errorf("wrong name %q", arg)
		}

		c := TransformArg{}
		c.FromTypeName = filepath.Base(y[0])
		c.FromPkgPath = t.PkgBase
		c.ToPkgPath = t.PkgBase
		c.ToTypeName = filepath.Base(y[1])

		if strings.Index(y[0], "/") > -1 {
			// test if the dir exists locally to the package being generated,
			// if so, update the import path to its absolute path,
			// if not, assume it is already an absolute package path.
			d := filepath.Dir(y[0])
			if _, err := os.Stat(d); os.IsExist(err) {
				c.FromPkgPath = d
			}
		}

		if strings.Index(y[1], "/") > -1 {
			// if the package path contains a /,
			// build a new out package made of t.base+p,
			// otherwise, it is the package being generated
			d := filepath.Dir(y[1])
			c.ToPkgPath = d
		}
		c.ToPath = filepath.Join(c.ToPkgPath, c.ToTypeName+".go")
		c.ToPath = strings.ToLower(c.ToPath)
		t.Args = append(t.Args, c)
	}
	return t, nil
}
