// Package status gathers information about the status of Go packages.
package status

import (
	"go/doc"
)

// Package contains status information about a package.
type Package struct {
	// Path is the import path of the package.
	Path string
	// CanGet specifies if it's possible to "go get" the package.
	CanGet bool
	// CanParse specifies if it's possible to parse the package's source.
	CanParse bool
	// CanBuild specifies if it's possible to "go build" the package.
	CanBuild bool
	// IsFork specifies if the package is a fork or not.
	IsFork bool
	// d is the package documentation for the entire package.
	d *doc.Package
}

// NewPkg returns a new package for the specified import path.
func NewPkg(importPath string) (pkg *Package) {
	pkg = &Package{
		Path: importPath,
	}
	return pkg
}
