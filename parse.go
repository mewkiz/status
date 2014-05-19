package status

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Parse parses the package's source and updates it's CanParse status.
func (pkg *Package) Parse() (err error) {
	pkg.CanParse = false
	fset := token.NewFileSet()
	mode := parser.ParseComments
	astPkgs, err := parser.ParseDir(fset, pkg.absPath(), isPkgFile, mode)
	if err != nil {
		return fmt.Errorf("Parse failed: %s", err.Error())
	}
	var astPkg *ast.Package
	if len(astPkgs) == 1 {
		// select the ast package.
		for _, p := range astPkgs {
			astPkg = p
		}
	} else if len(astPkgs) > 1 {
		// return an error if there are more than one package.
		var buf bytes.Buffer
		for _, p := range astPkgs {
			if buf.Len() > 0 {
				fmt.Fprintf(&buf, ", ")
			}
			fmt.Fprintf(&buf, p.Name)
		}
		return fmt.Errorf("Parse failed. '%s' contains more than one package: %s", pkg.absPath(), buf.Bytes())
	}
	pkg.d = doc.New(astPkg, filepath.Clean(pkg.Path), 0)
	pkg.CanParse = true
	return nil
}

// absPath returns the absolute import path of the package.
func (pkg *Package) absPath() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", pkg.Path)
}

// TODO(u): Make use of go/build and remove isPkgFile and isGoFile.

// isPkgFile returns the status of isGoFile while ignoring test files
// ("_test.go" suffix).
func isPkgFile(fi os.FileInfo) bool {
	if !isGoFile(fi) {
		return false
	}
	if strings.HasSuffix(fi.Name(), "_test.go") {
		// skip test files.
		return false
	}
	return true
}

// isGoFile returns true for files with ".go" extension while ignoring hidden
// files ("." prefix).
func isGoFile(fi os.FileInfo) bool {
	if fi.IsDir() {
		// skip directories.
		return false
	}
	name := fi.Name()
	if len(name) >= 1 && name[0] == '.' {
		// skip hidden files ("." prefix).
		return false
	}
	if filepath.Ext(name) != ".go" {
		// skip files without ".go" extension.
		return false
	}
	return true
}
