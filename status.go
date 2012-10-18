package status

// Package contains status information about a package.
type Package struct {
	// Path is the import path of the package.
	Path string
	// CanGet specifies if it's possible to "go get" the package.
	CanGet bool
	// CanBuild specifies if it's possible to "go build" the package.
	CanBuild bool
}
